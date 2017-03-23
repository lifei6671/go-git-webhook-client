package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"go-git-webhook-client/controllers"
)

var(
	route = mux.NewRouter()
)


func RegisterRoutes()  {

	route.HandleFunc("/socket",controllers.WebSocketServer).Name("web_socket_server")
	route.HandleFunc("/payload/{key}",controllers.Payload).Name("payload")
	route.HandleFunc("/token",controllers.Token).Name("token").Methods("POST")

	http.Handle("/",route)
}
