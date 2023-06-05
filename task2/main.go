package main

import (
	"fmt"
	"log"
	"net/http"
	controllers "task2/controllers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const INTERFACE = "127.0.0.1"
const PORT = "8000"
const SERVER = INTERFACE + ":" + PORT

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/spots", controllers.GetSpotsController).Methods("GET")

	fmt.Println("Listening on port 8000")
	log.Fatal(http.ListenAndServe(SERVER, router))
}
