package main

import (
	"auth-mux-c09/database"
	"auth-mux-c09/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	database.Connect()

	r := mux.NewRouter()

	//	Cords
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
	})

	//	Registrar rutas
	routes.Setup(r)

	handler := c.Handler(r)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
