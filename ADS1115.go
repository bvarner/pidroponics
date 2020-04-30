package pidroponics

import (
	"container/ring"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"time"
)

const SPS = 128

type ADSChannel struct {
	Emitter 	`json:"-"`
	Name 		string
	Channel		string

	samples		*ring.Ring `json:"-"`
}

type ADS1115 struct {
	Channels	[]ADSChannel
	Initialized	bool

	devPath		string
	devDevice	*os.File

	trigger		*os.File
	triggerTic  *time.Ticker

	idxTime		int
	idxVoltage 	int
}

/* Assumes an ADS1115 with four independent channels.
 */
func NewADS1115(devPath string) (*ADS1115, error) {
	var err error = nil

	a := &ADS1115 {
		devPath: devPath,
		Channels: []ADSChannel{
			{
				Name:    "A",
				Channel: "in_voltage0",
				samples: ring.New(SPS),
			},
			{
				Name:    "B",
				Channel: "in_voltage1",
				samples: ring.New(SPS),
			},
			{
				Name:    "C",
				Channel: "in_voltage2",
				samples: ring.New(SPS),
			},
			{
				Name:    "D",
				Channel: "in_voltage3",
				samples: ring.New(SPS),
			},
		},
	}

	// DRY, yo.
	initChannel := func(c ADSChannel) {
		n := c.samples.Len()
		for i := 0; i < n; i++ {
			c.samples.Value = -1
			c.samples = c.samples.Next()
		}
		// Set the sample rate to the size of the channel ring.
		deviceEcho(path.Join(a.devPath, c.Channel + "_sampling_frequency"), []byte(strconv.Itoa(c.samples.Len())), 0644)
	}

	for _, c := range a.Channels {
		initChannel(c)
	}

	// If the sysfs trigger doesn't exist, then we try to create one.
	triggerDev := "/sys/bus/iio/devices/iio_sysfs_trigger/trigger0"
	if _, err := os.Stat(triggerDev); err != nil {

		// Make sure we have the proper sysfs bits.
		if _, err := os.Stat("/sys/bus/iio/devices/iio_sysfs_trigger"); err != nil {
			fmt.Println("Sysfs Triggering Unavilable.", err)
			// TODO: Handle polling each channel instead.
			return a, err
		}

		// Create trigger0 if it doesn't exist.
		if _, err := os.Stat("/sys/bus/iio/devices/iio_sysfs_trigger/trigger0"); err != nil {
			// Create trigger0 since it does not exist
			if err := deviceEcho("/sys/bus/iio/devices/iio_sysfs_trigger/add_trigger", []byte("0"), 0200); err != nil {
				return a, err
			}
		}
	}

	// By the time we get here we know we have sysfstrigger0
	triggerName, err := ioutil.ReadFile(triggerDev + "/name")
	if err != nil {
		// TODO: Handle polling each channel instead.
		return a, err
	}

	// Disable the buffer and Set the trigger as the iio:device trigger.
	err = deviceEcho(a.devPath + "/buffer/enable", []byte("0"), 0)
	if err != nil {
		// TODO: Handle polling each channel instead.
		return a, err
	}
	if err := deviceEcho(a.devPath + "/trigger/current_trigger", triggerName, 0); err != nil {
		return a, err
	}

	// enable the channels we want...
	deviceEcho(a.devPath + "/scan_elements/in_timestamp_en", []byte("1"), 0644)
	deviceEcho(a.devPath + "/scan_elements/in_voltage0_en", []byte("1"), 0644)

	// Find out what index the items are.
	buf, err := ioutil.ReadFile(a.devPath + "/scan_elements/in_timestamp_index")
	if err != nil {
		return a, err
	}
	a.idxTime, err = strconv.Atoi(string(buf))

	buf, err = ioutil.ReadFile(a.devPath + "/scan_elements/in_voltage0_index")
	if err != nil {
		return a, err
	}
	a.idxVoltage, err = strconv.Atoi(string(buf))


	// Enable the buffering to the device file.
	err = deviceEcho(a.devPath + "/buffer/enable", []byte("1"), 0)

	// Identify the /dev device...
	_, devFileName := path.Split(a.devPath)
	a.devDevice, err = os.Open(path.Join("/dev", devFileName))
	if err != nil {
		return a, err
	}
	go a.readLoop()

	fmt.Println("Opening Trigger and starting up.")
	a.trigger, err = os.OpenFile(path.Join(triggerDev, "trigger_now"), os.O_WRONLY | os.O_SYNC, 0)
	if err != nil {
		return a, err
	}

	// Every tick, write to the file.
	a.triggerTic = time.NewTicker(time.Second / SPS)
	a.Initialized = true
	go a.tickerRead()

	return a, err
}

func (a *ADS1115) Close() {
	a.Initialized = false
	a.devDevice.Close()
	a.trigger.Close()
}

func (a *ADS1115) tickerRead() error {
	var err error = nil
	if a.triggerTic != nil {
		for range a.triggerTic.C {
			if !a.Initialized || err != nil {
				break
			}
			_, err = a.trigger.Write([]byte("1"))
		}
	}
	return err
}

func (a *ADS1115) Read() error {
	_, err := a.trigger.Write([]byte("1"))
	return err
}

/*
voltage0-voltage1:  (input, index: 0, format: le:S16/16>>0)
voltage0-voltage3:  (input, index: 1, format: le:S16/16>>0)
voltage1-voltage3:  (input, index: 2, format: le:S16/16>>0)
voltage2-voltage3:  (input, index: 3, format: le:S16/16>>0)
voltage0:  (input, index: 4, format: le:S16/16>>0)
voltage1:  (input, index: 5, format: le:S16/16>>0)
voltage2:  (input, index: 6, format: le:S16/16>>0)
voltage3:  (input, index: 7, format: le:S16/16>>0)
timestamp:  (input, index: 8, format: le:S64/64>>0)
*/
func (a *ADS1115) readLoop() {
	samp := make([]byte, 16)
	// TODO create a buffer.
	for {
		n, _ := a.devDevice.Read(samp)
		// should be two bytes per enabled sample channel + 4 bytes timestamp.
		v0 := binary.LittleEndian.Uint16(samp[0:2])
		v1 := binary.LittleEndian.Uint16(samp[2:4])
		v2 := binary.LittleEndian.Uint16(samp[4:6])
		v3 := binary.LittleEndian.Uint16(samp[6:8])
		ts := tsConvert(samp[8:16])

		// slice out the bytes and record the values.
		fmt.Println("ADS Read: ", n)
		fmt.Println("", ts, " v0:", v0, " v1:", v1, " v2:", v2, " v3:", v3)
		break
	}
}

func deviceEcho(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY, perm)
	defer f.Close()
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// Returns a int64 from an 8 byte buffer
func tsConvert(b []byte) int64 {
	_ = b[7] // bounds check hint to compiler; see golang.org/issue/14808
	return int64(b[0]) | int64(b[1])<<8 | int64(b[2])<<16 | int64(b[3])<<24 |
		int64(b[4])<<32 | int64(b[5])<<40 | int64(b[6])<<48 | int64(b[7])<<56
}
