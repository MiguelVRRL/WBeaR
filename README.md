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
    b.GET("/user/",  func(w http.ResponseWriter, r *http.Request) {
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
    "net/http"

    "github.com/MiguelVRRL/wbear"

)

func main() {
    b := WBeaR.NewBear()
    b.GET("/user/:name/",  func(w http.ResponseWriter, r *http.Request) {
         values := b.Values(r.URL)
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
    "net/http"

    "github.com/MiguelVRRL/wbear"
)

type HelloWorld struct {}
func (h *HelloWorld) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Println("hello world from a middleware")
}

func main() {
    b := WBeaR.NewBear()
    b.GET("/hello/",  func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w,"%v", "hello word from the handler")
    })
    b.UseGlobal(&HelloWorld{})
    fmt.Println("Run...")
    b.Run(":8080")
}
```
- ## Groping

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/MiguelVRRL/wbear"
)

func main() {
  b := wbear.NewBear()
   
  v1 := b.Group("/v1")
  {
      v1.GET("/register/:id",  func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w,id of the register: "%v", b.Values(r.URL)["id"])
      })
  }

  pannel := v1.Group("pannel")
  pannel.GET("/user/:uuid",  func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w,user's uuid: uuid: "%v", b.Values(r.URL)["uuid"])
  })

  fmt.Println("Run...")
  b.Run(":8080")
    
}

```
