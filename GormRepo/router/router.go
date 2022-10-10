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
		"/user",
		//update the user which is currently logged in
		service.IsAuthorized(http.HandlerFunc(service.UpdateUser)),
	},

	Route{
		"POST",
		"/user",
		http.HandlerFunc(service.AddUser),
	},

	Route{
		"GET",
		"/users/{id}",
		(http.HandlerFunc(service.GetUser)),
	},

	Route{
		"GET",
		"/users",
		(http.HandlerFunc(service.GetAllUser)),
	},
	Route{
		"DELETE",
		"/user",
		//delete the user which is currently logged in
		service.IsAuthorized(http.HandlerFunc(service.DeleteUser)),
	},
	Route{
		"DELETE",
		"/user/hobby",
		//delete the hobby of the user which is currently logged in
		service.IsAuthorized(http.HandlerFunc(service.DeleteHobbyOfAUser)),
	},
	Route{
		"DELETE",
		"/user/courses/{courseid}",
		//delete an associated course of the user which is currently logged in
		service.IsAuthorized(http.HandlerFunc(service.DeleteCourseOfAUser)),
	},
	Route{
		"PUT",
		"/courses/{id}",
		service.IsAuthorized(http.HandlerFunc(service.UpdateCourse)),
	},
	Route{
		"POST",
		"/course",
		http.HandlerFunc(service.AddCourse),
		//service.IsAuthorized(http.HandlerFunc(service.AddCourse)),
	},
	Route{
		"GET",
		"/courses/{id}",
		(http.HandlerFunc(service.GetCourse)),
	},
	Route{
		"GET",
		"/courses",
		(http.HandlerFunc(service.GetAllCourses)),
	},
	Route{
		"DELETE",
		"/courses/{id}",
		service.IsAuthorized(http.HandlerFunc(service.DeleteCourse)),
	},
}
