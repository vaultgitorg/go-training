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

func deleteCar(wr http.ResponseWriter, req *http.Request) {
	if b, s, i := validateMethod(req, "DELETE"); !b {
		http.Error(wr, s, i)
		return
	}

	carId, err := strconv.Atoi(req.FormValue("carId"))
	if err != nil {
		log.Printf("car id is not in valid format: %v", err)
		http.Error(wr, http.StatusText(400), http.StatusBadRequest)
		return
	}
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

	r, err := pdb.Exec("delete from main.car where id = $1", carId)
	insert, errIns := r.LastInsertId()
	rowsAffected, errAff := r.RowsAffected()
	log.Printf("Last record id: %d, %v, Rows affected: %d, %v \n", insert, errIns, rowsAffected, errAff)

	if err != nil {
		log.Printf("couldn't delete record(%d) from database: %v", carId, err)
		http.Error(wr, http.StatusText(500), http.StatusInternalServerError)
		return
	}

}

