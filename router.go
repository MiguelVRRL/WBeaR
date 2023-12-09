package wbear

import (
	"path"
	"strings"

)

// represents the parameters into the path to the registry
type param map[string]any//ubication of the parameters
type ubicationParam map[int]string
type HandlerBear func(c *Context)

// Param represents a Paht and the router into Bear
type mapRoutes map[string]router

type router struct {
	path string
	handler map[string]HandlerBear
	params param
  middlewares []middleware
}


// match between pathRegister and pathHttp
func match(b mapRoutes, pathUrl string) string {
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
func getKeys(path string) (param)  {
	Params := make(param)
	pathSplit := strings.Split(path,"/")
	for  ubication,word := range pathSplit {
		if word != "" && word[0] == 58 {
      Params[word[1:]] = ubication -1 
		}
	} 
	return Params
}


