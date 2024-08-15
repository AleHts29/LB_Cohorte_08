package routes

import (
	"auth-mux-c09/controllers"
	"github.com/gorilla/mux"
)

func Setup(r *mux.Router) {
	r.HandleFunc("/api/register", controllers.Register).Methods("POST")
	r.HandleFunc("/api/login", controllers.Login).Methods("POST")
	r.HandleFunc("/api/user", controllers.User).Methods("GET")
	r.HandleFunc("/api/logout", controllers.Logout).Methods("POST")
}
