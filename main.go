package main

import (
	"net/http"
	"go-git-webhook-client/conf"
	"log"
	"go-git-webhook-client/routers"
)

func main()  {

	routers.RegisterRoutes()

	err := http.ListenAndServe(":" + conf.GetString("httpport","8080"), nil)

	if err != nil {
		log.Fatal(err)
	}
}


