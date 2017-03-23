package controllers

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func Payload(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)


	fmt.Fprint(w,params)
}
