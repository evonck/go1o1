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

    todo = DecodeJson(r)

    //Check for malforme Json on the Name param
    if strings.EqualFold(todo.Name, "") {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    } 
    //Create the todo in the db
    err := Gdb.Create(&todo).Error  
    if err != nil {
        //If an error return Conflict
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusConflict)    
        return
    }
    //Return the new created Todo
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(todo); err != nil {
        panic(err)
    }
}

func TodoUpdate(w http.ResponseWriter, r *http.Request) {
    var todo Todo
    var todoUpdate Todo

    vars := mux.Vars(r)
    todoId := vars["todoId"]
    //Find the todo by Id
    err := Gdb.Where("id = ?", todoId).Find(&todo).Error       
    if err != nil {
        //If no Todo return Not alowed
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusMethodNotAllowed) 
        return
    } 
    //Find the new Updated json Todo
    todoUpdate = DecodeJson(r)
    //Check the new todo Name is not null if not update
    //How to deal with empty state as we have a default as false
    //Do we care if send a bad Id on the update json not same as the one on the URL
     if !strings.EqualFold(todo.Name, "") {
        todo.Name = todoUpdate.Name
    } 
    todo.State = todoUpdate.State
    err = Gdb.Save(&todo).Error
    if err != nil {
        //If error during update send back Conflict
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusConflict)
        log.Fatal(err)
        return
    } 
    //Return OK
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(todo); err != nil {
        panic(err)
    }
}

func TodoDelete(w http.ResponseWriter, r *http.Request) {
    var todo Todo
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    //Find  Todo by Id
    err := Gdb.Where("id = ?", todoId).Find(&todo).Error       
    if err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusConflict)
        return
    }
    //Delete
    err = Gdb.Delete(&todo).Error
    if err != nil {
        //If error send back conflict
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusConflict)
        return
    } 
    //If ok return 200
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

//Function that decode the Json Request to create a Todo object
func DecodeJson(r *http.Request) Todo{
    var todo Todo
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &todo); err != nil {
            panic(err)
    }
    return todo
}