package main

import (
	"log"
	"net/http"

	"github.com/JPanettieri/go-seedCrudMySQL/pkg/routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterSeedRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3306", r))
}
