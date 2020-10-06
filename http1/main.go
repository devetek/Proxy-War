package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/devetek/error-hanlder/types"
	"golang.org/x/net/http2"
)

func checkErr(err error, msg string) {
	if err == nil {
		return
	}
	fmt.Printf("ERROR: %s: %s\n", msg, err)
	os.Exit(1)
}

func main() {
	H1CServerUpgrade()
}

var URL_HTTP1 = os.Getenv("BACKEND_API_HTTP1")
var URL_HTTP2 = os.Getenv("BACKEND_API_HTTP2")

func H1CServerUpgrade() {
	fmt.Printf("Listening [0.0.0.0:8081]...\n")
	http.HandleFunc("/", FrontendServer)
	http.HandleFunc("/api", ApiServer)
	http.ListenAndServe(":8081", nil)
}

func FrontendServer(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request coming to frontend HTTP1....\n")
	backednResponse := OpenConnectionHTTP2()
	json, err := json.Marshal(backednResponse)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Service-Name", "http1-frontend")
	w.Write(json)
}

func ApiServer(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request coming to backend HTTP1....\n")
	profile := types.Person{"Http1", os.Getenv("SERVICE_NAME")}

	json, err := json.Marshal(profile)
	if err != nil {
		fmt.Printf("ApiServer Error 1: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func OpenConnectionHTTP2() types.Person {
	var defaultResponse = types.Person{}

	client := http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}

	resp, err := client.Get(URL_HTTP2)

	if err != nil {
		fmt.Printf("OpenConnectionHTTP2 Error 1: %v", err)
		return defaultResponse
	}

	dec := json.NewDecoder(resp.Body)

	err = dec.Decode(&defaultResponse)

	if err != nil {
		fmt.Printf("OpenConnectionHTTP2 Error 2: %v", err)
		return defaultResponse
	}

	return defaultResponse
}
