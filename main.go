package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	dsn := "mysql:fg0SrkOT6f707V3gzF4e7922O+nXZf1IfkvOVF9pqKU=@tcp(mysql-odsv:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Println("success connect to db")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Halo Dunia"))
	})

	err = http.ListenAndServe(":3000", mux)
	if err != nil {
		panic(err)
	}

	log.Println("server started")
}
