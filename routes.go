package main

import "net/http"

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
        "Users",
        "GET",
        "/users",
        Users,
    },
    Route{
        "Locations",
        "POST",
        "/locations",
        Locations,
    },
}