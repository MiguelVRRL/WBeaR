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

    "github.com/MiguelVRRL/WBeaR"
)

func main() {
	b := WBeaR.NewBear()
    b.Register("/user/",  func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w,"%v", "dummy data :D")
    })
	fmt.Println("Run...")
	http.ListenAndServe(":8080", &b)
}
```

- ## Dinamics urls:
```go
package main

import (
	"fmt"
	"net/http"

    "github.com/MiguelVRRL/WBeaR"
)

func main() {
	b := WBeaR.NewBear()
    b.Register("/iser/:name/",  func(w http.ResponseWriter, r *http.Request) {
         values := b.Values(r.URL)
        fmt.Fprintf(w,"name of user: %v",values["name"])
    })
	fmt.Println("Run...")
	http.ListenAndServe(":8080", &b)
}
```
- ## Middlewares
```go
package main

import (
	"fmt"
	"net/http"

    "github.com/MiguelVRRL/WBeaR"
)

type HelloWorld struct {}
func (h *HelloWorld) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world from a middleware")
}

func main() {
	b := WBeaR.NewBear()
    b.Register("/hello/",  func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w,"%v", "hello word from the handler")
    })
	b.UseGlobal(&HelloWorld{})
	fmt.Println("Run...")
	http.ListenAndServe(":8080", &b)
}
```
