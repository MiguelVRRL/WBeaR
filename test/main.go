package main

import (
    "fmt"
    "net/http"

    "github.com/MiguelVRRL/wbear"
)

func main() {
  b := wbear.NewBear()
  b.Register("/user/:uuid/",  func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w,"%v", b.Values(r.URL)["uuid"])
  })
   
  v1 := b.Group("/v1")
  v1.Register("/register/:id",  func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w,"%v", b.Values(r.URL)["id"])
  })
  v2 := v1.Group("/v2")
  v2.Register("/register/:uuid",  func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w,"%v", b.Values(r.URL)["uuid"])
  })

  fmt.Println("Run...")
  http.ListenAndServe(":8080", b)
    
}
