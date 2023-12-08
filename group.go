package wbear

import "net/http"


type Group struct {
  prefix string
  *Bear
}


// returns a subgroup 
func (g *Group) Group(prefix string) *Group {
  return &Group{prefix: g.prefix+prefix, Bear: g.Bear}
}


//these functions register the different handlers with their respective http method and their respective prefixes
func (g *Group) GET(path string, handler http.HandlerFunc) {
  g.Bear.register(g.prefix+path,http.MethodGet,handler)
}


func (g *Group) POST(path string, handler http.HandlerFunc) {
  g.Bear.register(g.prefix+path,http.MethodPost,handler)
}

func (g *Group) DELETE(path string, handler http.HandlerFunc) {
  g.Bear.register(g.prefix+path,http.MethodDelete,handler)
}


func (g *Group) PATCH(path string, handler http.HandlerFunc) {
  g.Bear.register(g.prefix+path,http.MethodPatch,handler)
}

func (g *Group) PUT(path string, handler http.HandlerFunc) {
  g.Bear.register(g.prefix+path,http.MethodPut,handler)
}


func (g *Group) OPTIONS(path string, handler http.HandlerFunc) {
  g.Bear.register(g.prefix+path,http.MethodOptions,handler)
}

func (g *Group) HEAD(path string, handler http.HandlerFunc) {
  g.Bear.register(g.prefix+path,http.MethodHead,handler)
}




