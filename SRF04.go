package pidroponics

import (
	"container/ring"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"time"
)

type Srf04 struct {
	Emitter		`json:"-"`
	Name		string
	Initialized bool

	devDevice	string
	readPath    string
	readTic		*time.Ticker `json:"-"`
	samples		*ring.Ring `json:"-"`
}

type Srf04State struct {
	Device		string
	Distance	float32
	Timestamp	int64
	sampleCount	int
	sum			int
}

func NewSrf04(devPath string) (*Srf04, error){
	var err error = nil

	s := &Srf04 {
		devDevice: devPath,
		Name: "",
		Initialized: false,
		samples: ring.New(10),
		readTic: nil,
		readPath: path.Join(devPath, "in_distance_raw"),
	}
	s.EmitterID = s

	// Initialize the ring with -1 for all the values.
	n := s.samples.Len()
	for i := 0; i < n; i++ {
		s.samples.Value = -1
		s.samples = s.samples.Next()
	}

	err = s.Initialize("", nil)
	if err == nil {
		_, err = s.Read()
	}

	return s, err
}

func (s *Srf04) Initialize(name string, readtic *time.Ticker) error {
	s.Name = name
	s.readTic = readtic

	_, err := os.Stat(s.readPath)
	f, err := os.Open(s.readPath)
	if err == nil {
		defer f.Close()
	}
	buf := make([]byte, 1)
	// This will likely err. We'll expect that.
	// Any error other than a timeout implies we have a device connected.
	_, err = f.Read(buf)
	if err != nil && os.IsTimeout(err) {
		fmt.Println("Timeout error. Not connected.")
		s.Initialized = false
	} else {
		fmt.Println("Non-timeout. Connected.")
		// Non-timeout. Likely -EIO. We're present and accounted for.
		s.Initialized = true
		err = nil
	}

	// start the background polling loop
	go s.tickerRead()

	return err
}



func (s *Srf04) GetState() *Srf04State {
	state := &Srf04State{
		Device:      s.Name,
		Distance:    -1,
		Timestamp:   time.Now().Unix(),
		sampleCount: 0,
		sum:         0,
	}

	// Add them up
	s.samples.Do(func(v interface{}) {
		val := v.(int)
		if val > 0 {
			state.sampleCount++
			state.sum += val
		}
	})

	// Do the division
	if state.sampleCount > 0 {
		state.Distance = float32(state.sum) / float32(state.sampleCount)
	}

	return state
}

func (s *Srf04) tickerRead() {
	if s.readTic != nil {
		for range s.readTic.C {
			if !s.Initialized {
				break
			}
			s.Read()
		}
	}
}

func (s *Srf04) Read() (int, error) {
	out, err := exec.Command("cat", s.readPath).Output()

	if err != nil {
		log.Fatal(err)
	}

	val, err := strconv.Atoi(string(out))
	if err != nil {
		s.samples.Value = val
		s.samples = s.samples.Next()
	}

	return val, err
}

