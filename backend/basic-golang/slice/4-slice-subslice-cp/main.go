package main

import "fmt"

// Disini kita akan mencoba untuk melakukan subslicing pada slice.
// Coba langsung gunakan function append ketika melakukan subslicing.
// contoh slice = append (slice, slice2[0:3])

// Silahkan copy slice dan mempunyai value "Marcus", "is", "known", "to", "be" dan "a", "philosopher"
// Silahkan print slice tersebut
// Expected output : [Marcus is known to be a philosopher]
func main() {
	slice := []string{"Marcus", "is", "known", "to", "be", "one", "of", "five", "greatest", "emperors", "of", "rome",
		"Aurelius", "is", "also", "known", "to", "be", "a", "philosopher"}

	// TODO: answer here
	length := len(slice)

	slice1 := slice[0:5]                // 0=Marcus 1=is 2=known 3=to 4=be
	slice2 := slice[length-2 : length]  // 18=a 19=philosopher
	slice3 := append(slice1, slice2...) // 0=Marcus 1=is 2=known 3=to 4=be 5=a 6=philosopher

	fmt.Println("len =", length)
	fmt.Println(slice3)
	apa_ya()
}

func apa_ya() {
	fmt.Print("apa?")
}
