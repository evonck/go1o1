package main

import (
    "github.com/jinzhu/gorm"
    _ "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
    "strings"
    "log"
    "os"
)

func getConnectionString() string {
    host := os.Getenv("MYSQL_PORT_3306_TCP_ADDR")
    port :=  "3306"
    user := "root"
    pass :=  ""
    dbname :=  "todo"
    protocol :=  "tcp"
    dbargs := " "

    if strings.Trim(dbargs, " ") != "" {
        dbargs = "?" + dbargs
    } else {
        dbargs = ""
    }
    return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s?charset=utf8&parseTime=True", 
        user, pass, protocol, host, port, dbname)
}

var InitDb func() = func(){
    connectionString := getConnectionString()
    var err error   
    Gdb, err = gorm.Open("mysql", connectionString); 
    if err != nil {
        log.Fatal(err)
    } else {
        Gdb.AutoMigrate(&Todo{})
    }
}