package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Shashika2001/bookstore-api/routes"
)

func main() {
	router := routes.SetupRouter()
	fmt.Println("Server running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
