package main

import (
  "fmt"
  "log"
  "net/http"
)
func whatsUpHandler(w http.ResponseWriter,r *http.Request)  {
  if r.URL.Path != "/whatsUp" {
    http.Error(w, "Status Not found" , http.StatusNotFound)
    return
  }

  if r.Method != "GET" {
    http.Error(w, "Method is not supported" , http.StatusNotFound)
    return
  }

  fmt.Fprintf(w, "Whats up bro")
}

func formHandler(w http.ResponseWriter,r *http.Request)  {
  if err := r.ParseForm(); err != nil{
    fmt.Fprintf(w, "ParseForm() err: %v", err)
    return
  }

  fmt.Fprintf(w, "POST request successful")
  name := r.FormValue("name")
  address := r.FormValue("address")
  town := r.FormValue("town")
  state := r.FormValue("state")
  fmt.Fprintf(w, "Name = %s\n", name)
  fmt.Fprintf(w, "Address = %s\n", address)
  fmt.Fprintf(w, "Town = %s\n", town)
  fmt.Fprintf(w, "Address = %s\n", address)
  fmt.Fprintf(w, "State = %s\n", state)
}


func main(){
  fileServer := http.FileServer(http.Dir("./front-end"))
  http.Handle("/",fileServer)
  http.HandleFunc("/form",formHandler)
  http.HandleFunc("/whatsUp", whatsUpHandler)

  fmt.Printf("starting server at port 7080 \n")
  if err := http.ListenAndServe(":7080", nil); err !=nil{
    log.Fatal(err)
  }
}
