package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Halo Dunia"))
	})

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		panic(err)
	}

	log.Println("server started")
}
