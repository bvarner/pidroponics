package main

import (
	"flag"
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	"github.com/bvarner/pidroponics"
	"log"
	"net/http"
	"os"
	"time"
)
var handler http.Handler

var inletWaterLevel *pidroponics.HCSR04

var drainWaterLevel *pidroponics.HCSR04

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
	var err error = nil

	inletWaterLevel, err = pidroponics.NewHCSR04("GPIO5", "GPIO6")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Inlet Water Level Not Initialized: ", err);
	}

	drainWaterLevel, err = pidroponics.NewHCSR04("GPIO16", "GPIO17")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Drain Water Level Not Initialized: ", err);
	}

	poller := time.NewTicker(1 * time.Second)
	go func() {
		for range poller.C {
			inletWaterLevel.MeasureDistance()
			fmt.Println("Inlet Wather Level: ", inletWaterLevel.Distance)
//			drainWaterLevel.MeasureDistance()
//			fmt.Println("Drain Wather Level: ", drainWaterLevel.Distance)
		}
	}()

	fmt.Println("Setting up HTTP server...")

	handler = http.FileServer(rice.MustFindBox("webroot").HTTPBox())
	fmt.Println("Found the rice box.")

	// Setup the handlers.
	http.HandleFunc("/", RootHandler)

	// Setup the SSE Event Handler. This comes from the 'broker'.
	//http.HandleFunc("/events", broker.ServeHTTP)

	cert := flag.String("cert", "/etc/ssl/certs/pidroponics.pem", "The certificate for this server.")
	certkey := flag.String("key", "/etc/ssl/certs/pidroponics.pem", "The key for the server cert.")

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
}