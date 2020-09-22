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

const url = "http://l2-envoy"

func H1CServerUpgrade() {
	fmt.Printf("Listening [0.0.0.0:8081]...\n")
	http.HandleFunc("/", FrontendServer)
	http.HandleFunc("/api", ApiServer)
	http.ListenAndServe(":8081", nil)
}

func FrontendServer(w http.ResponseWriter, r *http.Request) {
	backednResponse := HttpClientExample()
	json, err := json.Marshal(backednResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func ApiServer(w http.ResponseWriter, r *http.Request) {
	profile := types.Person{"Http1", os.Getenv("SERVICE_NAME")}

	json, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func HttpClientExample() types.Person {
	var defaultResponse = types.Person{}

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
