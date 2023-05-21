package WBeaR

import (
	"net/http"
	"net/url"
	"path"
	"regexp"
	"sort"
	"strings"

	"golang.org/x/exp/maps"
)

// represents the parameters into the path to the registry
type Param map[string]string
//ubication of the parameters
type ubicationParam map[int]string

// Param represents a Paht and the router into Bear
type mapRoutes map[string]Router

// type in charge of storing the routes and middlewares
type Bear struct {
	routes mapRoutes
	Middlewares []Middleware
}
// type in charge of Matching the values of the path
var keyPath, _ = regexp.Compile(":[a-zA-Z0-9]{1,}")
// type in charge of Matching the values of the Param

//var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type Router struct {
	path string
	handler http.HandlerFunc
	params Param
	ubication ubicationParam
}

// Router Register a new url
func (b *Bear) Register(path string, handler http.HandlerFunc) { 
	params, ubication := getKeys(path)
	chpath := changePath(path)
	router := Router{path: chpath, handler: handler, params: params, ubication: ubication}
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
	for  i := range pathSplit {
		if keyPath.MatchString(pathSplit[i]) {
			path += "*/" 
		} else if pathSplit[i] != "" {
			path +=  pathSplit[i] + "/"
		}
	} 
	return path
}


// get keys of the path structure
func getKeys(path string) (Param, ubicationParam)  {
	Params := make(Param)
	ubication := make(ubicationParam)
	pathSplit := strings.Split(path,"/")
	numUbication := 0
	for  i := range pathSplit {
		if keyPath.MatchString(pathSplit[i]) {
			Params[pathSplit[i][1:]] = pathSplit[i][1:]
			ubication[numUbication] = pathSplit[i][1:]
			numUbication++
		}
	} 
	return Params, ubication
}

// get values of the path common
func  (b Bear) Values(pathUrl *url.URL)  Param {
	key := Match(b.routes,pathUrl.Path[1:])
	// we use to check if the exits a value
	keysplit := strings.Split(key,"/")
	pathSplit := strings.Split(pathUrl.Path[1:],"/")
	// ubication of the all keys in b.routes[key]
	keysUbication := maps.Keys(b.routes[key].ubication)
	sort.Ints(keysUbication)
	// this incremment the place of the list of keys
	incremment := 0
	for i := 0; i<len(pathSplit);i++ {
		// if field exist in path:
		if keysplit[i] == "*" {
			// this is equal to: Bear.routes[key].params[ubication] == valueOfPath
			b.routes[key].params[b.routes[key].ubication[incremment]] = pathSplit[i]
			incremment++
		}
	}
	return b.routes[key].params

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
