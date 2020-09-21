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
	"errors"
	"net/http"
	"strconv"
)

func validateMethod(req *http.Request, method string) (bool, string, int) {
	if req.Method != method {
		return false, http.StatusText(405), http.StatusMethodNotAllowed
	}
	return true, "", 0
}

// getParams extracting parameter from http.Request interface
func getParams(req *http.Request) (Car, error) {
	c := Car{}
	brand, err := strconv.Atoi(req.FormValue("brand"))
	if err != nil {
		return Car{}, err
	}
	c.brandId = brand

	model, err := strconv.Atoi(req.FormValue("model"))
	if err != nil {
		return Car{}, err
	}

	c.modelId = model
	c.banType = req.FormValue("banType")
	price := req.FormValue("price")
	c.currency = req.FormValue("currency")

	if c.banType == "" || c.currency == "" || price == "" {
		return Car{}, errors.New("input parameters are not valid")
	}

	f, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return Car{}, err
	}
	c.price = f

	return c, nil
}

func intToBool(i int) bool {
	return i == 1
}
