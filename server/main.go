package main

import (
	"fmt"
	"net/http"
	"os"
)

func getData(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Hello world\n")
}

func main() {
    http.HandleFunc("/data", getData)

    // Retrieve the server environment variable
    serverEnv := os.Getenv("SERVER_ENV")

    // Run the server based on the environment
    if serverEnv == "DEV" {
        fmt.Println("Running in DEV mode on port 8080...")
        http.ListenAndServe(":8080", nil) // Local development
    } else if serverEnv == "PROD" {
        fmt.Println("Running in PROD mode on port 443 with HTTPS...")
        err := http.ListenAndServeTLS(
            ":443",
            "/etc/letsencrypt/live/api.sheribo.site/fullchain.pem",
            "/etc/letsencrypt/live/api.sheribo.site/privkey.pem",
            nil,
        )
        if err != nil {
            fmt.Printf("Failed to start HTTPS server: %v\n", err)
        }
    } else {
        fmt.Println("Environment variable SERVER_ENV not set or unrecognized. Exiting...")
    }
}