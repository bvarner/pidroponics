package pidroponics

import (
	"container/ring"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"time"
)

type Srf04 struct {
	Emitter		`json:"-"`
	Name		string

	devDevice	string
	readPath    string
	readTic		*time.Ticker `json:"-"`
	samples		*ring.Ring `json:"-"`
	initialized bool
}

type Srf04State struct {
	Device		string
	Distance	int
	Timestamp	int64
}

func NewSrf04(devPath string) (*Srf04, error){
	var err error = nil

	s := &Srf04 {
		devDevice: devPath,
		Name: "",
		initialized: false,
		samples: ring.New(250),
		readTic: nil,
		readPath: path.Join(devPath, "in_distance_raw"),
	}
	s.EmitterID = s

	// Initialize the ring with -1s
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
	s.initialized = err == nil

	// start the background polling loop
	go s.tickerRead()

	return err
}

func (s *Srf04) Close() {
	s.initialized = false
}

func (s *Srf04) tickerRead() {
	if s.readTic != nil {
		for range s.readTic.C {
			if !s.initialized {
				log.Println("Terminating polling loop for: ", s.Name)
				return
			}

			s.Read()
		}
	} else {
		log.Println("Background polling for srf04@", s.devDevice, " disabled. Reinitialize with a time.Ticker to enable.")
	}
}

func (s *Srf04) Read() (int, error) {
	var err error = nil
	if s.initialized {
		// read the value from the sensor device
		buf, err := ioutil.ReadFile(s.readPath)
		if err == nil {
			i, err := strconv.Atoi(string(buf))
			if err == nil {
				s.samples.Value = i
				s.samples = s.samples.Next()
			}
		}
	}

	return s.samples.Prev().Value.(int), err
}

