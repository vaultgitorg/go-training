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

import "net/http"

func createCar(wr http.ResponseWriter, req *http.Request) {
	if b, s, i := validateMethod(req, "POST"); !b {
		http.Error(wr, s, i)
		return
	}

	// getting parameters from request
	c, err := getParams(req)
	if err != nil {
		http.Error(wr, http.StatusText(400) + " | Error message:" + err.Error(), http.StatusBadRequest)
		return
	}

	// inserting values into database
	_, err = pdb.Exec("insert into main.car(brand, model, ban_type, price, currency) values($1, $2, $3, $4, $5)", c.brandId, c.modelId, c.banType, c.price, c.currency)
	if err != nil {
		http.Error(wr, http.StatusText(500)+" | Error message "+err.Error(), http.StatusInternalServerError)
		return
	}
}

