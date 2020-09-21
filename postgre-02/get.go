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
	"fmt"
	"log"
	"net/http"
)

func getCarList(wr http.ResponseWriter, req *http.Request) {
	if b, s, i := validateMethod(req, "GET"); !b {
		http.Error(wr, s, i)
		return
	}

	brand := req.FormValue("brand")
	banType := req.FormValue("banType")

	rows, err := pdb.Query("select c.id, b.brand_name, m.model_name, c.ban_type, c.miles, c.price, c.currency, c.on_active_loan, c.can_be_bartered, c.description from main.car c join main.brand b on c.brand = b.id join main.model m on m.id = c.model where b.brand_name = $1 and ban_type = $2", brand, banType)
	if err != nil {
		log.Printf("error while fetching data from db: %v", err)
		http.Error(wr, http.StatusText(500), http.StatusInternalServerError)
	}

	defer rows.Close()

	cars := make([]Car, 0)
	found := false
	for rows.Next() {
		found = true
		car := Car{}
		loanFlag := 0
		barterFlag := 0
		err = rows.Scan(&car.id, &car.brand, &car.model, &car.banType, &car.miles, &car.price, &car.currency, &loanFlag, &barterFlag, &car.description)
		if err != nil {
			log.Fatalf("cannot read data from cursor: %v | car: %v", err, car)
		}
		car.canBeBartered = intToBool(barterFlag)
		car.onActiveLoan = intToBool(loanFlag)

		cars = append(cars, car)
	}

	//	err = sql.ErrNoRows

	if !found {
		http.NotFound(wr, req)
		return
	}

	if err = rows.Err(); err != nil {
		log.Fatalf("something went wrong in cursor: %v", err)
	}

	for _, c := range cars {
		fmt.Fprintf(wr, "%d, %s, %s, %s, $%.2f\n", c.id, c.brand, c.model, c.banType, c.miles)
	}
}
