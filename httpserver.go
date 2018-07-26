package main

import (
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!!!\n")
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Path)
}

func HttpServ() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/", echoHandler)

	http.ListenAndServe(":12345", mux)

}

func FileServ() {
	http.ListenAndServe(":12346", http.FileServer(http.Dir("./storage")))
}
