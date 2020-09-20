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
)

type Car struct {
	miles         float64
	id            int
	brand         string
	model         string
	banType       string
	price         float64
	currency      string
	onActiveLoan  bool
	canBeBartered bool
	description   string
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
	fmt.Println("main function executed")
	defer pdb.Close()

	s := TempStruct{}
	fmt.Println(s)

	/*
	   select c.id, b.brand_name, m.model_name, c.ban_type, c.miles, c.price, c.currency, c.on_active_loan, c.can_be_bartered, c.description from car c join brand b on c.brand = b.id join model m on m.id = c.model;
	*/

	rows, err := pdb.Query("select c.id, b.brand_name, m.model_name, c.ban_type, c.miles, c.price, c.currency, c.on_active_loan, c.can_be_bartered, c.description from main.car c join main.brand b on c.brand = b.id join main.model m on m.id = c.model;")
	if err != nil {
		log.Fatalf("error while fetching data from db: %v", err)
	}
	defer rows.Close()

	cars := make([]Car, 0)
	for rows.Next() {
		car := Car{}
		loanFlag := 0
		barterFlag := 0
		rows.Columns()
		err := rows.Scan(&car.id, &car.brand, &car.model, &car.banType, &car.miles, &car.price, &car.currency, &loanFlag, &barterFlag, &car.description)
		if err != nil {
			log.Fatalf("cannot read data from cursor: %v | car: %v", err, car)
		}
		car.canBeBartered = intToBool(barterFlag)
		car.onActiveLoan = intToBool(loanFlag)

		cars = append(cars, car)
	}

	if err = rows.Err(); err != nil {
		log.Fatalf("something went wrong in cursor: %v", err)
	}

	for _, c := range cars {
		fmt.Println(c)
	}

}

func intToBool(i int) bool {
	return i == 1
}
