package main

import (
	"log"
	"net/http"
	"project/db"
	"project/handlers"
)

func main() {
	database := db.InitDB()

	http.HandleFunc("/api/commonTable", handlers.GetTable(database))
	http.HandleFunc("/api/addRecord", handlers.AddRecord(database))
	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}