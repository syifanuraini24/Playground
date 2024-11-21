package main

import "fmt"

//fungsi printWord akan melakukan print satu persatu nilai parameter yang diterimanya
//contoh nilai parameter yang diterima
//("selamat","pagi","siang",sore)
//outputnya
//selamat
//pagi
//siang
//sore
func main() {
	printWord("selamat")
	printWord("pagi")
	printWord("siang")
	printWord("sore")
}
func printWord(words ...string) {
	result := ""
	for _, words := range words {
		result += "" + words
	}
	fmt.Println(result)
}

// TODO: answer here
