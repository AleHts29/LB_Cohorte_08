package main

import (
	"c10_practica/internal/config"
	"c10_practica/internal/database"
	"c10_practica/internal/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	//	Lectura del archivo config.yaml - configuraicones
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	//	Conectarme a la DB
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	//	Me fijo si la tabla users existe y si no la creo
	if err := database.AutoMigrate(db); err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	//	instancio el router de gorilla mux
	router := mux.NewRouter()
	handlers.RouterHandlers(router, db)

	//	Levantamos el servidor
	log.Printf("Server starting on %s", cfg.ServerAddress)
	if err := http.ListenAndServe(cfg.ServerAddress, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
