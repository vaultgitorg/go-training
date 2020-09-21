/*
Copyright 2020 Elkhan Ibrahimov

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Car struct {
	id            int
	brand         string
	model         string
	banType       string
	miles         float64
	price         float64
	currency      string
	onActiveLoan  bool
	canBeBartered bool
	description   string
	brandId       int
	modelId       int
}

var pdb *sql.DB

func init() {
	fmt.Println("opening database connection")
	var err error
	pdb, err = sql.Open("postgres", "postgres://user1:password1@localhost/turbo_az?sslmode=disable")
	if err != nil {
		log.Fatalf("cannot connect to the postgresql database: %v", err)
	}
	fmt.Println("Connected to database...")
}

func main() {

	http.HandleFunc("/cars", getCarList)
	http.HandleFunc("/cars/delete", deleteCar)
	http.HandleFunc("/cars/create", createCar)
	http.HandleFunc("/cars/update", updateCar)
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		log.Fatal("Cannot start port to listen")
	}
}
