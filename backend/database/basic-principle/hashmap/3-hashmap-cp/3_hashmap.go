package main

import (
	"strings"
)

type PrimaryKey int

type SecondaryKey string

type UserRow struct {
	ID   PrimaryKey   //ID must be unique
	Name SecondaryKey //Name can be non-unique
	Age  int
}

type IndexByID map[PrimaryKey]UserRow

type IndexByName map[SecondaryKey][]PrimaryKey

type UserDB struct {
	ByID   IndexByID
	ByName IndexByName
}

func NewUser() *UserDB {
	return &UserDB{
		ByID:   make(IndexByID),
		ByName: make(IndexByName),
	}
}

func (db *UserDB) Insert(name string, age int) {
	// prepare data
	primaryKey := PrimaryKey(len(db.ByID) + 1)
	secondaryKey := SecondaryKey(name)

	user := &UserRow{
		ID:   primaryKey,
		Name: secondaryKey,
		Age:  age,
	}

	// save data
	// save to ByID
	db.ByID[primaryKey] = *user

	// save to ByName
	db.ByName[secondaryKey] = append(db.ByName[secondaryKey], primaryKey)
}

func (db *UserDB) WhereByID(id PrimaryKey) *UserRow {
	m, ok := db.ByID[id]
	if !ok {
		return nil
	}
	return &m
}

func (db *UserDB) WhereByName(name SecondaryKey) []*UserRow {
	ids := db.ByName[name]
	rows := make([]*UserRow, len(ids))
	for i, id := range ids {
		user := db.WhereByID(id)
		rows[i] = user
	}
	return rows
}

func (db *UserDB) WhereNameBeginsWith(name string) []*UserRow {
	rows := make([]*UserRow, 0)
	for key := range db.ByName {
		if strings.HasPrefix(string(key), name) {
			rows = append(rows, db.WhereByName(key)...)
		}
	}
	return rows
}
