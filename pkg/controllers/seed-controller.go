package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/JPanettieri/go-seedCrudMySQL/tree/master/pkg/models"
	"github.com/JPanettieri/go-seedCrudMySQL/tree/master/pkg/utils"
	"github.com/gorilla/mux"
)

var NewSeed models.Seed

func GetSeed(w http.ResponseWriter, r *http.Request) {
	NewSeed := models.GetAllSeeds()
	res, _ := json.Marshal(NewSeed)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetSeedBySeason(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	season := vars["season"]
	seedDetails, _ := models.GetSeedBySeason(season)
	res, _ := json.Marshal(seedDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateSeed(w http.ResponseWriter, r *http.Request) {
	CreateSeed := &models.Seed{}
	utils.ParseBody(r, CreateSeed)
	s := CreateSeed.CreateSeed()
	res, _ := json.Marshal(s)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteSeed(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	seed := models.DeleteSeed(name)
	res, _ := json.Marshal(seed)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateSeed(w http.ResponseWriter, r *http.Request) {
	updateSeed := &models.Seed{}
	utils.ParseBody(r, updateSeed)
	vars := mux.Vars(r)
	name := vars["name"]

	seedDetails, db := models.GetSeedByName(name)
	if updateSeed.Name != "" {
		seedDetails.Name = updateSeed.Name
	}
	if updateSeed.Season != "" {
		seedDetails.Season = updateSeed.Season
	}
	if updateSeed.Rainfall != "" {
		seedDetails.Rainfall = updateSeed.Rainfall
	}
	db.Save(&seedDetails)
	res, _ := json.Marshal(seedDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
