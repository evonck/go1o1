package main

import (
    "net/http"
    "fmt"
    "github.com/gorilla/mux"
    "encoding/json"
    "github.com/jinzhu/gorm"
    "io"
    "log"
    "strings"
    "io/ioutil"
)

var (
    Gdb gorm.DB
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello World!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    var todos []Todo

    err := Gdb.Find(&todos).Error  
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
        //var txn *gorm.DB = GetTransaction();
        err = Gdb.Create(&todo).Error  
       if err != nil {
            w.Header().Set("Content-Type", "application/json; charset=UTF-8")
            w.WriteHeader(http.StatusConflict)    
        } else {
            w.Header().Set("Content-Type", "application/json; charset=UTF-8")
            w.WriteHeader(http.StatusOK)
            if err := json.NewEncoder(w).Encode(todo); err != nil {
                panic(err)
            }
        }
    }
}

func TodoUpdate(w http.ResponseWriter, r *http.Request) {
    var todo Todo
    var todoUpdate Todo

    vars := mux.Vars(r)
    todoId := vars["todoId"]
    err := Gdb.Where("id = ?", todoId).Find(&todo).Error       
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
        err = Gdb.Update(&todo).Error
        if err != nil {
            w.Header().Set("Content-Type", "application/json; charset=UTF-8")
            w.WriteHeader(http.StatusConflict)
        } else {
            w.Header().Set("Content-Type", "application/json; charset=UTF-8")
            w.WriteHeader(http.StatusOK)
            if err := json.NewEncoder(w).Encode(todo); err != nil {
                panic(err)
            }
        }
    }
}

func TodoDelete(w http.ResponseWriter, r *http.Request) {
    var todo Todo
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    err := Gdb.Where("id = ?", todoId).Find(&todo).Error       
    if err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusConflict)
    }
    err = Gdb.Delete(&todo).Error
    if err != nil {
         w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusConflict)
    } else {
    
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusOK)
    }
}