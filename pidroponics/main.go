package main

import (
	"encoding/json"
	"flag"
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	"github.com/bvarner/pidroponics"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var broker *pidroponics.Broker

var handler http.Handler

var relays[] pidroponics.Relay
var transponders[] pidroponics.Srf04
var thermistors[] pidroponics.NTC100KThermistor
var phProbe pidroponics.AtlasScientificPhProbe

var relayURIMatcher = regexp.MustCompile("^/?relays/([0-3])$")
var transponderURIMatcher = regexp.MustCompile("^/?waterlevels/([0-2])$")
var thermistorURIMatcher = regexp.MustCompile("^/?temperatures/([0-2])$")

func redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://" + r.Host + r.RequestURI, http.StatusMovedPermanently)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	// Push some things if we know what our request is.
	if r.URL.Path == "/" || r.URL.Path == "/index.html" {
/*		p, ok := w.(http.Pusher)
		if ok {
			p.Push("/events", nil)
			p.Push("/style.css", nil)
			p.Push("/App.js", nil)
		}
*/	}

	handler.ServeHTTP(w, r)
}

func PhControl(w http.ResponseWriter, r *http.Request) {
	var err error = nil

	if r.Method == "GET" {
		err = json.NewEncoder(w).Encode(phProbe.GetState())
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func WaterLevelControl(w http.ResponseWriter, r *http.Request) {
	var err error = nil

	if r.Method == "GET" {
		if matched, _ := regexp.MatchString("^.waterlevels/?$", r.URL.Path); matched {
			// Dump the whole shebang.
			err = json.NewEncoder(w).Encode(transponders)
		} else {
			// look for an individual...
			matches := transponderURIMatcher.FindStringSubmatch(r.URL.Path)
			if len(matches) == 2 {
				idx, err := strconv.Atoi(matches[1])
				if err == nil {
					err = json.NewEncoder(w).Encode(transponders[idx].GetState())
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(http.StatusText(http.StatusBadRequest)))
			}
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func TempControl(w http.ResponseWriter, r *http.Request) {
	var err error = nil

	if r.Method == "GET" {
		if matched, _ := regexp.MatchString("^.temperatures/?$", r.URL.Path); matched {
			// Dump the whole shebang.
			err = json.NewEncoder(w).Encode(thermistors)
		} else {
			// look for an individual...
			matches := thermistorURIMatcher.FindStringSubmatch(r.URL.Path)
			if len(matches) == 2 {
				idx, err := strconv.Atoi(matches[1])
				if err == nil {
					err = json.NewEncoder(w).Encode(thermistors[idx].GetState())
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(http.StatusText(http.StatusBadRequest)))
			}
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func RelayControl(w http.ResponseWriter, r *http.Request) {
	var err error = nil

	if r.Method == "GET" {
		fmt.Println("URL: ", r.URL.Path)
		// check to see if this is the whole collection.
		if matched, _ := regexp.MatchString("^.*relays/?$", r.URL.Path); matched {
			// dump the whole shebang.
			err = json.NewEncoder(w).Encode(relays)
		} else {
			matches := relayURIMatcher.FindStringSubmatch(r.URL.Path)
			if len(matches) == 2 {
				idx, err := strconv.Atoi(matches[1])
				if err == nil {
					err = json.NewEncoder(w).Encode(relays[idx].GetState())
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(http.StatusText(http.StatusBadRequest)))
			}
		}
	} else if r.Method == "PUT" {
		if matches := relayURIMatcher.FindStringSubmatch(r.URL.Path); len(matches) == 2 {
			idx, err := strconv.Atoi(matches[1])
			if err == nil {
				decoder := json.NewDecoder(r.Body)
				var nr pidroponics.Relay
				err := decoder.Decode(&nr)
				if err == nil {
					relays[idx].Device = nr.Device
					relays[idx].Manual = nr.Manual
					err = relays[idx].SetOn(nr.IsOn)
				}
			}
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

// Main state machine control loop.
func controlLoop(tic *time.Ticker) {
	lights := relays[0]
	pump   := relays[1]
	fan    := relays[2]
	valve  := relays[3]

	ambientTemp := thermistors[2]

	for tock := range tic.C {
		hour := tock.Hour()
		minute := tock.Minute()

		if !lights.Manual {
			lights.SetOn(hour > 7 && hour < 22) // Lights on from 7AM to 10 PM.
		}

		// TODO: Calibration
		// TODO: Check proximity sensors to determine if pump should be able to be turned on.
		// TODO: Involve inlet & outlet temperature readings in pump & fan decisions.
		if !pump.Manual {
			pump.SetOn((minute < 15) || (minute > 20 && minute < 35) || (minute > 40 && minute < 55)) // On 15 minutes off 5.
			valve.SetOn(pump.IsOn)
		}

		if !fan.Manual {
			fan.SetOn(ambientTemp.GetState().Temperature > 26) // Anything over 26c (~80f)
		}
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	var err error = nil

	// Setup the SSE Broker.
	broker = pidroponics.NewBroker()
	broker.Start()

	fmt.Println("Enumerating devices...");
	transponderTicker := time.NewTicker(time.Second / 90)
	transponders, err = pidroponics.DetectSrf04(transponderTicker)
	if err != nil {
		log.Fatal("Failed to initialize transponders: ", err)
	}
	for _, transponder := range transponders {
		transponder.AddListener(broker.Outgoing)
	}

	relays, err = pidroponics.DetectRelays()
	if err != nil {
		log.Fatal("Failed to initialize relays: ", err)
	}
	for _, relay := range relays {
		relay.AddListener(broker.Outgoing)
	}

	thermistorTicker := time.NewTicker(time.Second / 10)
	thermistors, err = pidroponics.DetectNTC100KThermistors(thermistorTicker)
	if err != nil {
		log.Fatal("Failed to initialize thermistors: ", err)
	}
	for _, thermistor := range thermistors {
		thermistor.AddListener(broker.Outgoing)
	}

	phTicker := time.NewTicker(time.Second / 3)
	phProbe, err = pidroponics.NewAtlasScientificPhProbe("/sys/bus/platform/drivers/iio_hwmon/pidroponic-hwmon/hwmon/hwmon0/in1_input", phTicker)

	stateTicker := time.NewTicker(time.Second)
	go controlLoop(stateTicker)

	fmt.Println("Setting up HTTP server...")

	handler = http.FileServer(rice.MustFindBox("webroot").HTTPBox())
	fmt.Println("Found the rice box.")

	// Setup the handlers.
	http.HandleFunc("/", RootHandler)

	// Setup the SSE Event Handler. This comes from the 'broker'.
	http.HandleFunc("/events", broker.ServeHTTP)

	http.HandleFunc("/relays", RelayControl)
	http.HandleFunc("/relays/", RelayControl)

	http.HandleFunc("/temperatures", TempControl)
	http.HandleFunc("/temperatures/", TempControl)

	http.HandleFunc("/waterlevels", WaterLevelControl)
	http.HandleFunc("/waterlevels/", WaterLevelControl)

	http.HandleFunc("/ph", PhControl)

	http.Handle("/metrics", promhttp.Handler())

	cert := flag.String("cert", "/etc/ssl/certs/pidroponics.pem", "The certificate for this server.")
	certkey := flag.String("key", "/etc/ssl/certs/pidroponics-key.pem", "The key for the server cert.")

	flag.Parse()

	_, certerr := os.Stat(*cert)
	_, keyerr := os.Stat(*certkey)

	if certerr == nil && keyerr == nil {
		fmt.Println("SSL Configuration set up.")
		go func() {
			log.Fatal(http.ListenAndServe(":80", http.HandlerFunc(redirectTLS)));
		} ()
		log.Fatal(http.ListenAndServeTLS(":443", *cert, *certkey, nil))
	} else {
		fmt.Println("SSL Configuration not found.")
		log.Fatal(http.ListenAndServe(":80", nil))
	}

	return nil
}
