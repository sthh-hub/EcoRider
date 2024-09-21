package main

import (
	"fmt"
	"net/http"
)

// Handler function for the "/data" path
func getData(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Hello world\n") // Writing response to the client
}

func main() {
    // Register the handler function with the path "/data"
    http.HandleFunc("/data", getData)
    
    // Start the HTTP server on port 8080
    fmt.Println("Server starting on port 8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}