package main

import (
	"fmt"
	"learning/db"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Parameters:
// w http.ResponseWriter → lets you send a response to the client (write HTML, JSON, etc.).
// r *http.Request → contains all request info (method, headers, URL, body, etc.).
// Inside:
// fmt.Fprint(w, "Hello World") writes "Hello World" into the HTTP response body.
// Think of w like a stream going back to the browser.
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r)
}

func main() {
	http.HandleFunc("/", Home)

	db.Connect()
	// Starts the HTTP server on port 80 (default HTTP port).
	// The second argument is the "handler". If nil, it uses Go’s default HTTP multiplexer (the thing that matches routes to functions).
	// This function blocks forever until the server is stopped or crashes.
	http.ListenAndServe(":80", nil)

}
