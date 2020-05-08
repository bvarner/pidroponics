package pidroponics

import (
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
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

func DetectRelays() ([]Relay, error) {
	basedir := "/sys/bus/platform/drivers/leds-gpio/leds/leds"
	files, err := ioutil.ReadDir(basedir)
	if err != nil {
		return nil, err
	}

	var relays []Relay
	relayNames := []string{"Lights", "Pump", "Fan", "Valve"}


	for _, file := range files {
		if strings.HasPrefix(file.Name(), "relay") {
			relayNum, err := strconv.Atoi(file.Name()[len(file.Name()) - 1:])
			if err != nil {
				return relays, err
			}

			r := Relay {
				devDevice: path.Join(basedir, file.Name()),
				Device: relayNames[relayNum],
				initialized: false,
			}
			r.EmitterID = &r
			err = r.readState()
			if err != nil {
				break
			}

			relays = append(relays, r)
		}
	}

	return relays, err
}

func (r *Relay) eventName() string {
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
