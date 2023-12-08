package wbear

import (
	"net/http"
	"net/url"
	"strings"
)



type Group struct {
  prefix string
  *Bear
}

// type in charge of storing the routes and middlewares
type Bear struct {
	routes mapRoutes
	Middlewares []middleware
}

func (g *Group) Group(prefix string) *Group {
  return &Group{prefix: g.prefix+prefix, Bear: g.Bear}
}

func (g *Group) Register(path string, handler http.HandlerFunc) {
  g.Bear.Register(g.prefix+path,handler)
}

// it create a instance for bear
func NewBear() *Bear {
  b := &Bear{}
	(*b).routes = make(mapRoutes)
	return b
}


// create a group
func (b *Bear) Group(prefix string) *Group {
  return &Group{prefix: prefix, Bear: b}
  
}

// Router Register a new url
func (b *Bear) Register(path string, handler http.HandlerFunc) { 	
  params := getKeys(path)
	chpath := changePath(path)
	router := router{path: chpath, handler: handler, params: params}
	(*b).routes[chpath]= router
}



// get values of the path common
func  (b Bear) Values(pathUrl *url.URL)  param {
	originalPath := match(b.routes,pathUrl.Path[1:])
	pathSplit := strings.Split(pathUrl.Path[1:],"/")
	for key, value := range b.routes[originalPath].params {
		// if field exist in path:
			// this is equal to: Bear.routes[key].params[ubication] == valueOfPath
			b.routes[originalPath].params[key] = pathSplit[value.(int)]
	}
	return b.routes[originalPath].params

}

/*
----------------------------------------------------------------
func main() {
    var r Bear
    http.ListenAndServe(":8000", &r)
}
----------------------------------------------------------------
*/

func (b Bear) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	OriginalPath := match(b.routes,req.URL.Path[1:])
	if OriginalPath == "not found" {
		http.Error(w,http.StatusText(404),404)
		return
	}
 	r := b.routes[OriginalPath]
	r.handler(b.Execute(w, req))

}


