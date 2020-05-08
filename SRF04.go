package pidroponics

import (
	"container/ring"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type Srf04 struct {
	Emitter		`json:"-"`
	Name		string
	Initialized bool

	devDevice	string
	readPath    string
	readBuf		[]byte `json:"-"`
	readFile	*os.File `json:"-"`
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

func DetectSrf04(readtic *time.Ticker) ([]Srf04, error) {
	transponderNames := []string{"Sump", "Inlet", "Outlet"}

	files, err := ioutil.ReadDir("/sys/bus/platform/drivers/srf04-gpio")
	if err != nil {
		return nil, err
	}

	var sensors []Srf04
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "proximity@") {
			proximityNum, err := strconv.Atoi(file.Name()[len(file.Name()) - 1:])
			if err != nil {
				return sensors, err
			}
			devPath := path.Join("/sys/bus/platform/drivers/srf04-gpio", file.Name())

			// find the iio:deviceX path.
			// and update devPath.
			devFiles, err := ioutil.ReadDir(devPath)
			if err != nil {
				log.Fatal("Unable to open dev path:", devPath, err)
			}
			for _, file := range devFiles {
				if strings.HasPrefix(file.Name(), "iio:device") {
					devPath = path.Join(devPath, file.Name())
					break
				}
			}

			// Construct the object the right way.
			s := Srf04{
				devDevice:   devPath,
				Name:        transponderNames[proximityNum],
				Initialized: false,
				samples:     ring.New(30),
				readTic:     readtic,
				readPath:    path.Join(devPath, "in_distance_raw"),
				readBuf:     make([]byte, 4096),
			}
			s.EmitterID = &s

			// Initialize the ring with -1 for all the values.
			n := s.samples.Len()
			for i := 0; i < n; i++ {
				s.samples.Value = -1
				s.samples = s.samples.Next()
			}

			// Use the device number as the tick offset
			err = s.Initialize(proximityNum)
			if err == nil {
				_, err = s.Read()
			}

			fmt.Println("Added srf04: ", s.devDevice)
			sensors = append(sensors, s)
		}
	}

	return sensors, nil
}

func (s *Srf04) eventName() string {
	return s.Name
}

func (s *Srf04) Initialize(tickoffset int) error {
	_, err := os.Stat(s.readPath)
	f, err := os.OpenFile(s.readPath, os.O_RDONLY, os.ModeDevice)
	buf := make([]byte, 1)
	// This will likely err. We'll expect that.
	// Any error other than a timeout implies we have a device connected.
	_, err = f.Read(buf)
	if err != nil && os.IsTimeout(err) {
		s.Initialized = false
		s.readFile = nil
		f.Close()
	} else {
		// Non-timeout. Likely -EIO. We're present and accounted for.
		s.Initialized = true
		s.readFile = f
		err = nil
	}

	// start the background polling loop
	go s.tickerRead(tickoffset)

	return err
}

func (s *Srf04) Close() error {
	if s.readFile != nil {
		return s.readFile.Close()
	}
	return nil
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

	// TODO: Standard Deviation

	return state
}

func (s *Srf04) tickerRead(tickoffset int) {
	count := tickoffset
	if s.readTic != nil {
		for range s.readTic.C {
			if !s.Initialized {
				break
			}
			if count <= 0 {
				s.Read()
				count = 3
			}
			count--
		}
	}
}

func (s *Srf04) Read() (int, error) {
/** Reliable, but slow.
	out, err := exec.Command("cat", s.readPath).Output()

	if err != nil {
		log.Fatal(err)
	}

	val, err := strconv.Atoi(string(out[:len(out) - 2]))
	if err == nil {
		s.samples.Value = val
		s.samples = s.samples.Next()
	}
 */
	// Seek should tell us the new offset (0) and no err.
	bytesRead := 0
	_, err := s.readFile.Seek(0, 0)

	// Loop until N > 0 AND err != EOF && err != timeout.
	if err == nil {
		n := 0
		for {
			n, err = s.readFile.Read(s.readBuf)
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
		val, err := strconv.Atoi(string(s.readBuf[:bytesRead-1]))
		if err == nil {
			s.samples.Value = val
			s.samples = s.samples.Next()
		}
		// TODO: Use a different ticker for this.
		go func() {s.Emit(s.GetState())}()
		return val, err
	}

	return 0, err
}
