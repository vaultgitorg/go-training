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
	"log"
	"net/http"
	"strconv"
)

func updateCar(wr http.ResponseWriter, req *http.Request) {
	if b, s, i := validateMethod(req, "PUT"); !b {
		http.Error(wr, s, i)
		return
	}

	// getting parameters from request
	carId, err := strconv.Atoi(req.FormValue("carId"))
	row := pdb.QueryRow("select $1 as car_id from main.car where id = $2", carId, carId)
	err = row.Scan(&carId)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("raised ErrNoRows error: %v", err)
		http.NotFound(wr, req)
		return
	case err != nil:
		http.Error(wr, http.StatusText(500)+" | Error message:"+err.Error(), http.StatusInternalServerError)
		return
	}

	c, err := getParams(req)
	if err != nil {
		http.Error(wr, http.StatusText(400)+" | Error message:"+err.Error(), http.StatusBadRequest)
		return
	}

	r, err := pdb.Exec("update main.car set brand=$1, model=$2, ban_type=$3, price=$4, currency=$5 where id=$6", c.brandId, c.modelId, c.banType, c.price, c.currency, carId)
	if err != nil {
		http.Error(wr, http.StatusText(500)+" | Error message:"+err.Error(), http.StatusInternalServerError)
		return
	}
	insert, errIns := r.LastInsertId()
	rowsAffected, errAff := r.RowsAffected()
	log.Printf("Last record id: %d, %v, Rows affected: %d, %v \n", insert, errIns, rowsAffected, errAff)

}
