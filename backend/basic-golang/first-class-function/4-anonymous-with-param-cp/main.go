package main

import "fmt"

func main() {
	//fungsi yang mengembalikan nilai pangkat dua dari parameter yang diterima
	//contoh input parameter yang diterima (3)
	//maka fungsi akan mengembalikan 9
	square := func(n int) int {
		return n * n
	}(3)
	// TODO: answer here
	fmt.Println(square)
}
