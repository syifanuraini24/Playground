package main

import (
	"fmt"
)

func searchStudentPresentIndirect(students []string, siswa string, len int) string {
	var absensi string
	if len == -1 {
		absensi = siswa + " tidak hadir"
	} else if students[len] == siswa {
		absensi = siswa + " hadir"
	} else {
		return searchStudentPresentIndirect(students, siswa, len-1)
	}
	return absensi
}

func searchStudentPresent(students []string, siswa string) string {
	len := len(students)
	return searchStudentPresentIndirect(students, siswa, len-1)
}

func main() {
	studentsPresent := []string{"Dino", "Gilang", "Rangga", "Baren", "Dedi", "Dewi", "Juan"}

	//Check if student is present
	var student string
	fmt.Print("Masukkan nama siswa yang ingin dicari : ")
	fmt.Scanln(&student)

	result := searchStudentPresent(studentsPresent, student)
	fmt.Println(result)
}
