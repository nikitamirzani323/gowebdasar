package main

import (
	"net/http"

	"golangwebdasar/router"
)

func main() {
	app := router.Init()
	http.ListenAndServe(":7072", app)

}
