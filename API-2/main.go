package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")

		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Oioioiiii", http.StatusBadRequest)
			return
		}
		log.Printf("data %s\n", d)
		fmt.Fprintf(w, "Hello %s", d)
	})

	http.HandleFunc("/bye", func(http.ResponseWriter, *http.Request) {
		log.Println("Bye World")
	})

	http.ListenAndServe(":9090", nil)
}
