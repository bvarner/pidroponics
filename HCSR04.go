package pidroponics

import (
	"fmt"
	"log"
	"math"
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
	"sync"
	"time"
)

type HCSR04 struct {
	TriggerPin		gpio.PinIO	`json:"-"`
	EchoPin			gpio.PinIO	`json:"-"`

	sync.Mutex					`json:"-"`
	triggeredAt		int64 		`json:"-"`

	Distance		float64
}

func NewHCSR04(triggerPinName string, echoPinName string)(*HCSR04, error) {
	var err error = nil
	if _, err = host.Init(); err != nil {
		log.Fatal(err)
	}

	h := &HCSR04{
		TriggerPin: gpioreg.ByName(triggerPinName),
		EchoPin: gpioreg.ByName(echoPinName),
		Distance: math.NaN(),
		triggeredAt: 0,
	}

	// Set the trigger low by default.
	h.TriggerPin.Out(gpio.Low)
	// Set the echo up to listen.
	err = h.EchoPin.In(gpio.PullDown, gpio.BothEdges)

	return h, err;
}

func (h *HCSR04) MeasureDistance() (error) {
	h.Lock()
	defer h.Unlock()

	var echoStart = int64(0)
	var echoEnd  = int64(0)

	h.TriggerPin.Out(gpio.Low)
	h.triggeredAt = time.Now().UnixNano()
	h.TriggerPin.Out(gpio.High)
	time.Sleep(10 * time.Microsecond)
	h.TriggerPin.Out(gpio.Low)

	if h.EchoPin.WaitForEdge(-1) {
		echoStart = time.Now().UnixNano()
		if h.EchoPin.WaitForEdge(-1) {
			echoEnd = time.Now().UnixNano()
		}
	}

	fmt.Println("triggered: ", h.triggeredAt, " echoStart: ", echoStart, " echoEnd: ", echoEnd)
	fmt.Println("   pulse duration: ", float64(echoEnd - echoStart));

	h.Distance = (float64(echoEnd - echoStart) * float64(time.Nanosecond) / float64(time.Second)) * 17150

	return nil
}
