package main

import (
   "encoding/json"
   "fmt"
   "io"
   "io/ioutil"
   "net/http"
)

func Index(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(writer, "Hi there, I love %s!\n", request.URL.Path[1:])

    writer.WriteHeader(http.StatusOK)
}

func Users(writer http.ResponseWriter, request *http.Request) {
   fmt.Fprintf(writer, "%s\n", request.URL.Path[1:])

   writer.WriteHeader(http.StatusOK)
}

func SpotsCreate(writer http.ResponseWriter, request *http.Request) {
   fmt.Fprintf(writer, "%s\n", request.URL.Path[1:])

   var spot Spot
   body, err := ioutil.ReadAll(io.LimitReader(request.Body, 1048576))
   if err != nil {
      panic(err)
   }
   if err := request.Body.Close(); err != nil {
      panic(err)
   }
   err = json.Unmarshal(body, &spot)
   if err != nil {
      writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
      writer.WriteHeader(http.StatusUnprocessableEntity)
      if err := json.NewEncoder(writer).Encode(err); err != nil {
         panic(err)
      }
   } else {
      writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
      writer.WriteHeader(http.StatusCreated)
      if err := json.NewEncoder(writer).Encode(spot); err != nil {
         panic(err)
      }
   }
}