package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// Set the URL of the application running on localhost:8083
	url := "http://localhost:8083"

	// Infinite loop to repeat the process every 30 seconds
	for {
		// Fetch the message from the server
		response, err := http.Get(url)
		if err != nil {
			log.Printf("Error connecting to server: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}
		defer response.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("Error reading response body: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Echo the message from the server to the console
		fmt.Printf("Received message: %s\n", body)

		// Wait for 5 seconds before repeating the process
		time.Sleep(5 * time.Second)
	}
}
