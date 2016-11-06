package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var config Config

//go build -x github.com/x1m3/xtreamer

func serveVideoAction(w http.ResponseWriter, request *http.Request) {
var video VideoProxy

	vars := mux.Vars(request)
	filename := config.basepath + vars["filename"]
	log.Println("GET " + vars["filename"])
	video.init(filename)
	defer video.close()
	http.ServeContent(w, request, filename, time.Now(),&video)
}

func main() {

	config.load()
	router := mux.NewRouter()
	router.HandleFunc(`/stream/{filename:[a-zA-Z0-9\/\._]+}`, serveVideoAction).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
