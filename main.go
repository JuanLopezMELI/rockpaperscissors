package main

import (
	"log"
	"net/http"
	"rockpaperscissors/handlers"
)

func main() {
	//Create router
	router := http.NewServeMux()

	//Serve static files
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	port := ":8080"
	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
