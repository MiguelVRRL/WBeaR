package wbear

import (
	"fmt"
	"net/http"
	"testing"
)

type HelloApp struct {}
func (h *HelloApp) ServeContext(c *Context) {
    fmt.Println("hello world from a middleware app")
}

func BenchmarkRegister( b *testing.B) {
  bear := NewBear()
  for i := 0; i < b.N; i++ {
    bear.register("/user/:uuid/data/pannel/:id", http.MethodGet,func(c *Context) {},&HelloApp{}) 
  }

}

