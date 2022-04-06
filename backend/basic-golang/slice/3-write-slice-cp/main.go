package main

import "fmt"

// Lanjutan nomor 2
// Sehabis menambahkan "Olleh" pada slice tersebut coba ubah nilai "World" menjadi "Marcus"
// dan "Olleh" menjadi "Aurelius"
func main() {
	slice := []string{"Hello", "World"}
	slice = append(slice, "Olleh")

	for i, v := range slice {
		if v == "World" {
			slice[i] = "Marcus"
		}
		if v == "Olleh" {
			slice[i] = "Aurelius"
		}
	}

	fmt.Print(slice)
}
