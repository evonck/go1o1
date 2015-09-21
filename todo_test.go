package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var (
	server               *httptest.Server
	reader               io.Reader //Ignore this for now
	todosURL             string
	todosIDURL           string
	todosIDUnexistantURL string
)

func init() {
	server = httptest.NewServer(NewRouter())
	todosURL = fmt.Sprintf("%s/todos", server.URL)
	todosIDURL = fmt.Sprintf("%s/todos/1", server.URL)
	todosIDUnexistantURL = fmt.Sprintf("%s/todos/105", server.URL)
	var envMysqlSetting = os.Getenv("MYSQL_TEST_ENV")
	if strings.EqualFold(envMysqlSetting, "") {
		log.Fatal("No Database environnement set up as MYSQL_TEST_ENV")
	}
	InitDb(envMysqlSetting)
}

/*****************************
        Create
*****************************/
//Test basic creation
func TestCreateTodo(t *testing.T) {
	todoJSON := `{"Name": "Test", "State": true}`

	reader = strings.NewReader(todoJSON)

	request, err := http.NewRequest("POST", todosURL, reader)

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
	todoJSON := `{"Id":1, "Name": "Test", "State": true}`

	reader = strings.NewReader(todoJSON)

	request, err := http.NewRequest("POST", todosURL, reader)

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
	todoJSON := `{Name": "" ,"State": true}`

	reader = strings.NewReader(todoJSON)

	request, err := http.NewRequest("POST", todosURL, reader)

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
	todoJSON := `{"Name": "", "State": true}`

	reader = strings.NewReader(todoJSON)

	request, err := http.NewRequest("POST", todosURL, reader)

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
	todoJSON := `{"Id":1, "Name": "test", "State": true}`

	reader = strings.NewReader(todoJSON)

	request, err := http.NewRequest("POST", todosURL, reader)

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

	request, err := http.NewRequest("GET", todosURL, nil)

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
	todoJSON := `{"Name": "Test1", "State": false}`

	reader = strings.NewReader(todoJSON)

	request, err := http.NewRequest("PUT", todosIDURL, reader)

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
	todoJSON := `{"Id":2 ,"Name": "Test1", "State": false}`

	reader = strings.NewReader(todoJSON)

	request, err := http.NewRequest("PUT", todosIDURL, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

func TestUpdateWrongJson(t *testing.T) {
	todoJSON := `{"Id":2 ,Name": "Test1", "State": false}`

	reader = strings.NewReader(todoJSON)

	request, err := http.NewRequest("PUT", todosIDURL, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 405 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

func TestUpdateUnexistantTodos(t *testing.T) {
	todoJSON := `{"Name": "Test1", "State": false}`

	reader = strings.NewReader(todoJSON)

	request, err := http.NewRequest("PUT", todosIDUnexistantURL, reader)

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

	request, err := http.NewRequest("DELETE", todosIDURL, nil)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

func TestDeleteUnexistantTodos(t *testing.T) {

	request, err := http.NewRequest("DELETE", todosIDUnexistantURL, nil)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 409 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
