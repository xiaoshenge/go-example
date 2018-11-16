package main
import (
 "fmt"
 "github.com/julienschmidt/httprouter"
 "net/http"
 "log"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
 fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
 fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
 router := httprouter.New()
 //router.GET("/", Index)
 //router.GET("/hello/:name", Hello)
 //router.POST("/",Index)
 //router.POST("/aa/:name/:id", Index)
 router.POST("/ab/:id", Index)

 //router.Debug()
 router.POST("/aa", Index)

 router.Debug()

 log.Fatal(http.ListenAndServe(":8081", router))
}