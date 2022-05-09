package main

import (
	"fmt"
	"net/http"
	"time"
)

// Dari contoh yang diberikan, buatlah sebuah handler dengan menggunakan HandlerFunc yang menampilkan nama hari, bulan, dan tahun.
// Hint, gunakan time.Weekday, time.Day, time.Month, dan time.Year

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		t := time.Now()
		weekday := t.Weekday()
		day := t.Day()
		month := t.Month()
		year := t.Year()

		result := fmt.Sprintf("%v, %v %v %v", weekday, day, month, year)

		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(result))
	}
}
