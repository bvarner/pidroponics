package main

import (
	"encoding/json"
	"flag"
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	"github.com/bvarner/pidroponics"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var broker *pidroponics.Broker

var handler http.Handler

var relays[4]*pidroponics.Relay
var relayMatcher = regexp.MustCompile("^/?relays/([0-3])$")

var transponders[3]*pidroponics.Srf04


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

func RelayControl(w http.ResponseWriter, r *http.Request) {
	var err error = nil

	if r.Method == "GET" {
		fmt.Println("URL: ", r.URL.Path)
		// check to see if this is the whole collection.
		if matched, _ := regexp.MatchString("^.*relays/?$", r.URL.Path); matched {
			// dump the whole shebang.
			err = json.NewEncoder(w).Encode(relays)
		} else {
			matches := relayMatcher.FindStringSubmatch(r.URL.Path)
			if len(matches) == 2 {
				idx, err := strconv.Atoi(matches[1])
				if err == nil {
					err = json.NewEncoder(w).Encode(relays[idx].GetState())
				}
			} else {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("400 - Bad Request"))
			}
		}
	} else if r.Method == "PUT" {
		if matches := relayMatcher.FindStringSubmatch(r.URL.Path); len(matches) == 2 {
			idx, err := strconv.Atoi(matches[1])
			if err == nil {
				decoder := json.NewDecoder(r.Body)
				var nr pidroponics.Relay
				err := decoder.Decode(&nr)
				if err == nil {
					relays[idx].Device = nr.Device
					err = relays[idx].SetOn(nr.IsOn)
				}
			}
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("500 - Method Not Supported"))
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
	}
}


func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Setup the SSE Broker.
	broker = pidroponics.NewBroker()
	broker.Start()

	// Detect IIO Devices
	// /sys/bus/iio/devices/iio:devicen/name
	//    srf04
	//    ads1015
	fmt.Println("Enumerating devices...");
	files, err := ioutil.ReadDir("/sys/bus/iio/devices")
	if err != nil {
		return err
	}

	transponderIdx := 0
	adcIdx := 0

	for _, file := range files {
		if strings.HasPrefix(file.Name(), "iio:device") {
			devpath := path.Join("/sys/bus/iio/devices", file.Name())
			devnamepath := path.Join(devpath, "name")
			devnamebuf, err := ioutil.ReadFile(devnamepath)
			if err != nil {
				return err
			}

			devname := string(devnamebuf)
			// Convert to string so we're working with proper encoding before we drop the last rune.
			devname = devname[:len(devname) - 1]

			if devname == "srf04" {
				fmt.Println("Transponder[", transponderIdx, "] at: ", devpath)
				transponders[transponderIdx], err = pidroponics.NewSrf04(devpath)
				if err == nil {
					transponderIdx++
				}
			}

			if devname == "ads1015" {
				fmt.Println("ADC[", adcIdx, "] at: " + devpath)
				adcIdx++
			}
		}
	}

	// Now that we know how many transponders we have, initialize them with a ticker for polling their state.
	// TODO: Allow for setting / loading maps for devicenames to functions.
//	transponderTicker := time.NewTicker(time.Second / 90)
	transponderTicker := time.NewTicker(time.Second / 90)
	for idx, transponder := range transponders {
		if transponder != nil {
			if idx == 0 {
				err = transponder.Initialize("Sump", transponderTicker, 0)
			}
			if idx == 1 {
				err = transponder.Initialize("Inlet", transponderTicker, 1)
			}
			if idx == 2 {
				err = transponder.Initialize("Outlet", transponderTicker, 2)
			}
			if err != nil {
				transponder.Close()
			}
		}
	}

	// Enumerate Relay Devices. Setup Those.
	// /sys/class/leds/relay0/brightness
	files, err = ioutil.ReadDir("/sys/class/leds")
	if err != nil {
		return err
	}
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "relay") {
			devpath := path.Join("/sys/class/leds", file.Name())

			idx, err := strconv.Atoi(file.Name()[len(file.Name()) - 1:])
			if err != nil {
				log.Fatal("Unable to determine index of :" + file.Name())
			}
			relays[idx], err = pidroponics.NewRelay(devpath, "")
			if err != nil {
				log.Fatal("Unable to initialize: " + devpath, err)
			}
			relays[idx].AddListener(broker.Outgoing)

			// TODO: Allow for setting / loading maps for devicenames to functions.
			// for now, we'll just hard-code them.
			if idx == 0 {
				relays[idx].Device = "Lights"
			}
			if idx == 1 {
				relays[idx].Device = "Pump"
			}
			if idx == 2 {
				relays[idx].Device = "Fan"
			}
			if idx == 3 {
				relays[idx].Device = "Valve"
			}
		}
	}

	// TODO: Load settings that map Devices -> Functional Names.

	// TODO: Setup clock trigger... on clock trigger...

	// TODO: Check current clock time. Compare to desired device states.

	fmt.Println("Setting up HTTP server...")

	handler = http.FileServer(rice.MustFindBox("webroot").HTTPBox())
	fmt.Println("Found the rice box.")

	// Setup the handlers.
	http.HandleFunc("/", RootHandler)

	// Setup the SSE Event Handler. This comes from the 'broker'.
	http.HandleFunc("/events", broker.ServeHTTP)

	http.HandleFunc("/relays", RelayControl)
	http.HandleFunc("/relays/", RelayControl)


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
