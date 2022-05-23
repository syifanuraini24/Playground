package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type employee struct {
	ID     int
	Name   string
	Salary int
	Role   string
}

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		fmt.Println("error opening database:", err)
	}

	// rows, _ := db.Query("SELECT * FROM employee")
	// var employees []employee = make([]employee, 0)
	// defer rows.Close()
	// for rows.Next() {
	// 	var employeeSatu employee
	// 	rows.Scan(&employeeSatu.ID, &employeeSatu.Name, &employeeSatu.Salary, &employeeSatu.Role)
	// 	fmt.Println(employeeSatu)
	// 	employees = append(employees, employeeSatu)
	// }

	row := db.QueryRow("SELECT * FROM employee WHERE ID = ?", 1)
	var emp employee
	row.Scan(&emp.ID, &emp.Name, &emp.Salary, &emp.Role)

	fmt.Println("You are successfully opening the database.")
	defer db.Close()
}
