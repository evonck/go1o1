package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	Gdb gorm.DB
)

type Db interface {
	Create(value interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Delete(value interface{}, where ...interface{}) *gorm.DB
	AutoMigrate(values ...interface{}) *gorm.DB
}

//Say Hello world :)
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

//Set Header for the front to allow cross platform origin
func SetHeader(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE")
}

//Allow acces to the OPTION url
func AllowAcces(w http.ResponseWriter, r *http.Request) {
	SetHeader(&w)
	w.WriteHeader(http.StatusOK)
}

//Send the list of Todo in the database
func TodoShow(w http.ResponseWriter, r *http.Request) {
	var todos []Todo
	err := Gdb.Find(&todos).Error
	if err != nil {
		panic(err)
	}
	SetHeader(&w)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

//Create a Todo and store it to the database
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	SetHeader(&w)
	todo, err := DecodeJson(r)
	if err != nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Your Json is probably malformed")
		return
	}
	//Check for malforme Json on the Name param
	if strings.EqualFold(todo.Name, "") {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Name parameter can not be empty")
		return
	}
	//Create the todo in the db
	err = Gdb.Create(&todo).Error
	if err != nil {
		//If an error return Conflict
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintln(w, "An error occur :(")
		return
	}
	//Return the new created Todo
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

//Update an existing Todo
func TodoUpdate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	var todoUpdate Todo
	SetHeader(&w)

	vars := mux.Vars(r)
	todoId := vars["todoId"]
	//Find the todo by Id
	err := Gdb.Where("id = ?", todoId).Find(&todo).Error
	if err != nil {
		//If no Todo return Not alowed
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "The id doesn't exist")
		return
	}
	//Find the new Updated json Todo
	todoUpdate, err = DecodeJson(r)
	if err != nil {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Your Json is probably malformed")
		return
	}

	//Check the new todo Name is not null if not update
	if !strings.EqualFold(todoUpdate.Name, "") {
		todo.Name = todoUpdate.Name
	}
	todo.State = todoUpdate.State
	err = Gdb.Save(&todo).Error
	if err != nil {
		//If error during update send back Conflict
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Fatal(err)
		fmt.Fprintln(w, "An error occur on the server")
		return
	}
	//Return OK
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

//Delete a todo
func TodoDelete(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	SetHeader(&w)

	//Find  Todo by Id
	err := Gdb.Where("id = ?", todoId).Find(&todo).Error
	if err != nil {
		//No Id found
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintln(w, "The id doesn't exist")
		return
	}
	//Delete
	err = Gdb.Delete(&todo).Error
	if err != nil {
		//If error send back conflict
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintln(w, "An error occur on the server")
		return
	}
	//If ok return 200
	w.WriteHeader(http.StatusOK)
}

//Function that decode the Json Request to create a Todo object
func DecodeJson(r *http.Request) (Todo, error) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatal(err)
		return todo, err
	}
	if err := r.Body.Close(); err != nil {
		log.Fatal(err)
		return todo, err
	}
	if err := json.Unmarshal(body, &todo); err != nil {

		return todo, err
	}
	return todo, err
}
