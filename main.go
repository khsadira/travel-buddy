package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/projet-x/api"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "The used port for the HTTP server")
	flag.Parse()

	http.HandleFunc("/basic-ask", api.HandleBasicAsk)

	log.Printf("the server is listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatalf("the server has stopped: %v", err)
	}
}
