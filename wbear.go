package wbear

import (
	"net/http"
	"net/url"
	"strings"
)

type Context struct {
	Request   *http.Request
  Writer    http.ResponseWriter
  *Bear
}



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
func (b *Bear) register(path string,method string, handler HandlerBear, mwf...middleware) { 	
	chpath := changePath(path)
  //check if there is a handler (minimun)
  if len((*b).routes[chpath].handler) != 0 {
    if entry, ok :=  (*b).routes[chpath]; ok {  
	    entry.handler[method] = handler
      entry.middlewares = append(entry.middlewares,mwf... )
    }
  } else {
    params := getKeys(path)
    methodHandler :=  make(map[string]HandlerBear)
    methodHandler[method] = handler 
    router := router{ handler: methodHandler , params: params, middlewares: mwf}
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
 
func (b *Bear) GET(path string, handler HandlerBear, mwf ...middleware) {
  b.register(path,http.MethodGet,handler,mwf...)
}


func (b *Bear) POST(path string, handler HandlerBear, mwf ...middleware) {
  b.register(path,http.MethodPost,handler, mwf...)
}

func (b *Bear) DELETE(path string, handler HandlerBear, mwf ...middleware) {
  b.register(path,http.MethodDelete,handler, mwf...)
}


func (b *Bear) PATCH(path string, handler HandlerBear, mwf ...middleware) {
  b.register(path,http.MethodPatch,handler, mwf...)
}

func (b *Bear) PUT(path string, handler HandlerBear, mwf ...middleware) {
  b.register(path,http.MethodPut,handler, mwf...)
}


func (b *Bear) OPTIONS(path string, handler HandlerBear, mwf ...middleware) {
  b.register(path,http.MethodOptions,handler, mwf...)
}

func (b *Bear) HEAD(path string, handler HandlerBear, mwf ...middleware) {
  b.register(path,http.MethodHead,handler, mwf...)
}




/*
----------------------------------------------------------------
func main() {
    var r Bear
    http.ListenAndServe(":8000", &r)
}
----------------------------------------------------------------
*/

func (b *Bear) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	OriginalPath := match(b.routes,req.URL.Path[1:])
	if OriginalPath == "not found" {
		http.Error(w,http.StatusText(404),404)
		return
	}
  context := &Context{Request: req, Writer: w, Bear: b}
 	r := b.routes[OriginalPath]
  // execute the Middlewares of the single path
  r.execute(context)
  if  _,ok := r.handler[req.Method]; ok{ 
    r.handler[req.Method](b.execute(context))
  } else {
    http.Error(w, "Method http no supported", 401)
  }
}

// initialize the server
func (b *Bear) Run(port string) {
  http.ListenAndServe(port, b)
}
