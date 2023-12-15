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
  routertest["/user/*/data/pannel21"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel22"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel23"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel24"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel25"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel26"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel27"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel28"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel29"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel30"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel31"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel32"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel33"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel34"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel35"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel36"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel37"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel38"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel39"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/data/pannel40"] = router{handler: handler , params: Paramstest }


  for i := 0; i < b.N; i++ {
      match(routertest,"/user/uuid/data/pannel")
  }
}

func TestMatchEnd(t *testing.T) {
  Paramstest := make(param)
  routertest := make(mapRoutes)
  handler := make(map[string]HandlerBear)
  routertest["/user/*/data/"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/pannel2/"] = router{handler: handler , params: Paramstest }
  routertest["/user/data/pannel3/"] = router{handler: handler , params: Paramstest }
  routertest["/*/data/pannel4/"] = router{handler: handler , params: Paramstest }
  routertest["/data/*/user/*/pannel5/"] = router{handler: handler , params: Paramstest }
  routertest["/song/pannel/*/"] = router{handler: handler , params: Paramstest }
  routertest["/user/data/"] = router{handler: handler , params: Paramstest }

  assert := "/song/pannel/*/"
  path := "/song/pannel/23"
  key := match(routertest, path)
  if key != assert {
    t.Fatalf("failed match. assert: %s key: %s", assert,key )
  }

}

func TestMatchMiddle(t *testing.T) {
  Paramstest := make(param)
  routertest := make(mapRoutes)
  handler := make(map[string]HandlerBear)
  routertest["/user/*/data/"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/pannel2/"] = router{handler: handler , params: Paramstest }
  routertest["/user/data/pannel3/"] = router{handler: handler , params: Paramstest }
  routertest["/*/data/pannel4/"] = router{handler: handler , params: Paramstest }
  routertest["/data/*/user/*/pannel5/"] = router{handler: handler , params: Paramstest }
  routertest["/song/pannel/*/"] = router{handler: handler , params: Paramstest }
  routertest["/user/data/"] = router{handler: handler , params: Paramstest }

  assert := "/user/*/data/"
  path := "/user/2324/data/"
  key := match(routertest, path)
  if key != assert {
    t.Fatalf("failed match. assert: %s key: %s", assert,key )
  }
}

func TestMatchBegin(t *testing.T) {
  Paramstest := make(param)
  routertest := make(mapRoutes)
  handler := make(map[string]HandlerBear)
  routertest["/user/*/data/"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/pannel2/"] = router{handler: handler , params: Paramstest }
  routertest["/user/data/pannel3/"] = router{handler: handler , params: Paramstest }
  routertest["/*/data/pannel4/"] = router{handler: handler , params: Paramstest }
  routertest["/data/*/user/*/pannel5/"] = router{handler: handler , params: Paramstest }
  routertest["/song/pannel/*/"] = router{handler: handler , params: Paramstest }
  routertest["/user/data/"] = router{handler: handler , params: Paramstest }

  assert := "/*/data/pannel4/"
  path := "/rudel/data/pannel4"
  key := match(routertest, path)
  if key != assert {
    t.Fatalf("failed match. assert: %s key: %s", assert,key )
  }

}


func TestMatchFail(t *testing.T) {
  Paramstest := make(param)
  routertest := make(mapRoutes)
  handler := make(map[string]HandlerBear)
  routertest["/user/*/data/"] = router{handler: handler , params: Paramstest }
  routertest["/user/*/pannel2/"] = router{handler: handler , params: Paramstest }
  routertest["/user/data/pannel3/"] = router{handler: handler , params: Paramstest }
  routertest["/*/data/pannel4/"] = router{handler: handler , params: Paramstest }
  routertest["/data/*/user/*/pannel5/"] = router{handler: handler , params: Paramstest }
  routertest["/song/pannel/*/"] = router{handler: handler , params: Paramstest }
  routertest["/user/data/"] = router{handler: handler , params: Paramstest }

  assert := "not found"
  path := "/data/2334/user/234416g6/pannel3"
  key := match(routertest, path)
  if key != assert {
    t.Fatalf("failed match. assert: %s key: %s", assert,key )
  }

}



func BenchmarkGetKeys( b *testing.B) {
  for i := 0; i < b.N; i++ {
    getKeys("/user/:uuid/data/pannel/:id")
  }

}

