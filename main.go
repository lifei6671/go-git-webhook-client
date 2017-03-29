package main

import (
	"net/http"
	"github.com/lifei6671/go-git-webhook-client/conf"
	"log"
	"github.com/lifei6671/go-git-webhook-client/routers"
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


