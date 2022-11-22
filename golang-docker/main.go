package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá mundo"))
	})
}

//docker build . -t my-server
//docker run -p 8080:8080 my-server
