package wbear

import (

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
  handler := make(map[string]HandlerBear)
  routertest["/user/*/data/pannel"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel2"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel3"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel4"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel5"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel6"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel7"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel8"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel9"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel10"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel11"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel12"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel13"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel14"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel15"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel16"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel17"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel18"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel19"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel20"] = router{handler: handler , params: Paramstest }

  for i := 0; i < b.N; i++ {
      match(routertest,"/user/uuid/data/pannel")
  }
}

func BenchmarkGetKeys( b *testing.B) {
  for i := 0; i < b.N; i++ {
    getKeys("/user/:uuid/data/pannel/:id")
  }

}

