package main

import "fmt"

//memanggil fungsi goodMorning
//fungsi goodMorning akan melakukan print selamat pagi + name yang didapat dari parameter fungsi
func main() {
	goodMorning("teman")
	goodMorning("teman 2")

}
func goodMorning(name string) {
	fmt.Println("selemat pagi", name)
	// TODO: answer here
}
