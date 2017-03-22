package main

import (
	"net/http"
	"go-git-webhook-client/conf"
	"log"
	"go-git-webhook-client/routers"
	"go-git-webhook-client/controllers"
)

func main()  {


	routers.MapRoute("default","/",&controllers.PayloadController{},routers.Parameter{},"GET")

	routers.RegisterRoutes()

	err := http.ListenAndServe(":" + conf.GetString("httpport","8080"), nil)

	if err != nil {
		log.Fatal(err)
	}
}


