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

//psql -U grealyve -d stockapi
//CREATE DATABASE stockapi;
//CREATE TABLE stocks(stockid SERIAL PRIMARY KEY, name TEXT, price INT, company TEXT);