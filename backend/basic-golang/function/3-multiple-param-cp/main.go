package main

import "fmt"

//Memanggil fungsi goodAfternoon
//dari dalam good afternoon akan dilakukan print "selamat sore name1 dan name2"
func main() {
	goodAfternoon("adi", "anti")
	goodAfternoon("ado", "suci")

}
func goodAfternoon(name, name1 string) {
	fmt.Printf("Selamat sore %s dan %s\n", name, name1)
}

// TODO: answer here
