package main

import (
    "fmt"
    "net/http"
    "sync"
)

// Define a global counter and a mutex for thread safety
var requestCount int
var mu sync.Mutex

func main() {
    // Print a message to the console
    fmt.Println("Starting server...")

    // Set up a route for "/"
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Lock the mutex to safely increment the counter
        mu.Lock()
        requestCount++
        mu.Unlock()

        // Return the current request count
        fmt.Fprintf(w, "Hello, World! This is request number %d\n", requestCount)
    })

    // Define the port
    port := 8083
    fmt.Printf("Server running on port %d\n", port)

    // Start the HTTP server
    if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
        fmt.Println("Server error:", err)
    }

    fmt.Println("Server stopped.")
}
