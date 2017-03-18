package main

import (
    "fmt"
    "log"
    "net/http"
    "regexp"
)

var validPath = regexp.MustCompile("^/(users|locations)/([a-zA-Z0-9]*)$")

func mainHandler(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(writer, "Hi there, I love %s!", request.URL.Path[1:])
}

func userHandler(writer http.ResponseWriter, request *http.Request, extras string) {
   fmt.Fprintf(writer, "%s + %s", request.URL.Path[1:], extras)
}

func locationHandler(writer http.ResponseWriter, request *http.Request, extras string) {
   switch request.Method {
      case "GET":
         //not allowed
         log.Printf("%s %s %s", request.Method, request.URL.Path, http.StatusMethodNotAllowed)
         writer.WriteHeader(http.StatusMethodNotAllowed)
      case "POST":
         //
         fmt.Fprintf(writer, "%s + %s inside post", request.URL.Path[1:], extras)
      case "PUT":
         //
         fmt.Fprintf(writer, "%s + %s inside put", request.URL.Path[1:], extras)
      case "DELETE":
         //not allowed
         log.Printf("%s %s %s", request.Method, request.URL.Path, http.StatusMethodNotAllowed)
         writer.WriteHeader(http.StatusMethodNotAllowed)
      default:
         //die
         log.Printf("%s %s %s", request.Method, request.URL.Path, http.StatusBadRequest)
         writer.WriteHeader(http.StatusBadRequest)
   }
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {
        matches := validPath.FindStringSubmatch(request.URL.Path)
        if matches == nil {
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
    
    log.Fatal(http.ListenAndServe(":8080", nil))
}


