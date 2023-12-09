package main

import (
    "fmt"
    "net/http"

    "github.com/MiguelVRRL/wbear"
)

func handler(c *wbear.Context) {
    fmt.Fprintf(c.Writer,"get %v", c.Values(c.Request.URL))
}

type HelloWorld struct {}
func (h *HelloWorld) ServeContext(c *wbear.Context) {
    fmt.Println("hello world from a middleware")
}


func main() {
  b := wbear.NewBear()

  v1 := b.Group("/v1")
  v1.UseGroup(&HelloWorld{})
  v1.GET("/register/:id",handler)  
  fmt.Println("Run...")
  http.ListenAndServe(":8080", b)
    
}
