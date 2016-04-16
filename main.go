package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/sir-wiggles/auth/views"
)

func main() {
	ssl := flag.Bool("ssl", false, "flag enable ssl")
	port := flag.String("port", "3000", "port to run server on")
	keyPem := flag.String("key", "key.pem", "location of your key.pem file")
	certPem := flag.String("cert", "cert.pem", "location of your cert.pem file")
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/login", api.LoginHandler)

	router := handlers.LoggingHandler(os.Stdout, r)
	log.Printf("Serving on port :%s with ssl %t", *port, *ssl)

	if *ssl {
		http.ListenAndServeTLS(fmt.Sprintf(":%s", *port), *certPem, *keyPem, router)
	} else {
		http.ListenAndServe(fmt.Sprintf(":%s", *port), router)
	}
}
