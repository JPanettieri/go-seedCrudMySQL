package models

import (
	"github.com/JPanettieri/go-seedCrudSQL/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Seed struct {
	gorm.Model
	Name     string `gorm: "json:""name"`
	Season   string `gorm: "json:""season"`
	Rainfall string `gorm: "json:""rainfall"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Seed{})
}

func (s *Seed) CreateSeed() *Seed {
	db.NewRecord(s)
	db.Create(&s)
	return s
}

func GetAllSeeds() []Seed {
	var Seeds []Seed
	db.Find((&Seeds))
	return Seeds
}

func GetSeedBySeason(Season string) (*Seed, *gorm.DB) {
	var getSeed Seed
	db := db.Where("SEASON=?", Season).Find(&getSeed)
	return &getSeed, db
}

func GetSeedByName(Name string) (*Seed, *gorm.DB) {
	var getSeed Seed
	db := db.Where("NAME=?", Name).Find(&getSeed)
	return &getSeed, db
}

func DeleteSeed(Name string) Seed {
	var seed Seed
	db.Where("NAME=?", Name).Delete(seed)
	return seed
}
