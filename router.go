package wbear

import (
	"strings"
)

// represents the parameters into the path to the registry
type param map[string]any//ubication of the parameters
type ubicationParam map[int]string
type HandlerBear func(c *Context)

// Param represents a Paht and the router into Bear
type mapRoutes map[string]router

type router struct {

	handler map[string]HandlerBear
	params param
  middlewares []middleware
}



// match between pathRegister and pathHttp
func match(routes mapRoutes, pathUrl string) string {
  if pathUrl != "" && pathUrl[len(pathUrl)-1] != '/' {
    pathUrl += "/"
  }
  pathSplit := strings.Split(pathUrl, "/")
  for key := range routes {
    keySplit := strings.Split(key, "/")
    if len(keySplit) < len(pathSplit) || len(keySplit) > len(pathSplit) {
      continue
    }
    for index := range keySplit {
      if keySplit[index] != pathSplit[index] && keySplit[index] != "*" {
        break
      } else if index == len(keySplit)  -1 {
        return key
      }
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


