package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josue/tlc_example/app/api"
)

func main() {

	//gin.SetMode("release")
	r := gin.Default()
	
	api.LoadRoutes(r)


	server := http.Server{
		Addr:    ":8080",
		Handler: r,
		//TLSConfig: tlsConfig,
	}

	//err = server.ListenAndServeTLS("", "")
	log.Println(server.ListenAndServe())

}
