package main

import (
	"log"
	"net/http"
	"parallel-programming/assignment3"
	"time"
)

// Import the package that contains the server and client code

func main() {
	// Run the server in a separate goroutine
	go func() {
		mux := http.NewServeMux()

		// Assuming you have defined your handlers in srv.go
		mux.HandleFunc("/kvs/", assignment3.KVSPut)
		mux.HandleFunc("/kvs/", assignment3.KVSGet)
		mux.HandleFunc("/kvs/", assignment3.KVSDel)

		log.Println("Starting server on :3000")
		if err := http.ListenAndServe(":3000", mux); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait a moment for the server to start
	time.Sleep(1 * time.Second)

	// Now run your client operations, assuming they are defined in Client.go
	assignment3.MainClient()
}
