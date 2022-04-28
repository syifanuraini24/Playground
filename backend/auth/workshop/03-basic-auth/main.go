package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// start the server on port 8000
	fmt.Println("Starting Server at port :8000")
	log.Fatal(http.ListenAndServe(":8000", Routes()))
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	// encrypt token dari username menggunakan bcrypt kemudian kirim ke user kedalam cookie
	mux.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		cookieName := "token"
		var creds Credentials

		// Task:  Buat JSON body diconvert menjadi credential struct & return bad request ketika terjadi kesalahan decoding json

		// TODO: answer here
		body := r.Body
		defer body.Close()
		err := json.NewDecoder(body).Decode(&creds)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Task: Ambil password dari username yang dipakai untuk login

		// TODO: answer here
		actualPassword := users[creds.Username]

		// Task: return unauthorized jika password salah

		// TODO: answer here
		if creds.Password != actualPassword {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Task: 1. Buat token string menggunakan bcrypt dari credential username
		// 		 2. return internal server error ketika terjadi kesalahan encrypting token

		// TODO: answer here
		passwordByte := []byte(creds.Password)
		token, err := bcrypt.GenerateFromPassword(passwordByte, 10)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		tokenString := string(token)

		// Task: Set token string kedalam cookie response

		// TODO: answer here
		expTime := time.Now().Add(30 * time.Minute)
		http.SetCookie(w, &http.Cookie{
			Name:    cookieName,
			Value:   tokenString,
			Expires: expTime,
		})
	})

	return mux
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
