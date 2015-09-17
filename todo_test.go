package main

import (
    "fmt"
    "io"
    "net/http/httptest"
    "testing"
    "strings"
    "net/http"
    main "evonck/todo"
    "log"
    "os"

)

var (
    server   *httptest.Server
    reader   io.Reader //Ignore this for now
    todosUrl string
    todosIdUrl string
    todosIdUnexistantUrl string
)

func init() {
    server = httptest.NewServer(main.NewRouter()) 
    todosUrl = fmt.Sprintf("%s/todos", server.URL) 
    todosIdUrl = fmt.Sprintf("%s/todos/1", server.URL) 
    todosIdUnexistantUrl = fmt.Sprintf("%s/todos/105", server.URL) 
    var envMysqlSetting = os.Getenv("MYSQL_TEST_ENV")
    if (strings.EqualFold(envMysqlSetting, "") ){    
        log.Fatal("No Database environnement set up as MYSQL_TEST_ENV")
    }
    main.InitDb(envMysqlSetting)
}

/*****************************
        Create
*****************************/
//Test basic creation
func TestCreateTodo(t *testing.T) {
    todoJson := `{"Name": "Test", "State": true}`

    reader = strings.NewReader(todoJson)

    request, err := http.NewRequest("POST", todosUrl, reader) 

    res, err := http.DefaultClient.Do(request)

    if err != nil {
        t.Error(err) 
    }
    if res.StatusCode != 201 {
        t.Errorf("Success expected: %d", res.StatusCode) 
    }
}

//Test with creation with Id

func TestCreateWithId(t *testing.T) {
    todoJson := `{"Id":1, "Name": "Test", "State": true}`

    reader = strings.NewReader(todoJson)

    request, err := http.NewRequest("POST", todosUrl, reader) 

    res, err := http.DefaultClient.Do(request)

    if err != nil {
        t.Error(err) 
    }
    if res.StatusCode != 201 {
        t.Errorf("Success expected: %d", res.StatusCode) 
    }
}

//Test creation with bad Json
func TestCreateBadJson(t *testing.T) {
    todoJson := `{Name": "" ,"State": true}`

    reader = strings.NewReader(todoJson)

    request, err := http.NewRequest("POST", todosUrl, reader) 

    res, err := http.DefaultClient.Do(request)

    if err != nil {
        t.Error(err) 
    }
    if res.StatusCode != 405 {
        t.Errorf("Success expected: %d", res.StatusCode) 
    }
}

//Test with creation of empty Name

func TestCreateEmptyName(t *testing.T) {
    todoJson := `{"Name": "", "State": true}`

    reader = strings.NewReader(todoJson)

    request, err := http.NewRequest("POST", todosUrl, reader) 

    res, err := http.DefaultClient.Do(request)

    if err != nil {
        t.Error(err) 
    }
    if res.StatusCode != 405 {
        t.Errorf("Success expected: %d", res.StatusCode) 
    }
}

//Test with creation with Id already exist

func TestCreateWhitExistingId(t *testing.T) {
    todoJson := `{"Id":1, "Name": "test", "State": true}`

    reader = strings.NewReader(todoJson)

    request, err := http.NewRequest("POST", todosUrl, reader) 

    res, err := http.DefaultClient.Do(request)

    if err != nil {
        t.Error(err) 
    }
    if res.StatusCode != 409 {
        t.Errorf("Success expected: %d", res.StatusCode) 
    }
}


/*****************************
        Get
*****************************/

func TestGetTodos(t *testing.T) {

    request, err := http.NewRequest("GET", todosUrl, nil) 

    res, err := http.DefaultClient.Do(request)

    if err != nil {
        t.Error(err) 
    }
    if res.StatusCode != 200 {
        t.Errorf("Success expected: %d", res.StatusCode) 
    }
}

/*****************************
        Update
*****************************/

func TestUpdateTodos(t *testing.T) {
    todoJson := `{"Name": "Test1", "State": false}`

    reader = strings.NewReader(todoJson)

    request, err := http.NewRequest("PUT", todosIdUrl, reader) 

    res, err := http.DefaultClient.Do(request)

    if err != nil {
        t.Error(err) 
    }
    if res.StatusCode != 200 {
        t.Errorf("Success expected: %d", res.StatusCode) 
    }
}


//Update Url Id 1 but json Id 2 
func TestUpdateWrongIdTodos(t *testing.T) {
    todoJson := `{"Id":2 ,"Name": "Test1", "State": false}`

    reader = strings.NewReader(todoJson)

    request, err := http.NewRequest("PUT", todosIdUrl, reader) 

    res, err := http.DefaultClient.Do(request)

    if err != nil {
        t.Error(err) 
    }
    if res.StatusCode != 200 {
        t.Errorf("Success expected: %d", res.StatusCode) 
    }
}

func TestUpdateWrongJson(t *testing.T) {
    todoJson := `{"Id":2 ,Name": "Test1", "State": false}`

    reader = strings.NewReader(todoJson)

    request, err := http.NewRequest("PUT", todosIdUrl, reader) 

    res, err := http.DefaultClient.Do(request)

    if err != nil {
        t.Error(err) 
    }
    if res.StatusCode != 405 {
        t.Errorf("Success expected: %d", res.StatusCode) 
    }
}

func TestUpdateUnexistantTodos(t *testing.T) {
    todoJson := `{"Name": "Test1", "State": false}`

    reader = strings.NewReader(todoJson)

    request, err := http.NewRequest("PUT", todosIdUnexistantUrl, reader) 

    res, err := http.DefaultClient.Do(request)

    if err != nil {
        t.Error(err) 
    }
    if res.StatusCode != 405 {
        t.Errorf("Success expected: %d", res.StatusCode) 
    }
}

/*****************************
        Delete
*****************************/

func TestDeleteTodos(t *testing.T) {

    request, err := http.NewRequest("DELETE", todosIdUrl, nil) 

    res, err := http.DefaultClient.Do(request)

    if err != nil {
        t.Error(err) 
    }
    if res.StatusCode != 200 {
        t.Errorf("Success expected: %d", res.StatusCode) 
    }
}

func TestDeleteUnexistantTodos(t *testing.T) {

    request, err := http.NewRequest("DELETE", todosIdUrl, nil) 

    res, err := http.DefaultClient.Do(request)

    if err != nil {
        t.Error(err) 
    }
    if res.StatusCode != 409 {
        t.Errorf("Success expected: %d", res.StatusCode) 
    }
}