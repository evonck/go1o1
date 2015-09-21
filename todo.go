package main

//Todo represent a todo task
type Todo struct {
	ID    int `gorm:"primary_key"`
	Name  string
	State bool
}

//Todos slice of Todo
type Todos []Todo
