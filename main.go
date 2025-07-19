package main

import (
	"beck1/pages/home"
	"log"

	"github.com/rohanthewiz/rweb"
)

func main() {
	/*
		TODO
		- move css to own file
		- templatize menu
		- pick out some decent Google fonts
		- create hierarchy of pages and components
	*/

	// Initialize the server with options
	s := rweb.NewServer(rweb.ServerOptions{
		Address: ":8000",
		Verbose: true,
	})

	// Add request info middleware for stats
	s.Use(rweb.RequestInfo)
	s.ElementDebugRoutes()

	// Static routes
	s.StaticFiles("/css/", "assets/css", 1)
	s.StaticFiles("/img/", "assets/img", 1)

	// Define routes
	s.Get("/", home.HomeHandler)

	log.Println("Starting server on http://localhost:8000")
	log.Fatal(s.Run())
}
