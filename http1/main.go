package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

type Person struct {
	Name string `json:"name"`
	Log  string `json:"log"`
}

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

const url = "http://localhost:8080"

func H1CServerUpgrade() {
	fmt.Printf("Listening [0.0.0.0:8081]...\n")
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8081", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	backednResponse := HttpClientExample()
	json, err := json.Marshal(backednResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func HttpClientExample() Person {
	var defaultResponse = Person{}

	client := http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}

	resp, err := client.Get(url)

	if err != nil {
		return defaultResponse
	}

	dec := json.NewDecoder(resp.Body)

	err = dec.Decode(&defaultResponse)

	if err != nil {
		return defaultResponse
	}

	return defaultResponse
}
