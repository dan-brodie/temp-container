package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"flag"
)

var (
	hostname string
	httpAddr string
)

func webServer(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s - [%s] %s %s %s\n", hostname, time.Now().Format(time.RFC1123), r.RemoteAddr, r.URL, r.Method)
	fmt.Fprintf(w, "Hostname: %s", hostname)
}
// add /healthz endpoint

func main() {
	flag.StringVar(&httpAddr, "http", ":8080", "Port to Listen on")
	flag.Parse()

	var err error
	hostname, err = os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Serving")
	http.HandleFunc("/", webServer)
	log.Fatalln(http.ListenAndServe(httpAddr, nil))
}