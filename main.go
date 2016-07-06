package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

var (
	port         string
	response     string
	wsdlResponse string
	sleep        int
)

func init() {
	flag.StringVar(&port, "port", "8080", "Port to listen on")
	flag.StringVar(&response, "response", "", "response so answer with")
	flag.StringVar(&wsdlResponse, "wsdl", "", "response so answer on ?WSDL")
	flag.IntVar(&sleep, "delay", 0, "Delay before responding with response")
}

func main() {
	flag.Parse()
	http.HandleFunc("/", handler)
	log.Println("Starting server on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("| %s | %s | %s", r.RemoteAddr, r.Method, r.RequestURI)
	if getWsdl(r.URL) {
		fmt.Fprint(w, wsdlResponse)
		return
	}

	time.Sleep(time.Duration(sleep) * time.Millisecond)
	fmt.Fprint(w, response)
}

func getWsdl(u *url.URL) bool {

	if _, ok := u.Query()["WSDL"]; ok {
		return true
	}
	if _, ok := u.Query()["wsdl"]; ok {
		return true
	}
	return false
}
