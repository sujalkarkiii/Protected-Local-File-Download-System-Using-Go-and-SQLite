package main

import (
	"centralserver/controllers"
	"fmt"
	"log"
	"net/http"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
func main() {
	DB, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
    panic("failed to connect database")
  }

	server:=http.NewServeMux()
	
	server.HandleFunc("/post",controllers.Handleautth)
	fs:=http.FileServer(http.Dir("./Public"))
	server.Handle("/",fs)

    erro := http.ListenAndServe(":8081", server)
    if erro != nil {
		log.Fatal(erro)
    }
	defer fmt.Println("Program exited")
}
