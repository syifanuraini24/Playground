package main

import (
	"encoding/base64"
	"fmt"
)

// print encode string 'tokenrahasia' kedalam base64
func main() {
	str := "tokenrahasia"
	fmt.Println("1:", str)

	bytes := []byte(str)
	fmt.Println("2:", bytes)

	encodedStr := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println("3:", encodedStr)

	decodedBytes, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		fmt.Println("error decoding string:", err.Error())
		return
	}
	fmt.Println("4:", decodedBytes)

	decodedStr := string(decodedBytes)
	fmt.Println("5:", decodedStr)
}
