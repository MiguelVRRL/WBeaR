package WBeaR

import (
	"net/http"
	"net/url"
	"path"
	//"sort"
	"strings"

	//"golang.org/x/exp/maps"
)

// represents the parameters into the path to the registry
type Param map[string]any//ubication of the parameters
type ubicationParam map[int]string

// Param represents a Paht and the router into Bear
type mapRoutes map[string]Router

// type in charge of storing the routes and middlewares
type Bear struct {
	routes mapRoutes
	Middlewares []Middleware
}


// type in charge of Matching the values of the Param

//var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Router struct {
	path string
	handler http.HandlerFunc
	params Param

	// ubication ubicationParam
}

// Router Register a new url
func (b *Bear) Register(path string, handler http.HandlerFunc) { 
	params := getKeys(path)
	chpath := changePath(path)
	router := Router{path: chpath, handler: handler, params: params}
	(*b).routes[chpath]= router
}
func NewBear() Bear {
	b := new(Bear)
	(*b).routes = make(mapRoutes)
	return *b
}
// match between pathRegister and pathHttp
func Match(b mapRoutes, pathUrl string) string {
	//keys := maps.Keys(b)
	for _,e:= range b {
		if v, err := path.Match(e.path,pathUrl+"/"); err == nil && v {
			return e.path
		} else if v, err := path.Match(e.path,pathUrl); err == nil && v {
			return e.path
		} 
	}
	return "not found"
}


// change path of: /foo/bar/:id to: /foo/bar/*
func changePath(path string) string {
	pathSplit := strings.Split(path,"/")
  path = ""
  for  index,word := range pathSplit {
    if word != "" && word[0] != 58 {
    	path +=  pathSplit[index] + "/"
      continue
    }
    path += "*/"
  }
  return path[2:]
}



// get keys of the path structure
func getKeys(path string) (Param)  {
	Params := make(Param)
	pathSplit := strings.Split(path,"/")
	for  ubication,word := range pathSplit {
		if word != "" && word[0] == 58 {
      Params[word[1:]] = ubication -1 
		}
	} 
	return Params
}

// get values of the path common
func  (b Bear) Values(pathUrl *url.URL)  Param {
	originalPath := Match(b.routes,pathUrl.Path[1:])
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
	OriginalPath := Match(b.routes,req.URL.Path[1:])
	if OriginalPath == "not found" {
		http.Error(w,http.StatusText(404),404)
		return
	}
 	r := b.routes[OriginalPath]
	r.handler(b.Execute(w, req))

}
