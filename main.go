package main

import (
	"github.com/tasty-road/routers"
)

func main() {

	routersInit := routers.InitRouter()
	routersInit.Run(":8000")
	/*server := &http.Server{
		Addr:           "8000",
		Handler:        routersInit,
		ReadTimeout:    60,
		WriteTimeout:   60,
		//MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", "8000")

	server.ListenAndServe()*/
}
