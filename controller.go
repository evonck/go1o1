package main

import (
    "net/http"
    "fmt"
    "github.com/gorilla/mux"
    "encoding/json"
    "github.com/jinzhu/gorm"
    "database/sql"
    "io"
    "log"
    "strings"
    "io/ioutil"
)

var (
    Gdb gorm.DB
)

func GetTransaction() *gorm.DB {
    txn := Gdb.Begin()
        if txn.Error != nil {
            panic(txn.Error)
        }
        return txn
}

func CommitTransaction(txn *gorm.DB) bool{
    if txn == nil {
        return false
    }
    txn.Commit()
    if err := txn.Error; err != nil && err != sql.ErrTxDone {
        Rollback(txn)
        return false
    }
    return true
}

func Rollback(txn *gorm.DB) bool{
   if txn == nil {
        return false
    }
    txn.Rollback()
    if err := txn.Error; err != nil && err != sql.ErrTxDone {
        return false
    }
    txn = nil
    return true
}


func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello World!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    var txn *gorm.DB = GetTransaction();
    var todos []Todo

    err := txn.Find(&todos).Error  
    if err != nil {
        panic(err)
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
    var todo Todo
    vars := mux.Vars(r)
    log.Print(vars)

    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    log.Print(body)

    if err := json.Unmarshal(body, &todo); err != nil {
            log.Print(todo)
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }
    if strings.EqualFold(todo.Name, "") {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusMethodNotAllowed)
    } else {
        var txn *gorm.DB = GetTransaction();
        txn.Create(&todo)
        succes := CommitTransaction(txn)
        if (succes) {
            w.Header().Set("Content-Type", "application/json; charset=UTF-8")
            w.WriteHeader(http.StatusOK)
            if err := json.NewEncoder(w).Encode(todo); err != nil {
                panic(err)
            }
        } else {
            w.Header().Set("Content-Type", "application/json; charset=UTF-8")
            w.WriteHeader(http.StatusConflict)
        }
    }
}

func TodoUpdate(w http.ResponseWriter, r *http.Request) {
    var todo Todo
    var todoUpdate Todo

    vars := mux.Vars(r)
    todoId := vars["todoId"]
    var txn *gorm.DB = GetTransaction();
    err := txn.Where("id = ?", todoId).Find(&todo).Error       
    if err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusMethodNotAllowed) 
    } else {
        body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
        if err != nil {
            panic(err)
        }
        if err := r.Body.Close(); err != nil {
            panic(err)
        }
        if err := json.Unmarshal(body, &todoUpdate); err != nil {
            if err := json.NewEncoder(w).Encode(err); err != nil {
                panic(err)
            }
        }
        todo.Name = todoUpdate.Name
        todo.State = todoUpdate.State
        txn.Update(&todo)
        CommitTransaction(txn)
        succes := CommitTransaction(txn)
        if (succes) {
            w.Header().Set("Content-Type", "application/json; charset=UTF-8")
            w.WriteHeader(http.StatusOK)
            if err := json.NewEncoder(w).Encode(todo); err != nil {
                panic(err)
            }
        } else {
            w.Header().Set("Content-Type", "application/json; charset=UTF-8")
            w.WriteHeader(http.StatusConflict)
        }
    }
}

func TodoDelete(w http.ResponseWriter, r *http.Request) {
    var todo Todo
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    var txn *gorm.DB = GetTransaction();
    err := txn.Where("id = ?", todoId).Find(&todo).Error       
    if err != nil {
         w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusConflict)
    }
    txn.Delete(&todo)
    succes := CommitTransaction(txn)
    if (succes) {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)
    } else {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusConflict)
    }
}