package main

import (
	"fmt"
	"net/http"
	"time"
)

// Dari contoh yang telah diberikan, buatlah http.ServeMux yang memiliki dua route
// Route pertama yaitu "/time" yang menghandle TimeHandler
// Route kedua yaitu "/hello" yang menghandle SayHelloHandler

func TimeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		result := fmt.Sprintf("%v, %v %v %v", t.Weekday(), t.Day(), t.Month(), t.Year())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	} // TODO: replace this
}

func SayHelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name != "" {
			result := fmt.Sprintf("Hello, %v!", name)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(result))
			return
		}

		result := "Hello there"
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	} // TODO: replace this
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	// TODO: answer here
	mux.HandleFunc("/time", TimeHandler())
	mux.HandleFunc("/hello", SayHelloHandler())
	return mux

}
