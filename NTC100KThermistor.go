package pidroponics

import (
	"container/ring"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path"
	"strconv"
	"strings"
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
}

type ThermistorState struct {
	Name 	string
	Temperature float64
	Timestamp	int64
	sampleCount	int
	sum			int
}

func DetectNTC100KThermistors(readtic *time.Ticker) ([]NTC100KThermistor, error) {
	thermistorNames := []string{"Sump", "Inlet", "Ambient"}

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
			}
			t.EmitterID = &t

			// Initialize samples to NaN, a float64.
			n := t.samples.Len()
			for i := 0; i < n; i++ {
				t.samples.Value = math.NaN()
				t.samples = t.samples.Next()
			}

			err = t.Initialize()
			if err == nil {
				_, err = t.Read()
			}

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
	if t.readFile != nil {
		t.readFile.Close()
	}
	return nil
}

func (t *NTC100KThermistor) GetState() *ThermistorState {
	// TODO: Implement
	return nil
}

func (t *NTC100KThermistor) tickerRead() {
	// TODO: implement
	return
}

func (t *NTC100KThermistor) Read() (float64, error) {
	// TODO: implement
	return 0, nil
}
