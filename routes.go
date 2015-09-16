package main

import (
    "net/http"

)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "Index",
        "GET",
        "/todos",
        TodoShow,
    },
    Route{
        "TodoIndex",
        "POST",
        "/todos",
        TodoCreate,
    },
    Route{
        "TodoShow",
        "PUT",
        "/todos/{todoId}",
        TodoUpdate,
    },
    Route{
        "TodoShow",
        "DELETE",
        "/todos/{todoId}",
        TodoDelete,
    },
}