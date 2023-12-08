package main

import (
    "fmt"
    "net/http"

    "github.com/MiguelVRRL/wbear"
)

func main() {
  b := wbear.NewBear()

   
  v1 := b.Group("/v1")
  v1.GET("/register/:id",  func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w,"get %v", b.Values(r.URL)["id"])
  })
  
  v1.POST("/register/:id",  func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w,"post %v", b.Values(r.URL)["id"])
  })
  
  
  fmt.Println("Run...")

  http.ListenAndServe(":8080", b)
    
}
