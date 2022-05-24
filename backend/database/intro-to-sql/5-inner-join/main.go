package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "../example.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	orderRepo := NewOrderRepository(db)

	orders, err := orderRepo.FetchOrders()
	if err != nil {
		panic(err)
	}

	// Print all orders
	fmt.Printf("ID\tStudentID\tStudentName\tProductID\tProductName\tProductPrice\tQuantity\tOrderDate\tTotalPrice\n")
	for _, o := range orders {
		fmt.Printf("%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n",
			o.ID, o.StudentID, o.StudentName, o.ProductID, o.ProductName, o.ProductPrice, o.Quantity, o.OrderDate.Format("2006-02-01"), o.TotalPrice)
	}
}
