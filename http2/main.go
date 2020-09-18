package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
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
	H2CServerUpgrade()
}

// This server supports "H2C upgrade" and "H2C prior knowledge" along with
// standard HTTP/2 and HTTP/1.1 that golang natively supports.
func H2CServerUpgrade() {
	h2s := &http2.Server{}

	profile := Person{"Alex", os.Getenv("SERVICE_NAME")}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request coming....")

		json, err := json.Marshal(profile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	})

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: h2c.NewHandler(handler, h2s),
	}

	fmt.Printf("Listening [0.0.0.0:8080]...\n")
	checkErr(server.ListenAndServe(), "while listening")
}

// This server only supports "H2C prior knowledge".
// You can add standard HTTP/2 support by adding a TLS config.
func H2CServerPrior() {
	server := http2.Server{}

	l, err := net.Listen("tcp", "0.0.0.0:8080")
	checkErr(err, "while listening")

	fmt.Printf("Listening [0.0.0.0:8080]...\n")
	for {
		conn, err := l.Accept()
		checkErr(err, "during accept")

		server.ServeConn(conn, &http2.ServeConnOpts{
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Printf("New Connection: %+v\n", r)
				fmt.Fprintf(w, "Hello, %v, http: %v, welcome to service %s", r.URL.Path, r.TLS == nil, os.Getenv("SERVICE_NAME"))
			}),
		})
	}
}
