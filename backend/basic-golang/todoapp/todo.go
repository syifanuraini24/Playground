package main

import "time"

type Item struct {
	Title    string
	Deadline time.Time
}

type Todos struct {
	items []Item
}

func (todos *Todos) Add(item Item) {
	todos.items = append(todos.items, item)
}

func (todos *Todos) GetAll() []Item {
	return todos.items
}

func (todos *Todos) GetUpcoming() []Item {
	return []Item{todos.items[len(todos.items)-1]}
}

func NewItem(title string, deadline time.Time) Item {
	return Item{title, deadline}
}

func NewTodos() Todos {
	return Todos{}
}
