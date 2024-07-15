// main.go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world! from  %s\n", os.Getenv("HOSTNAME"))
}

func server(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("You received the server which is running on port: %s \n", os.Getenv("PORT"))
}
func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	fmt.Printf("Starting server on port %s...\n", port)

	http.HandleFunc("/server", server)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
