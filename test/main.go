package main

import (
    "fmt"

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

    b.GET("/user/:name", func(c *wbear.Context)  {
      values := c.Values(c.Request.URL)
      fmt.Println(values)
      fmt.Fprintf(c.Writer,"name of user: %v",values["name"])
    })
    fmt.Println("Run...")
    b.Run(":8080")
}
