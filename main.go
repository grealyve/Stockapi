package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/grealyve/golang-stockapi/router"
	
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on 8000...")

	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
