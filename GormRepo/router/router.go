package router

import (
	"gorm/service"
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Method  string
	Pattern string
	Handler http.Handler
}

type Routes []Route

func CreateRoute() *mux.Router {

	router := mux.NewRouter()
	for _, route := range routes {

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.Handler)

	}

	return router
}

var routes = Routes{

	Route{
		"POST",
		"/login",
		http.HandlerFunc(service.Login),
	},

	Route{
		"PUT",
		"/update/user",
		//update the user which is currently logged in
		service.IsAuthorized(http.HandlerFunc(service.UpdateUser)),
	},

	Route{
		"POST",
		"/add/user",
		http.HandlerFunc(service.AddUser),
	},

	Route{
		"GET",
		"/get/{id}/user",
		(http.HandlerFunc(service.GetUser)),
	},

	Route{
		"GET",
		"/get/users",
		(http.HandlerFunc(service.GetAllUser)),
	},
	Route{
		"DELETE",
		"/delete/user",
		//delete the user which is currently logged in
		service.IsAuthorized(http.HandlerFunc(service.DeleteUser)),
	},
	Route{
		"DELETE",
		"/delete/hobby",
		//delete the hobby of the user which is currently logged in
		service.IsAuthorized(http.HandlerFunc(service.DeleteHobbyOfAUser)),
	},
	Route{
		"DELETE",
		"/delete/user/course/{courseid}",
		//delete an associated course of the user which is currently logged in
		service.IsAuthorized(http.HandlerFunc(service.DeleteCourseOfAUser)),
	},
	Route{
		"PUT",
		"/update/{id}/course",
		service.IsAuthorized(http.HandlerFunc(service.UpdateCourse)),
	},
	Route{
		"POST",
		"/add/course",
		http.HandlerFunc(service.AddCourse),
		//service.IsAuthorized(http.HandlerFunc(service.AddCourse)),
	},
	Route{
		"GET",
		"/get/{id}/course",
		(http.HandlerFunc(service.GetCourse)),
	},
	Route{
		"GET",
		"/get/courses",
		(http.HandlerFunc(service.GetAllCourses)),
	},
	Route{
		"DELETE",
		"/delete/{id}/course",
		service.IsAuthorized(http.HandlerFunc(service.DeleteCourse)),
	},
}
