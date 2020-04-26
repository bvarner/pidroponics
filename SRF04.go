package pidroponics

import (
	"container/ring"
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

type Srf04 struct {
	Emitter		`json:"-"`
	Name		string

	devDevice	string
	readPath    string
	readFile	*os.File
	readTic		*time.Ticker `json:"-"`
	samples		*ring.Ring `json:"-"`
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
		samples: ring.New(250),
		readTic: nil,
		readPath: path.Join(devPath, "in_distance_raw"),
		readFile: nil,
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
	if err == nil {
		s.readFile, err = os.Open(s.readPath)
	}

	// start the background polling loop
	go s.tickerRead()

	return err
}

func (s *Srf04) Close() {
	s.readFile.Close()
}

func (s *Srf04) tickerRead() {
	if s.readTic != nil {
		for range s.readTic.C {
			if s.readFile == nil {
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
	samp := make([]byte, 32)

	n, err := s.readFile.Read(samp)
	if err != nil {
		log.Fatal("Error reading from: ", s.readPath, err)
	}
	fmt.Println("Read ", n, " bytes from: ", s.readPath, " Err: ", err)

	return 0, err
	//
	//
	//
	//var err error = nil
	//if s.initialized {
	//	// read the value from the sensor device
	//	buf, err := ioutil.ReadFile(s.readPath)
	//	// In the case that we get zero bytes, we consider this an unexpected EOF
	//	// and an 'unconnected' device.
	//	fmt.Println("Read: ", len(buf), " bytes from ", s.readPath)
	//	if len(buf) == 0 {
	//		err = io.ErrUnexpectedEOF
	//	}
	//
	//	if err == nil {
	//		i, err := strconv.Atoi(string(buf))
	//		if err == nil {
	//			fmt.Println("raw distance: ", i)
	//			s.samples.Value = i
	//			s.samples = s.samples.Next()
	//		}
	//	}
	//}
	//
	//return s.samples.Prev().Value.(int), err
}

