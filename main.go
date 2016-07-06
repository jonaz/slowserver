package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	port     string
	response string
	sleep    int
)

func init() {
	flag.StringVar(&port, "p", "8080", "Port to listen on")
	flag.StringVar(&response, "r", "", "response so answer with")
	flag.IntVar(&sleep, "d", 0, "Delay before responding")
}

func main() {
	flag.Parse()
	http.HandleFunc("/", handler)
	log.Println("Starting server on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(sleep) * time.Millisecond)
	fmt.Fprint(w, response)
	log.Printf("| %s | %s | %s", r.RemoteAddr, r.Method, r.RequestURI)
}
