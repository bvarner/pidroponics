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
	triggeredWhen	*time.Time	`json:"-"`
	echoStart		*time.Time	`json:"-"`
	echoEnd			*time.Time	`json:"-"`

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
	}

	// Set the trigger low by default.
	h.TriggerPin.Out(gpio.Low)
	// Set the echo up to listen.
	err = h.EchoPin.In(gpio.PullDown, gpio.BothEdges)
	if err == nil {
		go func() {
			var maxResult, _ = time.ParseDuration("38us")
			for {
				// On edge change of the echo...
				h.EchoPin.WaitForEdge(-1)
				var t = time.Now()
				if h.echoStart == nil {
					h.echoStart = &t
				} else if h.echoEnd == nil {
					h.echoEnd = &t
				}

				// compute the distance, clear the value.
				if h.echoStart != nil && h.echoEnd != nil {
					// compute this down to centimeters
					var delay = h.echoEnd.Sub(*h.echoEnd).Seconds();
					fmt.Println("Delay: ", delay)
					if delay < maxResult.Seconds() {
						h.Distance = float64(delay) / 58
					} else {
						// Or if we can't detect things in front of us...
						h.Distance = math.NaN()
					}

					// Clear Status
					h.echoStart = nil;
					h.echoEnd = nil;
				}
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
	h.TriggerPin.Out(gpio.High)
	time.Sleep(pulse)
	h.TriggerPin.Out(gpio.Low)

	return nil
}
