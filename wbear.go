package wbear

import (
	"net/http"
	"net/url"
	"strings"
)

// type in charge of storing the routes and middlewares
type Bear struct {
	routes mapRoutes
	Middlewares []middleware
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

// Router register a new url
func (b *Bear) register(path string,method string, handler http.HandlerFunc) { 	
	chpath := changePath(path)
  //check if there is a handler (minimun)
  if len((*b).routes[chpath].handler) != 0 {
	   (*b).routes[chpath].handler[method] = handler
  } else {
     params := getKeys(path)
     methodHandler :=  make(map[string]http.HandlerFunc)
     methodHandler[method] = handler 
     router := router{path: chpath, handler: methodHandler , params: params}
	   (*b).routes[chpath]= router
  }

}


// get values of the path common
func  (b Bear) Values(pathUrl *url.URL)  param {
	originalPath := match(b.routes,pathUrl.Path[1:])
	pathSplit := strings.Split(pathUrl.Path[1:],"/")
  paramsCopied := make(param)
  for key, value := range b.routes[originalPath].params  {
		paramsCopied[key] = pathSplit[value.(int)]
	}
  // return a copied because this does'nt change the original values
	return paramsCopied
}


//these functions register the different handlers with their respective http method
func (b *Bear) GET(path string, handler http.HandlerFunc) {
  b.register(path,http.MethodGet,handler)
}


func (b *Bear) POST(path string, handler http.HandlerFunc) {
  b.register(path,http.MethodPost,handler)
}

func (b *Bear) DELETE(path string, handler http.HandlerFunc) {
  b.register(path,http.MethodDelete,handler)
}


func (b *Bear) PATCH(path string, handler http.HandlerFunc) {
  b.register(path,http.MethodPatch,handler)
}

func (b *Bear) PUT(path string, handler http.HandlerFunc) {
  b.register(path,http.MethodPut,handler)
}


func (b *Bear) OPTIONS(path string, handler http.HandlerFunc) {
  b.register(path,http.MethodOptions,handler)
}

func (b *Bear) HEAD(path string, handler http.HandlerFunc) {
  b.register(path,http.MethodHead,handler)
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
	r.handler[req.Method](b.Execute(w, req))

}

func (b *Bear) Run(port string) {
  http.ListenAndServe(":8000", b)
}
