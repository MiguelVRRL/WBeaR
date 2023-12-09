package wbear


// interface for create a new Middleware
type middleware interface{ 
	ServeContext(c *Context) 
}

// add a new middleware to the Baear
func (b *Bear) UseGlobal(mwf ...middleware) {
	b.Middlewares = append(b.Middlewares, mwf...)
}

func (b *Bear) execute(c *Context) *Context {
	if len(b.Middlewares) == 0 { 
		return c
	}
	for i := 0; i < len(b.Middlewares); i++ {
		b.Middlewares[i].ServeContext(c)
	}
	return c
}

func (r *router) execute(c *Context) (*Context){
	if len(r.middlewares) == 0 { 
		return c
	}
	for i := 0; i < len(r.middlewares); i++ {
		r.middlewares[i].ServeContext(c)
	}
	return c
}


// add a new middleware to a Group
func (g *Group) UseGroup(mwf ...middleware) {
  g.middlewares = append(g.middlewares, mwf...)
}




