// frontend/main.go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://backend:3000/server")
	if err != nil {
		http.Error(w, "Error contacting backend", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading backend response", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Frontend received: %s", body)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5000", nil)
}
