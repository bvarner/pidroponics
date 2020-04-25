package pidroponics

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"time"
)

type Relay struct {
	Emitter		`json:"-"`
	Device		string
	IsOn		bool

	devDevice	string
	initialized bool
}

type RelayState struct {
	Device		string
	IsOn		bool
	Timestamp	int64
}

func NewRelay(devPath string, Device string)(*Relay, error) {
	var err error = nil

	r := &Relay{
		devDevice: devPath,
		Device: Device,
		initialized: false,
	}
	err = r.readState()

	return r, err
}

func (r *Relay) eventName() string {
	fmt.Println("Get eventName()")
	return r.Device
}

func (r *Relay) readState() error {
	statebuf, err := ioutil.ReadFile(path.Join(r.devDevice, "brightness"))
	if err != nil {
		return err
	}
	brightness, err := strconv.Atoi(string(statebuf)[:len(string(statebuf))-1])
	if err != nil {
		return err
	}
	r.IsOn = brightness > 0
	r.initialized = true

	return nil
}

func (r *Relay) writeState() error {
	var state []byte

	if r.IsOn {
		state = []byte("1")
	} else {
		state = []byte("0")
	}
	ioutil.WriteFile(path.Join(r.devDevice, "brightness"), state, os.ModeDevice)

	return nil
}

func (r *Relay) GetState() RelayState {
	return RelayState {
		Device:    r.Device,
		IsOn:      r.IsOn,
		Timestamp: time.Now().Unix(),
	}
}

func (r *Relay) SetOn(state bool) error {
	var err error = nil

	change := r.initialized && r.IsOn != state

	r.IsOn = state
	err = r.writeState()

	if err == nil && change {
		r.Emit(r.GetState())
	}

	return err
}
