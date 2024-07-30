package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// definimos estructura para el mapeo de la info
type Gasto struct {
	ID       string  `json:"id"`
	Concepto string  `json:"concepto"`
	Importe  float64 `json:"importe"`
	Fecha    string  `json:"fecha"`
}

type Reespuesta struct {
	Msg string `json:"msg"`
}

// Simulamos una DB
var gastos []Gasto

func getGastos(w http.ResponseWriter, r *http.Request) {
	//Definimos el header
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gastos)
}

func getGasto(w http.ResponseWriter, r *http.Request) {
	//Capturamos el param
	params := mux.Vars(r)

	//Buscamos el elemento en el array
	for _, item := range gastos {
		if item.ID == params["id"] {
			//Definimos el header
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	//	si no encontro el elemento
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&Gasto{})
}

func createGasto(w http.ResponseWriter, r *http.Request) {
	//r = {
	//	"id": "5",
	//		"concepto": "pagp de luz",
	//		"importe": 50,
	//		"fecha": "2024-10-12"
	//}
	var gasto Gasto
	_ = json.NewDecoder(r.Body).Decode(&gasto)

	//No se puede porque necesita hacer el decode
	//newGasgo := json.NewDecoder(r.Body)
	//gasto = &newGasgo

	// agregamos la info al array
	gastos = append(gastos, gasto)

	//Definimos el header
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gasto)
}

func updateGasto(w http.ResponseWriter, r *http.Request) {
	//Capturamos el param
	params := mux.Vars(r)

	for index, item := range gastos {
		if item.ID == params["id"] {

			//Eliminamos el elemnto
			gastos = append(gastos[:index], gastos[index+1:]...)

			var gasto Gasto
			_ = json.NewDecoder(r.Body).Decode(&gasto)

			gasto.ID = params["id"]
			gastos = append(gastos, gasto)

			//Definimos el header
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(gasto)
			return
		}
	}

	//	si no encontro el elemento
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&Gasto{})
}

func deleteGasto(w http.ResponseWriter, r *http.Request) {
	//Capturamos el param
	params := mux.Vars(r)

	for index, item := range gastos {
		if item.ID == params["id"] {
			//Eliminamos el elemnto
			gastos = append(gastos[:index], gastos[index+1:]...)

			//Definimos el header
			w.Header().Set("Content-Type", "application/json")
			//json.NewEncoder(w).Encode(gastos)
			json.NewEncoder(w).Encode(Reespuesta{Msg: "El elemento se borro de forma exitosa"})
			return
		}
	}

	//	si no encontro el elemento
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(Reespuesta{Msg: "El elemento No se pudo borrar de forma exitosa"})
}

func main() {
	// Inicializar algunos datos
	gastos = append(gastos, Gasto{ID: "1", Concepto: "compra de artículos", Importe: 50.00, Fecha: "2010-10-10"})
	gastos = append(gastos, Gasto{ID: "2", Concepto: "compra semanal de comida", Importe: 165.00, Fecha: "2010-10-20"})
	gastos = append(gastos, Gasto{ID: "3", Concepto: "reparación auto", Importe: 2400.00, Fecha: "2010-11-14"})
	gastos = append(gastos, Gasto{ID: "4", Concepto: "servicios de cable", Importe: 240.00, Fecha: "2010-11-10"})

	//	Creamos el enrutador usando Gorilla Mux
	r := mux.NewRouter()

	//	Rutas o endpoints
	r.HandleFunc("/gastos", getGastos).Methods("GET")
	r.HandleFunc("/gastos/{id}", getGasto).Methods("GET")
	r.HandleFunc("/gastos", createGasto).Methods("POST")
	r.HandleFunc("/gastos/{id}", updateGasto).Methods("PUT")
	r.HandleFunc("/gastos/{id}", deleteGasto).Methods("DELETE")

	//	Iniciamos el server
	log.Println("Server run on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
