package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
)

func getData(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Hello, Secure World!\n")
}

func main() {
    http.Handle("/", http.FileServer(http.Dir("./static")))
    http.HandleFunc("/data", getData)

    serverEnv := os.Getenv("SERVER_ENV")

    if serverEnv == "DEV" {
        fmt.Println("Running in DEV mode on port 8080...")
        http.ListenAndServe(":8080", nil)
    } else if serverEnv == "PROD" {
        fmt.Println("Running in PROD mode on port 443 with HTTPS...")

        // Set up a secure TLS configuration
        tlsConfig := &tls.Config{
            MinVersion: tls.VersionTLS12, // Minimum supported TLS version is 1.2
            CurvePreferences: []tls.CurveID{
                tls.CurveP256, tls.X25519,
            },
            PreferServerCipherSuites: true,
        }

        server := &http.Server{
            Addr:      ":443",
            TLSConfig: tlsConfig,
        }

        err := server.ListenAndServeTLS(
            "/etc/letsencrypt/live/app.sheribo.site/fullchain.pem",
            "/etc/letsencrypt/live/app.sheribo.site/privkey.pem",
        )
        if err != nil {
            log.Fatalf("Failed to start HTTPS server: %v", err)
        }
    } else {
        fmt.Println("Environment variable SERVER_ENV not set or unrecognized. Exiting...")
    }
}
