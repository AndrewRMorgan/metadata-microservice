package main

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := httprouter.New()
	router.GET("/", index)
	http.ListenAndServe(":"+port, router)
}

func index(res http.ResponseWriter, req *http.Request, ps httprouter.Params) {

}
