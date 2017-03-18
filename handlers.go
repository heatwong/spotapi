package main

import (
   // "encoding/json"
   "fmt"
   "net/http"
)

func Index(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(writer, "Hi there, I love %s!\n", request.URL.Path[1:])
}

func Users(writer http.ResponseWriter, request *http.Request) {
   fmt.Fprintf(writer, "%s\n", request.URL.Path[1:])
}

func Locations(writer http.ResponseWriter, request *http.Request) {
   fmt.Fprintf(writer, "%s\n", request.URL.Path[1:])
}