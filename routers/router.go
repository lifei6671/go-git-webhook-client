package routers

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Parameter map[string]interface{}

func RegisterRoutes()  {
	r := mux.NewRouter()

	r.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {

		return r.ProtoMajor == 0
	})


}
