package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Everything is fine %q", r.URL.Path[1:])
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
