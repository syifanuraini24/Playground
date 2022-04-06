package main

import "fmt"

// Disini kalian coba untuk membaca nilai dari slice yang diberikan.
// Lalu kalian akan menambahkan kata "Olleh" pada slice tersebut.
func main() {
	slice := []string{"Hello", "World"}
	slice = append([]string{"Hello"}, slice...)
	fmt.Print(slice)

	// a("Hello", "World", "!")
}

func a(s ...string) {
	for _, v := range s {
		fmt.Println(v)
	}
}
