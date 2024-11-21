package main

import "fmt"

func main() {
	// print selamat pagi menggunakan anonymous function
	// fungsi yang lansung dijalankan
	func() {
		fmt.Println("selamat pagi")
	}()
}
