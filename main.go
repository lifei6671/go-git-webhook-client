package main

import (
	"net/http"
	"go-git-webhook-client/conf"
	"log"
	"go-git-webhook-client/routers"
	"fmt"
)

func main()  {

	routers.RegisterRoutes()

	host := conf.GetString("httpport","8080")
	fmt.Println("http server Running on http://:",host)
	err := http.ListenAndServe(":" + host, nil)

	if err != nil {
		log.Fatal(err)
	}
}


