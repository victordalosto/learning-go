package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Print a message to the console
    fmt.Println("Starting server...")

    // Set up a route for "/"
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, World!")
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
