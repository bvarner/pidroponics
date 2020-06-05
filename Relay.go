package pidroponics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type Relay struct {
	Emitter `json:"-"`
	gauge	prometheus.Gauge
	Device  string
	IsOn    bool
	toggled  time.Time
	Manual  bool

	devDevice	string
	initialized bool
}

type RelayState struct {
	Device		string
	IsOn		bool
	Manual		bool
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

			var r = Relay{
				devDevice: path.Join(basedir, file.Name()),
				gauge: promauto.NewGauge(prometheus.GaugeOpts{
					Namespace:   "pidroponics",
					Subsystem:   "relay",
					Name:        relayNames[relayNum] + "_value",
					Help:        "State of a Relay. 0 = off, 1 = on",
				}),
				Device:      relayNames[relayNum],
				initialized: false,
				Manual:      false,
			}
			r.EmitterID = &r
			err = r.readState()
			if err != nil {
				break
			}

			fmt.Println("Added relay: ", r.devDevice)
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
	r.toggled = time.Now()
	if r.IsOn {
		r.gauge.Set(1)
	} else {
		r.gauge.Set(0)
	}
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
		Manual:	   r.Manual,
		Timestamp: time.Now().Unix(),
	}
}

func (r *Relay) SetOn(state bool) error {
	var err error = nil

	// Only change if we're initialized and different than the current state, and have been set longer than 2 seconds.
	change := r.initialized && r.IsOn != state && time.Now().Sub(r.toggled).Seconds() > 2
	if change {
		r.IsOn = state
		err = r.writeState()

		if err == nil {
			if r.IsOn {
				r.gauge.Set(1)
			} else {
				r.gauge.Set(0)
			}
			r.Emit(r.GetState())
		}
	}

	return err
}

func (r *Relay) SetManual(manual bool) {
	r.Manual = manual
}
