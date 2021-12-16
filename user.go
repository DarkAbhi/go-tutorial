package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"errors"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DB_URL = "host=localhost user=admin password=password dbname=hub port=5432 sslmode=disable TimeZone=Asia/Kolkata"

type User struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string
	Email     string `gorm:"not null"`
}

func InitialMigration() {
	DB, err = gorm.Open(postgres.Open(DB_URL), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&User{})
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	u := DB.First(&user, params["id"])
	if errors.Is(u.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found.")
		return
	}
	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	r.ParseForm()
	w.Header().Set("Content-Type", "application/json")
	var user User
	user.FirstName = r.FormValue("firstname")
	user.LastName = r.FormValue("lastname")
	user.Email = r.FormValue("email")
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}
