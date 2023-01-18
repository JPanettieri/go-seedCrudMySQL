package routes

import (
	"github.com/JPanettieri/go-seedCrudMySQL/tree/master/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterSeedRoutes = func(router *mux.Router) {
	router.HandleFunc("/seed/", controllers.CreateSeed).Methods("POST")
	router.HandleFunc("/seed/", controllers.GetSeed).Methods("GET")
	router.HandleFunc("/seed/{season}", controllers.GetSeedBySeason).Methods("GET")
	router.HandleFunc("/seed{season}", controllers.UpdateSeed).Methods("PUT")
	router.HandleFunc("/seed{season}", controllers.DeleteSeed).Methods("DELETE")
}
