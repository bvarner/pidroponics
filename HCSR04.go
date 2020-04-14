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
	echoStart		int64		`json:"-"`
	echoEnd			int64 		`json:"-"`

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
		echoStart: 0,
		echoEnd: 0,
	}

	// Set the trigger low by default.
	h.TriggerPin.Out(gpio.Low)
	// Set the echo up to listen.
	err = h.EchoPin.In(gpio.PullDown, gpio.BothEdges)
	if err == nil {
		go func() {
			for {
				// On edge change of the echo...
				h.EchoPin.WaitForEdge(-1)
				var t = time.Now()
				if h.echoStart == 0 {
					h.echoStart = t.UnixNano()
					h.echoEnd = 0
				} else if h.echoEnd == 0 {
					h.echoEnd = t.UnixNano()
				}

				// compute the distance, clear the value.
				if h.echoStart != 0 && h.echoEnd != 0 {
					// compute this down to centimeters
					var highTime = h.echoEnd - h.echoStart
					fmt.Println("echo end - start: ", h.echoEnd, " ", h.echoStart)
					fmt.Println("  highTime ns: ", highTime)
					fmt.Println("  highTime us: ", int64(time.Nanosecond) * highTime / int64(time.Microsecond))
					fmt.Println("    multplied: ", highTime * 340 / 2)
					fmt.Println("between trigger - start: ", h.triggeredAt, " ", h.echoStart)
					fmt.Println("     ", h.echoStart - h.triggeredAt)
					fmt.Println("     /58 = ", (h.echoStart - h.triggeredAt) / 58);
					fmt.Println("between trigger - end: ", h.triggeredAt, " ", h.echoEnd)
					fmt.Println("     ", h.echoEnd - h.triggeredAt)
					fmt.Println("     /58 = ", (h.echoEnd - h.triggeredAt) / 58);

					// Clear Status
					h.echoStart = 0
					h.echoEnd = 0
					h.triggeredAt = 0
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
	h.triggeredAt = time.Now().UnixNano()
	h.TriggerPin.Out(gpio.High)
	time.Sleep(pulse)
	h.TriggerPin.Out(gpio.Low)

	return nil
}
