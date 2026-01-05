package main

import (
	"centralserver/controllers"
	"centralserver/data"
	"centralserver/middleware"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// use below for login 
    //         {RollNo: 101, Department: "Computer"},
    //         {RollNo: 102, Department: "Electronics"},
    //         {RollNo: 103, Department: "Mechanical"},
    //     }



func main() {
	godotenv.Load()

	db, err := gorm.Open(sqlite.Open("files.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&data.Handlelingauthetication{})

	server := http.NewServeMux()

	// Pages
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./Public/index.html")
	})
server.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./Public/download.html")
})

	// Auth
	server.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		controllers.Handleautth(db, w, r)
	})

	// APIs
	server.HandleFunc("/api/files",
		middleware.JWTMiddleware(controllers.HandleFetchingFiles),
	)

	server.HandleFunc("/download",
		middleware.JWTMiddleware(controllers.HandleDownload),
	)

	fmt.Println("Server running on http://0.0.0.0:8081")
	log.Fatal(http.ListenAndServe(":8081", server))
}
