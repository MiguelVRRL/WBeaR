package benchmark

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

    "github.com/MiguelVRRL/WBeaR"
)


func BenchmarkBear(b *testing.B) {
    
    router := WBeaR.NewBear()
    router.Register("/:name",  func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w,"%v", router.Values(r.URL))
    })
    for n := 0; n < b.N; n++ {
        req, err := http.NewRequest("GET", "/Tiago", nil)

        if err != nil {
            b.Fatal(err)
        }
        rr := httptest.NewRecorder()

        router.ServeHTTP(rr, req)
    }
}