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


# Launch
## Makefile
 
  * Test
  ```shell
 	make test SRC_ROOT=src MYSQL_TEST_ENV=url
```
Launch a docker container and run the test in SRC_ROOT in it.

Example:
 ```shell
make test SRC_ROOT=/src/evonck/todo 				MYSQL_TEST_ENV=root:@tcp([192.168.99.100]:3306)/test?	charset=utf8&parseTime=True'
```

* Dockerize
```shell
  make dockerize SRC_ROOT=src MYSQL=url
```
  Launch a docker container with the go binary running
  
Example:
```shell
  make test SRC_ROOT=/src/evonck/todo MYSQL=root:@tcp([192.168.99.100]:3306)/todo?charset=utf8&parseTime=True'
```
 * clean
 	
    Clean the folder of the todo binary created in the docker container
    
## Docker-compose
You can launch the entire application using docker by runing:
```shell
  docker-copmose up
```
   This will create a docker go,mysql and nginx, it will run the test install the go binary and launch it. The application will be available on:
   
   $docker-machine ip/8080 : front
   
   $docker-machine ip/8081: api Todo
   
## Front-End
The fornt end of the application is set up to use localhost as the address for the api. You can modify that by setting up a cookie api_override to define the address of your api.

Example:

	document.cookie="api_override=192.168.99.100"
  