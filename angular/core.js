angular.module('mainApp', ['todoController', 'todoService']);


angular.module('todoController', [])

	// inject the Todo service factory into our controller
	.controller('mainController', ['$scope','$http','Todos', function($scope, $http, Todos) {
		$scope.formData = {};
		$scope.loading = true;

		var getTodo = function() {
			Todos.get()
			.success(function(data) {
				$scope.todos = data;
				$scope.loading = false;
				var task =0;
				for (var i = 0 ; i < $scope.todos.length; i++) {
				if ($scope.todos[i].State == false) {
					task ++;
				}
			}
			$scope.task = task;
			});
		};

		getTodo();
		$scope.createTodo = function() {

			if ($scope.formData.Name != undefined) {
				$scope.loading = true;

				Todos.create($scope.formData)
					.success(function(data) {
						$scope.loading = false;
						$scope.formData = {}; 
						getTodo();
					});

			}
		};
		
		$scope.updateTodo = function(todo) {
			$scope.loading = true;
			dataPut = '{"State":' + todo.State +'}';
			Todos.update(todo.Id,dataPut)
				.success(function(data) {
					$scope.loading = false;

				});
		};

		$scope.clear = function(id, state) {
			$scope.loading = true;
			for (var i = 0 ; i < $scope.todos.length; i++) {
				if ($scope.todos[i].State == true) {
					Todos.delete($scope.todos[i].Id)
					.success(function(data) {
					$scope.loading = false;
					getTodo();

				});
				}
			}	
		};
	}]);

angular.module('todoService', [])

	// super simple service
	// each function returns a promise object 
	.factory('Todos', ['$http',function($http) {
		return {
			get : function() {
				return $http.get('http://192.168.99.100:8081/todos');
			},
			create : function(todoData) {
				return $http.post('http://192.168.99.100:8081/todos', todoData);
			},
			delete : function(id) {
				return $http.delete('http://192.168.99.100:8081/todos/' + id);
			},
			update : function(id, todoData) {
				return $http.put('http://192.168.99.100:8081/todos/' + id, todoData);
			}
		}
	}]);
