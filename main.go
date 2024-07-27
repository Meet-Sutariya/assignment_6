package main

import (
	"fmt"
	"net/http"
	"os"
)

// subjectHandler returns the subject assigned to this VM
func subjectHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the subject from the environment variable
	subject := os.Getenv("SUBJECT")
	if subject == "" {
		subject = "No Subject Assigned"
	}

	// Create the HTML response
	response := fmt.Sprintf("<html><body><h1>Subject: %s</h1></body></html>", subject)

	// Write the response
	fmt.Fprint(w, response)
}

func main() {
	http.HandleFunc("/", subjectHandler)
	fmt.Println("Starting server on :80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		fmt.Printf("Could not start server: %s\n", err)
	}
}
