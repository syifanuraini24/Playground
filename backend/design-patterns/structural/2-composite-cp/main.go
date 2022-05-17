package main

import (
	"fmt"

	"github.com/ruang-guru/playground/backend/design-patterns/structural/2-composite-cp/perusahaan"
)

// CTO -> { VP1 -> { J1, J2 }, VP2 -> { J3 } }

func main() {
	j1 := perusahaan.Junior{}
	j2 := perusahaan.Junior{}
	j3 := perusahaan.Junior{}

	vp1 := perusahaan.VP{
		Subordinate: []perusahaan.Employee{j1, j2},
	}

	vp2 := perusahaan.VP{
		Subordinate: []perusahaan.Employee{j3},
	}

	cto := perusahaan.CTO{
		Subordinate: []perusahaan.Employee{vp1, vp2},
	}

	fmt.Printf("Total divison CEO salary: %d\n", cto.TotalDivisonSalary())  // 30 + 40 + 30 = 100
	fmt.Printf("Total divison VP-1 salary: %d\n", vp1.TotalDivisonSalary()) // 20 + 10 + 10 = 40
	fmt.Printf("Total divison VP-2 salary: %d\n", vp2.TotalDivisonSalary()) // 20 + 10 = 30
}
