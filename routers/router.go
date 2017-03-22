package routers

import (
	"github.com/gorilla/mux"
	"net/http"
	"go-git-webhook-client/controllers"
	"reflect"
	"fmt"
)

var(
	route = mux.NewRouter()
)

type Parameter map[string]interface{}

func RegisterRoutes()  {

	http.Handle("/",route)
}


func MapRoute(name string,path string,c controllers.ControllerInterface,params Parameter,methods... string){

	route.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		defer func() {
			//if r := recover();r != nil{
			//	fmt.Errorf("%v",r)
			//}
		}()

		controller := reflect.ValueOf(c)

		method := controller.MethodByName("Init")

		if method.IsValid() {
			requestValue := reflect.ValueOf(r)
			responseValue := reflect.ValueOf(w)
			method.Call([]reflect.Value{responseValue, requestValue})

			actionMethod := controller.MethodByName("IndexAction")

			if actionMethod.IsValid() {
				result := actionMethod.Call([]reflect.Value{})

				if len(result) >  0 {
					resultMethod := result[0].MethodByName("ExecuteResult")

					if resultMethod.IsValid() {
						resultMethod.Call([]reflect.Value{responseValue, requestValue})
					}
				}

			}

		}else{
			fmt.Fprint(w,"Method not found")
		}


	}).Methods(methods...).Name(name)
}