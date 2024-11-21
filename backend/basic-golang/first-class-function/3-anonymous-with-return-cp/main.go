package main

import "fmt"

func main() {
	// mengembalikan string selamat sore dengan anonymous function
	goodAfternoon := func() string {
		return "selamat sore dari anonymous function"
		// TODO: answer here
	}()

	fmt.Println(goodAfternoon)
}
