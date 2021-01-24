package main

import (
	"net/http"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"encoding/json"
	"fmt"
)

func newHttpServer() {
	router := mux.NewRouter()

	router.HandleFunc("/api/books", handleGetBooks)
	router.HandleFunc("/ready", handleReadinessProbe)

	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%v", configs.Port),
	}

	log.Info(fmt.Sprintf("Dolos listening to port %v", configs.Port))
	log.Fatal(server.ListenAndServe())
}

func handleGetBooks(w http.ResponseWriter, r *http.Request) {

	var books []Book
	
	for i := 0; i <= 10; i++  {
		newBook := Book{}
		gofakeit.Struct(&newBook)
		newBook.PublishDate = gofakeit.Date()
		books = append(books, newBook)
	}

	jbyte, err := json.Marshal(books)
	if err != nil {
		log.Error(err)
		w.Write([]byte("error occurred when retrieving books"))
		return
	}

	// var network bytes.Buffer
	// enc := gob.NewEncoder(&network)
	// encErr := enc.Encode(books)
	// if encErr != nil {
	// 	log.Error(encErr)
	// 	w.Write([]byte("error occurred when retrieving books"))
	// 	return
	// }

	_, werr := w.Write(jbyte)
	if err != nil {
		log.Error(werr)
	}
}

func handleReadinessProbe(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Dolos ready"))
}