package main

type EmployeeRow struct {
	ID        int // primary key
	Name      string
	Position  string
	Salary    int
	ManagerID int // foreign key -> Employee
}
type EmployeeDB []EmployeeRow

func NewEmployeeDB() *EmployeeDB {
	return &EmployeeDB{}
}

func (db *EmployeeDB) Where(id int) *EmployeeRow {
	for i := 0; i < len(*db); i++ {
		if (*db)[i].ID == id {
			return &(*db)[i]
		}
	}
	return nil
}

func (db *EmployeeDB) Insert(name string, position string, salary int, managerID int) {
	(*db) = append(*db, EmployeeRow{
		ID:        len(*db) + 1,
		Name:      name,
		Position:  position,
		Salary:    salary,
		ManagerID: managerID,
	})
}

func (db *EmployeeDB) Update(id int, name string, position string, salary int, managerID int) {
	// search employee by id
	employee := db.Where(id)

	// assign new value to properties
	employee.Name = name
	employee.Position = position
	employee.Salary = salary
	employee.ManagerID = managerID
}

func (db *EmployeeDB) Delete(id int) {
	// search employee by id
	for i, employee := range *db {
		if employee.ID == id {
			// remove existing employee with specified id from table
			*db = append((*db)[:i], (*db)[i+1:]...)
			return
		}
	}
}
