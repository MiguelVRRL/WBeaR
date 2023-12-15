# WBeaR.
A web framework minimalist based in Mux.

![alt text](https://github.com/MiguelVRRL/WBeaR/blob/main/logo/WBeaR.svg)

## Why was WBeaR created?.
Principally this project was created for learning more about the functionality of others
frameworks and this proposite this proposal remains valid but i want to improve this project
as you advance in knowledge and can give you a better route, maybe if this goes well
to orient it to a more professional and practical world.

## What does "Bear" offer?.
Although it may seem fun and even weird nothing new really, just as I said before over time I
will be imploring new things but for now it is just the simple project of a student passionate about go.

# Install WbeaR

```
go get -u github.com/MiguelVRRL/wbear
```

# using WBeaR

- ## Basic setup:
```go
package main

import (
    "fmt"
    "net/http"

    "github.com/MiguelVRRL/wbear"
)

func main() {
    b := WBeaR.NewBear()
    b.GET("/user/",  func(c *wbear.Context) {
        fmt.Fprintf(w,"%v", "dummy data :D")
    })
    fmt.Println("Run...")
    b.Run(":8080")
}
```

- ## Dinamics urls:
```go
package main

import (
    "fmt"

    "github.com/MiguelVRRL/wbear"

)

func main() {
    b := WBeaR.NewBear()
    b.GET("/user/:name/", func(c *wbear.Context)  {
         values := c.Values(c.r.URL)
         fmt.Fprintf(w,"name of user: %v",values["name"])
    })
    fmt.Println("Run...")
    b.Run(":8080")
}
```
- ## Middlewares
```go
package main

import (
    "fmt"

    "github.com/MiguelVRRL/wbear"
)

type HelloApp struct {}
func (h *HelloApp) ServeContext(c *wbear.Context) {
    fmt.Println("hello world from a middleware app")
}

type HelloGroup struct {}
func (h *HelloGroup) ServeContext(c *wbear.Context) {
    fmt.Println("hello world from a middleware Group")
}

type HelloPath struct {}
func (h *HelloPath) ServeContext(c *wbear.Context) {
    fmt.Println("hello world from a middleware single path")
}




func main() {
    b := WBeaR.NewBear()
    
    b.GET("/hello-single/",  func(c *Context) {
        fmt.Fprintf(w,"%v", "hello word from the handler")
    }, &HelloPath)

    b.GET("/hello/",  func(c *Context) {
        fmt.Fprintf(w,"%v", "hello word from the handler")
    })
    b.UseGlobal(&HelloApp{})
    
    v1 := b.Group("/v1") 
    v1.UseGroup(&HelloGroup)
    b.GET("/hello/",  func(c *Context) {
        fmt.Fprintf(w,"%v", "hello word from the handler")
    })
    

    fmt.Println("Run...")
    b.Run(":8080")
}
```
- ## Grouping

```go
package main

import (
    "fmt"

    "github.com/MiguelVRRL/wbear"
)

func main() {
  b := wbear.NewBear()
   
  v1 := b.Group("/v1")
  {
      v1.GET("/register/:id",  func(c *wbear.Context) {
      fmt.Fprintf(w,id of the register: "%v", c.Values(c.Request.URL)["id"])
      })
  }

  pannel := v1.Group("pannel")
  pannel.GET("/user/:uuid",  func(c *wbear.Context) {
      fmt.Fprintf(w,user's uuid: uuid: "%v", c.Values(c.Request.URL)["uuid"])
  })

  fmt.Println("Run...")
  b.Run(":8080")
    
}

```
- ## custom http method not suported
If you haven't added a handler for a certain method of a URL, you can respond with a custom html or a default one (included in the framework), saying that it is not supported.

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/MiguelVRRL/wbear"
)

func main() {
    b := WBeaR.NewBear()
    b.GET("/user/",  func(c *wbear.Context) {
        fmt.Fprintf(w,"%v", "dummy data :D")
    })
    wbear.HTTPMethodFailHTML("./template/HTTPMethodFailHTML.html")
    fmt.Println("Run...")
    b.Run(":8080")
}


```
