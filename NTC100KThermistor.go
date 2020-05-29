package pidroponics

import (
	"container/ring"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"
)

type NTC100KThermistor struct {
	Emitter		`json:"-"`
	Name 		string
	Initialized bool

	readPath	string
	readBuf		[]byte
	readFile	*os.File
	readTic		*time.Ticker
	samples		*ring.Ring
	sampleLock	sync.Mutex
	emitTic 	*time.Ticker
}

type ThermistorState struct {
	Name 	string
	Temperature float64
	Timestamp	int64
	sampleCount	int
	sum			float64
}

func DetectNTC100KThermistors(readtic *time.Ticker) ([]NTC100KThermistor, error) {
	thermistorNames := []string{"sump_temp", "inlet_temp", "ambient_temp"}

	files, err := ioutil.ReadDir("/sys/bus/platform/drivers/ntc-thermistor")
	if err != nil {
		return nil, err
	}

	var thermistors []NTC100KThermistor
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "therm@") {
			thermNum, err := strconv.Atoi(file.Name()[len(file.Name()) - 1:])
			if err != nil {
				return thermistors, err
			}
			devPath := path.Join("/sys/bus/platform/drivers/ntc-thermistor", file.Name(), "hwmon")

			// therm0/hwmon
			devFiles, err := ioutil.ReadDir(devPath)
			if err != nil {
				log.Fatal("Unable to open dev path:", devPath, err)
			}

			// hwmonX/temp1_input
			readPath := path.Join(devPath, devFiles[0].Name(), "temp1_input")

			_, err = os.Stat(readPath)
			if err != nil {
				return thermistors, err
			}

			t := NTC100KThermistor{
				Name:			thermistorNames[thermNum],
				Initialized: 	false,
				readPath:		readPath,
				readBuf:	 	make([]byte, 4096),
				readFile:		nil,
				readTic:		readtic,
				samples:		ring.New(10),
				emitTic: 		time.NewTicker(time.Second),
			}
			t.EmitterID = &t

			// Initialize samples to NaN, a float64.
			t.sampleLock.Lock()
			n := t.samples.Len()
			for i := 0; i < n; i++ {
				t.samples.Value = math.NaN()
				t.samples = t.samples.Next()
			}
			t.sampleLock.Unlock()

			err = t.Initialize()
			if err == nil {
				_, err = t.Read()
			}

			go t.emitLoop()
			fmt.Println("Added NTC100kThermistor: ", t.readPath)
			thermistors = append(thermistors, t)
		}
	}

	return thermistors, nil
}

func (t *NTC100KThermistor) eventName() string {
	return t.Name
}

func (t *NTC100KThermistor) Initialize() error {
	f, err := os.OpenFile(t.readPath, os.O_RDONLY, os.ModeDevice)
	buf := make([]byte, 1)
	_, err = f.Read(buf)
	if err != nil && os.IsTimeout(err) {
		t.Initialized = false
		t.readFile = nil
		f.Close()
	} else {
		t.Initialized = true
		t.readFile = f
		err = nil
	}

	go t.tickerRead()

	return err
}

func (t *NTC100KThermistor) Close() error {
	var err error = nil
	if t.readFile != nil {
		err = t.readFile.Close()
		t.readFile = nil
		t.Initialized = false
		t.emitTic.Stop()
	}
	return err
}

func (t *NTC100KThermistor) GetState() *ThermistorState {
	state := &ThermistorState{
		Name:        t.Name,
		Temperature: math.NaN(),
		Timestamp:   time.Now().Unix(),
		sampleCount: 0,
		sum:         0,
	}

	// Add up the samples
	t.sampleLock.Lock()
	defer t.sampleLock.Unlock()
	t.samples.Do(func(v interface{}) {
		val := v.(float64)
		if val != math.NaN() {
			state.sampleCount++
			state.sum += val
		}
	})

	// Do the division.
	if state.sampleCount > 0 {
		state.Temperature = state.sum / float64(state.sampleCount)
	}

	if math.IsNaN(state.Temperature) {
		state.Temperature = -999
	}

	// TODO: Standard Deviation

	return state
}

func (t *NTC100KThermistor) tickerRead() {
	if t.readTic != nil {
		for range t.readTic.C {
			if !t.Initialized {
				break
			}
			t.Read()
		}
	}
}

func (t *NTC100KThermistor) emitLoop() {
	for range t.emitTic.C {
		if t.Initialized {
			go func() {t.Emit(t.GetState())}()
		}
	}
}

func (t *NTC100KThermistor) Read() (float64, error) {
	// Seek should tell us the new offset (0) and no err.
	bytesRead := 0
	_, err := t.readFile.Seek(0, 0)

	// Loop until N > 0 AND err != EOF && err != timeout.
	if err == nil {
		n := 0
		for {
			n, err = t.readFile.Read(t.readBuf)
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
		val, err := strconv.ParseFloat(string(t.readBuf[:bytesRead-1]), 64)
		if err == nil {
			t.sampleLock.Lock()
			defer t.sampleLock.Unlock()
			t.samples.Value = val / 1000 // The value from the device is degrees celcius * 1000, to avoid decimal places.
			t.samples = t.samples.Next()
		}
		return val, err
	}

	return 0, err
}
