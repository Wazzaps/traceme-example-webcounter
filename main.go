package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func main() {
	http.HandleFunc("GET /", handleGet)
	http.HandleFunc("POST /increment", handleIncrement)

	if err := startHttpServer(); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	value := counter
	mutex.Unlock()

	fmt.Fprintf(w, "%d", value)
}

func handleIncrement(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	value := counter
	mutex.Unlock()

	fmt.Fprintf(w, "%d", value)
}

func startHttpServer() error {
	server := http.Server{}
	var listener net.Listener
	var err error
	if envListenUnix := os.Getenv("LISTEN_UNIX"); envListenUnix != "" {
		fmt.Println("Server starting on unix socket", envListenUnix)
		listener, err = net.Listen("unix", envListenUnix)
	} else {
		fmt.Println("Server starting on :8080")
		listener, err = net.Listen("tcp", ":8080")
	}
	if err != nil {
		return err
	}

	return server.Serve(listener)
}
