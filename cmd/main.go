package main

import (
	"log"
	"net/http"

	"github.com/JPanettieri/go-seedCrudMySQL/tree/master/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterSeedRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3306", r))
}
