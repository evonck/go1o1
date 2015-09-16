# Go 101

Basic Todo rest json using go:
 

 * Get
  ```shell
curl localhost:8080/todos
```
Return the list of Todos 

 * Create
 ```shell
curl -X POST -H "Content-Type: application/json" -d '{"Name":"Test","State":true}' 'http://localhost:8081/todos'
```
Return 201 

 * Update
 ```shell
curl -X PUT -H "Content-Type: application/json" -H " -d '{"Name":"Test","State":true}' 'http://localhost:8081/todos/{todoId}'
```
Return the updated Todo

* Delete
```shell
curl -X DELETE 'http://localhost:8081/todos/1'
```
Return 200
