package main

import "fmt"

func main() {
	//fungsi goodMorning melakukan print "selamat pagi"
	// TODO: answer here
	greet := func() {
		fmt.Println("selamat pagi")
	}
	greet()
	fmt.Printf("jenis variabelnya %T", greet)
}
