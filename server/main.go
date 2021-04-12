package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"totalAmount":{"value":517.9}}`))
	})

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
