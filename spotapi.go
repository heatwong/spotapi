package main

import (
    "fmt"
    "net/http"
    "regexp"
)

type Spot struct {
   lat float
   lon float
   alt float
}

var validPath = regexp.MustCompile("^/(users|locations)/([a-zA-Z0-9]*)$")

func mainHandler(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(writer, "Hi there, I love %s!", request.URL.Path[1:])
}

func userHandler(writer http.ResponseWriter, request *http.Request, extras string) {
   fmt.Fprintf(writer, "%s + %s", request.URL.Path[1:], extras)
}

func locationHandler(writer http.ResponseWriter, request *http.Request, extras string) {
   fmt.Fprintf(writer, "%s + %s", request.URL.Path[1:], extras)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {
        matches := validPath.FindStringSubmatch(request.URL.Path)
        if m == nil {
            http.NotFound(writer, request)
            return
        }
        fn(writer, request, matches[2])
    }
}

func main() {
    http.HandleFunc("/", mainHandler)
    http.HandleFunc("/users/", makeHandler(userHandler))
    http.HandleFunc("/locations/", makeHandler(locationHandler))
    http.ListenAndServe(":8080", nil)
}


