package main

type Todo struct {
	Id        int 	`gorm:"primary_key"`
	Name 			string 	
	State bool
}


type Todos []Todo
