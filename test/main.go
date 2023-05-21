package main

import (
	"fmt"
	"net/http"

    "github.com/MiguelVRRL/WBeaR"
)

func main() {
	b := WBeaR.NewBear()
    b.Register("/hello/:name/:id/:test",  func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w,"%v", b.Values(r.URL))
    })
	fmt.Println("Run...")
	http.ListenAndServe(":8080", &b)
}
/*
i:= 0
	for i < len(b.routes[key].params){
		for  k,v := range  b.routes[key].params {
			if v == k {
				copyM[k] = pathSplit[k]
				i++
			}
		} 
	}	
*/