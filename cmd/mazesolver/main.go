package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Mussabeheen/mazeSolver/pkg/maze"
	"github.com/gorilla/mux"
)

var port = 9000

func main() {
	controller := maze.NewController(maze.Newservice())
	router := mux.NewRouter()
	controller.RegisterRoutes(router)
	fmt.Println("Starting server :", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
