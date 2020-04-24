package main

import (
	"flag"
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	"github.com/bvarner/pidroponics"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

var broker *pidroponics.Broker

var handler http.Handler

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

func main() {
	if err := run(); err != nil {
		fmt.Fprint(os.Stderr, "error: %v\n", err)
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
	for _, file := range files {
		fmt.Println("    ..." + file.Name())
		if strings.HasPrefix(file.Name(), "iio:device") {
			devpath := path.Join("/sys/bus/iio/devices", file.Name())
			devnamebuf, err := ioutil.ReadFile( path.Join(devpath, "name"))
			if err != nil {
				return err
			}

			devname := string(devnamebuf)

			if devname == "srf04" {
				fmt.Println("Ultrasonic transponder at: " + devpath)
			}
			if devname == "ads1015" {
				fmt.Println("ADC at: " + devpath)
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
			fmt.Println("Creating device for: ", path.Join("/sys/class/leds", file.Name()))
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
	//http.HandleFunc("/events", broker.ServeHTTP)

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
