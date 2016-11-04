package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//go build -x github.com/x1m3/xtreamer

func serveVideoAction(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "video/mp4")
	w.Write([]byte("hola"))


}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/lala", serveVideoAction).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
