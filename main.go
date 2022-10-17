package main

import (
	"encoding/json"
	"github.com/debbysa/moleo/core/module"
	"github.com/debbysa/moleo/core/repository/category"
	"github.com/debbysa/moleo/handler/api"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var db *gorm.DB

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/moleo?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	categoryRepo := category.NewCategoryRepository(db)
	categoryUseCase := module.NewCategoryUsecase(categoryRepo)
	categoryHandler := api.NewCategoryHandler(categoryUseCase)

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handleHome)
	router.HandleFunc("/login", handleLogin).Methods("POST")

	router.HandleFunc("/category/{id}", categoryHandler.GetCategoryById).Methods("GET")
	router.HandleFunc("/category", categoryHandler.GetCategoryList).Methods("GET")
	router.HandleFunc("/category", categoryHandler.CreateCategory).Methods("POST")
	router.HandleFunc("/category/{id}", categoryHandler.UpdateCategory).Methods("PUT")
	router.HandleFunc("/category/{id}", categoryHandler.DeleteCategory).Methods("DELETE")

	err = http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatalf("error server serve = %v", err)
	}
}

type user struct {
	ID       int
	Username string
	Password string
}

type userRequest struct {
	Username string
	Password string
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Welcome to Moleo"))
	if err != nil {
		return
	}
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var userReq userRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&userReq)
	if err != nil {
		log.Fatalf("error when decode = %v", err)
	}

	var users []user
	result := db.Table("user").Where(&user{Username: userReq.Username, Password: userReq.Password})

	if result.Error != nil {
		w.Write([]byte("Login failed"))
	}

	if len(users) > 0 {
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": userReq.Username,
		}).SignedString([]byte("secret"))
		if err != nil {
			w.Write([]byte("Login failed"))
		}
		w.Write([]byte(token))
	} else {
		w.Write([]byte("Login failed"))
	}

}
