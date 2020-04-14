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
	err = h.EchoPin.In(gpio.PullDown, gpio.RisingEdge)
	if err == nil {
		go func() {
			for {
				// On edge change of the echo...
				h.EchoPin.WaitForEdge(-1)
				var t = time.Now().UnixNano()

				// compute the distance, clear the value.
				fmt.Println("between start - trigger: ", t, " ", h.triggeredAt)
				fmt.Println("     ", t - h.triggeredAt)
				fmt.Println("     /58 = ", (t - h.triggeredAt) / 58);
				h.triggeredAt = 0
			}
		}()
	} else {
		log.Print(err);
	}

	return h, err;
}

func (h *HCSR04) MeasureDistance() (error) {
	h.Lock()
	defer h.Unlock()

	var pulse time.Duration = 0

	pulse += 10 * time.Microsecond
	h.TriggerPin.Out(gpio.Low)
	h.triggeredAt = time.Now().UnixNano()
	h.TriggerPin.Out(gpio.High)
	time.Sleep(pulse)
	h.TriggerPin.Out(gpio.Low)

	return nil
}
