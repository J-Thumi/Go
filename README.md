# Golang 
The Go programming language is an open source project to make programmers more productive.
Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.

Every go project has a go.mod file, created by `go mod init path/to/be/used/when/importing/packages`

`package main`
Every Go file starts with a package declaration.
main is a special package in Go:
Any package named main is treated as an entry point for an executable program.
It must have a main() function — that’s where execution starts (like int main() in C).
Without package main (and main()), Go won’t compile your code into an executable — it would instead compile it into a library.

## Simple HTTP Server

```
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)          // Set handler for root URL
    fmt.Println("Starting server on :8080")
    err := http.ListenAndServe(":8080", nil)  // Start server on port 8080
    if err != nil {
        panic(err)
    }
}
```

`fmt`: Provides formatting functions, here used to write output to the response. More about the package click [here](https://pkg.go.dev/fmt)
`net/http`: Provides HTTP client and server implementations.More about it click [here](https://pkg.go.dev/net/http)

### This is the HTTP handler function.
```
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}
```

It has two parameters:
`w http.ResponseWriter`: An interface used to construct the HTTP response (write headers and body).
`r *http.Request`: Represents the incoming HTTP request (method, URL, headers, body, etc).
Inside, `fmt.Fprintln(w, "Hello, World!")` writes the string "`Hello, World!`" followed by a newline to the response body.
When a client (like your browser) makes a request, this handler sends back "`Hello, World!`".

### Registering the handler
`http.HandleFunc("/", handler)` Tells Go’s HTTP server that for requests to the path / (root URL), it should use the handler function.
You can register different handlers for different URL paths.

### Starting the server

```
fmt.Println("Starting server on :8080")
err := http.ListenAndServe(":8080", nil)
```

`fmt.Println` just prints a message to the terminal.
`http.ListenAndServe(":8080", nil)` starts the HTTP server on port 8080.
The second parameter is the handler to use for requests — `nil` means use the default `http.DefaultServeMux` which we registered the handler to.
This call blocks and runs the server in the foreground, accepting incoming connections.

###  Error handling

```
if err != nil {
    panic(err)
}
```

`ListenAndServe` returns an error if it fails to start or the server crashes.
If there’s an error, the program panics (crashes) and prints the error message.