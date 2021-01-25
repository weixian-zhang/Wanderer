package main

import (
	"io/ioutil"
	"net/http"

	//"path"
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func newHttpServer() {
	router := mux.NewRouter()

	router.HandleFunc("/api/store/books", handleGetBooks)
	router.HandleFunc("/ready", handleReady)

	server := &http.Server{
		Handler: router,
		Addr: fmt.Sprintf(":%v", configs.Port),
	}

	log.Fatal(server.ListenAndServe())
}

func handleGetBooks(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get(configs.BookApiBaseUrl + "/api/books")
	if LogIfError(err) {
		w.Write([]byte(fmt.Sprintf("Unable to contact Dolos service for books. \r\n%v", err.Error())))
		return
	}

	defer resp.Body.Close()

	barr, ioerr := ioutil.ReadAll(resp.Body)
	LogIfError(ioerr)

	books := &[]Book{}
	json.Unmarshal(barr, books)

	_, werr := w.Write(barr)
	LogIfError(werr)
}

func handleReady(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ready"))
}