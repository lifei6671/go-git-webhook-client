package main

import (
	"fmt"
	"net/http"
)

func main()  {

	http.ListenAndServe(":8888", nil)

	fmt.Println("Client ")
}


