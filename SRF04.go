package pidroponics

import (
	"container/ring"
	"encoding/binary"
	"fmt"
	"io"
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

	// start the background polling loop
	go s.tickerRead()

	return err
}

func (s *Srf04) tickerRead() {
	if s.readTic != nil {
		for range s.readTic.C {
			// TODO: Exit Condition
			s.Read()
		}
	} else {
		log.Println("Background polling for srf04@", s.devDevice, " disabled. Reinitialize with a time.Ticker to enable.")
	}
}

func (s *Srf04) Read() (int, error) {
	samp := make([]byte, 4)

	readFile, err := os.OpenFile(s.readPath, os.O_RDONLY, os.ModeDevice)
	if err != nil {
		return 0, err
	}
	defer readFile.Close()

	for {
		n, err := io.ReadFull(readFile, samp)
		if os.IsTimeout(err) {
			log.Fatal("  ReadFull TIMEOUT")
		}
		fmt.Println("ReadFull ", n, "bytes. Err: ", err)
		fmt.Println("    ", string(samp))
		fmt.Println("    ", binary.LittleEndian.Uint32(samp[0:4]))
		if err == io.EOF {
			break
		}
	}

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

