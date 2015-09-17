package main

import (
	_ "database/sql"
	_ "evonck/todo/Godeps/_workspace/src/github.com/go-sql-driver/mysql"
	"evonck/todo/Godeps/_workspace/src/github.com/jinzhu/gorm"
	"log"
	"strings"
)

var InitDb func(dbUrl string) = func(dbUrl string) {
	var err error
	Gdb, err = gorm.Open("mysql", dbUrl)
	if err != nil {
		log.Fatal(err)
	} else {
		Gdb.AutoMigrate(&Todo{})
	}
}

//root:@tcp([172.17.0.27]:3306)/todo?charset=utf8&parseTime=True
