package main

import (
	"centralserver/controllers"
	"centralserver/data"
	// "centralserver/middleware"
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("files.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&data.Handlelingauthetication{})
	server := http.NewServeMux()

	server.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		controllers.Handleautth(db, w, r)
	})
	// server.HandleFunc("/download", middleware.JWTMiddleware(controllers.Handledownload))

	erro := http.ListenAndServe(":8081", server)
	if erro != nil {
		log.Fatal(erro)
	}
	defer fmt.Println("Program exited")
}
