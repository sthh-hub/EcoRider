package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

func getData(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, Secure World!\n")
}

func main() {
	// Handler for HTTPS requests
	http.HandleFunc("/data", getData)

	// Setting up a TLS configuration with supported protocols
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// Starting the HTTPS server
	server := &http.Server{
		Addr:      ":443",
		TLSConfig: tlsConfig,
	}

	// Redirect HTTP to HTTPS
	go func() {
		log.Println("Starting HTTP to HTTPS redirect server on :80")
		err := http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
		}))
		if err != nil {
			log.Fatalf("HTTP redirect server failed: %v", err)
		}
	}()

	log.Println("Starting HTTPS server on :443")
	err := server.ListenAndServeTLS(
		"/etc/letsencrypt/live/api.sheribo.site/fullchain.pem",
		"/etc/letsencrypt/live/api.sheribo.site/privkey.pem",
	)
	if err != nil {
		log.Fatalf("Failed to start HTTPS server: %v", err)
	}
}