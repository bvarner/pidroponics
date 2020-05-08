package pidroponics

import (
	"container/ring"
	"os"
	"time"
)

type AtlasScientificPhProbe struct {
	Emitter 	`json:"="`

	readFile *os.File `json:"-"`
	readTic *time.Ticker `json:"-"`
	samples *ring.Ring `json:"-"`
}

func NewAtlasScientificPhProbe(devPath string) (*AtlasScientificPhProbe, error) {
	var err error = nil

	p := &AtlasScientificPhProbe {

	}
	p.EmitterID = p

	return p, err
}

func (p *AtlasScientificPhProbe) eventName() string {
	return "PhProbe"
}

