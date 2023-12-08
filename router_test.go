package wbear

import (
	"net/http"
	"testing"
)

var routertest mapRoutes
func BenchmarkChangePath(b *testing.B) {
    for i := 0; i < b.N; i++ {
      changePath("/user/:uuid/data/pannel")
    }
}

func BenchmarkMatch( b *testing.B) {
  Paramstest := make(param)
  routertest := make(mapRoutes)
  routertest["/user/*/data/pannel"] = Router{path: "/user/*/data/pannel",handler: func(w http.ResponseWriter,r  *http.Request){}, params: Paramstest }
  for i := 0; i < b.N; i++ {
      Match(routertest,"/user/uuid/data/pannel")
  }
}
