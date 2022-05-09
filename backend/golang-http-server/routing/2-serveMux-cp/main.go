package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Dari contoh yang telah diberikan, buatlah route untuk TimeHandler dan SayHelloHandler.
// Buatlah route "/time" pada TimeHandler dan "/hello" untuk SayHelloHandler dengan menggunakan ServeMux

var TimeHandler = func(writer http.ResponseWriter, request *http.Request) {
	// TODO: answer here
	t := time.Now()
	weekday := t.Weekday()
	day := t.Day()
	month := t.Month()
	year := t.Year()

	result := fmt.Sprintf("%v, %v %v %v", weekday, day, month, year)

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(result))
}

var SayHelloHandler = func(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	result := "Hello there"
	name := r.URL.Query().Get("name")
	if name != "" {
		result = fmt.Sprintf("Hello, %v!", name)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}

func main() {
	mux := http.NewServeMux()
	// TODO: answer here
	mux.HandleFunc("/time", TimeHandler)
	mux.HandleFunc("/hello", SayHelloHandler)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
