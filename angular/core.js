angular.module('mainApp', ['todoController', 'todoService']);


angular.module('todoController', [])

	// inject the Todo service factory into our controller
	.controller('mainController', ['$scope','$http', '$window' ,'Todos', function($scope, $http, $window, Todos) {
		$scope.formData = {};
		$scope.loading = true;
		// each function returns a promise object 
	    var apiBaseUrl = "localhost";
 		var parser = document.createElement('a');
		parser.href = $window.location.href;
		//Only use for docker-compose right now to launch fast
		apiBaseUrl = parser.hostname; 
	    var cookies = document.cookie.split(';');
	    for (var i = 0; i < cookies.length; i++) {
	      var cookie = cookies[i];
	      var trimmedCookie = cookie.match(/^\s*(.*)/)[1];
	      if (trimmedCookie.indexOf("api_override" + '=') === 0) {
	          apiBaseUrl = trimmedCookie.substring("api_override".length + 1, trimmedCookie.length);
	          break;
	      }
	    }

		var getTodo = function() {
			Todos.get(apiBaseUrl)
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

				Todos.create($scope.formData, apiBaseUrl)
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
			Todos.update(todo.ID,dataPut,apiBaseUrl)
				.success(function(data) {
					$scope.loading = false;

				});
		};

		$scope.clear = function(id, state) {
			$scope.loading = true;
			for (var i = 0 ; i < $scope.todos.length; i++) {
				if ($scope.todos[i].State == true) {
					Todos.delete($scope.todos[i].ID,apiBaseUrl)
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
	
			get : function(apiBaseUrl) {
				return $http.get('http://'+apiBaseUrl+':8081/todos');
			},
			create : function(todoData,apiBaseUrl) {
				return $http.post('http://'+apiBaseUrl+':8081/todos', todoData);
			},
			delete : function(id,apiBaseUrl) {
				return $http.delete('http://'+apiBaseUrl+':8081/todos/' + id);
			},
			update : function(id, todoData,apiBaseUrl) {
				return $http.put('http://'+apiBaseUrl+':8081/todos/' + id, todoData);
			}
		}
	}]);
