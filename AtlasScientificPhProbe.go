package pidroponics

import (
	"container/ring"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"time"
)

type AtlasScientificPhProbe struct {
	Emitter 	`json:"="`
	Initialized bool

	readPath	string
	readBuf		[]byte
	readFile 	*os.File `json:"-"`
	readTic 	*time.Ticker `json:"-"`
	sampleLock	sync.Mutex
	samples 	*ring.Ring `json:"-"`
	emitTic		*time.Ticker
}

type PhProbeState struct {
	Ph 			float32
	Timestamp   int64
	sampleCount	int
	sum			float32
}

// Creates a new AtlasScientificPhProbe from the given ADC Channel.
// This essentially creates a user-space driver.
func NewAtlasScientificPhProbe(readPath string, readTic *time.Ticker, emitTic *time.Ticker) (AtlasScientificPhProbe, error) {
	var err error = nil

	p := AtlasScientificPhProbe{
		Initialized: false,
		readPath:    readPath,
		readBuf:     make([]byte, 4096),
		readFile:    nil,
		readTic:     readTic,
		samples:     ring.New(3),
		emitTic:	 emitTic,
	}
	p.EmitterID = &p

	// Initialize the ring with -1
	p.sampleLock.Lock()
	n := p.samples.Len()
	for i := 0; i < n; i++ {
		p.samples.Value = float32(-1)
		p.samples = p.samples.Next()
	}
	p.sampleLock.Unlock()

	err = p.Initialize()
	if err == nil {
		_, err = p.Read()
	}

	go p.emitLoop()
	fmt.Println("Added Atlas Scientific PH: ", p.readPath)
	return p, err
}

func (p *AtlasScientificPhProbe) eventName() string {
	return "PhProbe"
}

func (p *AtlasScientificPhProbe) Initialize() error {
	_, err := os.Stat(p.readPath)
	f, err := os.OpenFile(p.readPath, os.O_RDONLY, os.ModeDevice)
	buf := make([]byte, 1)

	_, err = f.Read(buf)
	if err != nil && os.IsTimeout(err) {
		p.Initialized = false
		p.readFile = nil
		f.Close()
	} else {
		p.Initialized = true
		p.readFile = f
		err = nil
	}

	go p.tickerRead()

	return err
}

func (p *AtlasScientificPhProbe) Close() error {
	var err error = nil
	if p.readFile != nil {
		err = p.readFile.Close()
		p.readFile = nil
		p.Initialized = false
	}
	return err
}

func (p *AtlasScientificPhProbe) GetState() *PhProbeState {
	state := &PhProbeState{
		Timestamp:   time.Now().Unix(),
		Ph: 		-1,
		sampleCount: 0,
		sum:         0,
	}

	// Add them up
	p.sampleLock.Lock()
	defer p.sampleLock.Unlock()
	p.samples.Do(func(v interface{}) {
		val := v.(float32)
		if val > 0 {
			state.sampleCount++
			state.sum += val
		}
	})

	// Do the division
	if state.sampleCount > 0 {
		state.Ph = state.sum / float32(state.sampleCount)
	}

	// TODO: Standard Deviation
	return state
}

func (p *AtlasScientificPhProbe) tickerRead() {
	if p.readTic != nil {
		for range p.readTic.C {
			if !p.Initialized {
				break
			}
			p.Read()
		}
	}
}

func (p *AtlasScientificPhProbe) emitLoop() {
	for range p.emitTic.C {
		if p.Initialized {
			go func() {p.Emit(p.GetState())}()
		}
	}
}

func (p *AtlasScientificPhProbe) Read() (float32, error) {
	bytesRead := 0
	_, err := p.readFile.Seek(0, 0)

	// Loop until N > 0 AND err != EOF && err != timeout.
	if err == nil {
		n := 0
		for {
			n, err = p.readFile.Read(p.readBuf)
			bytesRead += n
			if os.IsTimeout(err) {
				// bail out.
				bytesRead = 0
				break
			}
			if err == io.EOF {
				// Success!
				break
			}
			// Any other err means 'keep trying to read.'
		}
	}

	// We shouldn't ever get here if we didn't read anything.
	if bytesRead > 0 { // paranoia
		val, err := strconv.ParseFloat(string(p.readBuf[:bytesRead-1]), 32)
		if err == nil {
			val = (-5.6548 * val) + 15.509

			p.sampleLock.Lock()
			defer p.sampleLock.Unlock()
			p.samples.Value = float32(val)
			p.samples = p.samples.Next()
		}
		return float32(val), err
	}

	return 0, err
}