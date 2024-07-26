package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello my name is Aman")
	})

	log.Println("Server is running on port 8100")
	http.ListenAndServe("[::]:8100", nil)
}
