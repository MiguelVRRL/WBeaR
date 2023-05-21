package WBeaR

import "net/http"


// interface for create a new Middleware
type Middleware interface{ 
	ServeHTTP(w http.ResponseWriter, r *http.Request) 
}
/*
type middleware interface {
	Middleware(handler http.Handler) http.Handler
}
*/
/*
func (r *Router) Use(mwf ...Middleware) {
	r.middlewares = append(r.middlewares, mwf...)
	
}
*/
// add a new middleware to the Baear
func (b *Bear) UseGlobal(mwf ...Middleware) {
	b.Middlewares = append(b.Middlewares, mwf...)
}

func (b *Bear) Execute(w http.ResponseWriter, r *http.Request)(  http.ResponseWriter,  *http.Request){
	if len(b.Middlewares) == 0 { 
		return w,r
	}
	for i := 0; i < len(b.Middlewares); i++ {
		b.Middlewares[i].ServeHTTP(w,r)
	}
	return w,r
}



