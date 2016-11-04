package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

//go build -x github.com/x1m3/xtreamer

func serveVideoAction(w http.ResponseWriter, request *http.Request) {
	basepath:=`C:\Users\xime\Desktop\`
	vars := mux.Vars(request)
	filename := basepath + vars["filename"]
	fp,_ := os.Open(filename)
	http.ServeContent(w, request, filename, time.Now(),fp)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc(`/stream/{filename:[a-zA-Z0-9\/\._]+}`, serveVideoAction).Methods("GET")



	log.Fatal(http.ListenAndServe(":8000", router))
}
