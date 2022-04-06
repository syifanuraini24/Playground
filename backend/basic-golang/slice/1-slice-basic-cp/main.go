package main

import "fmt"

// Disini teman teman akan mencoba untuk
// melakukan penambahan data pada slice.
// Buatlah variable slice dengan tipe data string.
// Lalu masukan nama kalian ke dalam slice.
// Expected outout: ["NamaPanggilan", "Nama Akhir"]
// Contoh [Zein Fahrozi]
// Outputkan jawabannya ya pastikan cap dan len nya adalah 2
func main() {
	var names []int
	names = append(names, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	cap := cap(names)
	len := len(names)
	fmt.Printf("cap=%d, len=%d, elements=%v", cap, len, names)
}
