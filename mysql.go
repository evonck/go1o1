package main

import (
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

//InitDb initialise the connection with the databases
var InitDb = func(dbUrl string) {
	var err error
	Gdb, err = gorm.Open("mysql", dbUrl)
	if err != nil {
		log.Fatal(err)
	} else {
		Gdb.AutoMigrate(&Todo{})
	}
}
