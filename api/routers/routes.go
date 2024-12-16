package routers

import (
	"server/controllers"
	"server/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/students/{id}", controllers.GetStudentsByID).Methods("GET")

	adminRouter := router.PathPrefix("/admin").Subrouter()

	adminRouter.Use(middleware.AuthMiddleware)

	adminRouter.HandleFunc("/course", controllers.CreateCourse).Methods("POST")
	adminRouter.HandleFunc("/students", controllers.GetAllStudents).Methods("GET")
	adminRouter.HandleFunc("/students", controllers.CreateStudent).Methods("POST")
	adminRouter.HandleFunc("/students/{id}", controllers.UpdateStudent).Methods("PUT")
	adminRouter.HandleFunc("/students/{id}", controllers.DeleteStudent).Methods("DELETE")
	return router
}
