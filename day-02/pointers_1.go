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

import "fmt"

type Human struct {
	isWorking bool
	name      string
	gender    string
	age       int
}

func main2() {
	//var h1 Human
	//h1.age = 22
	//h1.name = "Elxan"
	//h1.gender = "Male"
	//fmt.Println(h1)
	//
	//h2 := Human{name: "Vusal", isWorking: true}
	//fmt.Println(h2)
	//
	//
	//fmt.Println(Human{})

	humanValue := Human{name: "Elxan", isWorking: true, age: 55}
	humanAddress := &humanValue
	fmt.Printf("\n\n\n\nhumanValue=%v | humanAddress=%p\n", humanValue, humanAddress)
	buildHuman(humanAddress)
	fmt.Println(humanValue)
	fmt.Printf("humanValue=%v | humanAddress=%p\n\n\n\n", humanValue, humanAddress)
}

func buildHuman(v *Human) {
	fmt.Printf("v=%p | &v=%p\n", v, &v)
	changeHumanData(v)
}

func changeHumanData(v *Human) {
	fmt.Printf("v=%p | &v=%p\n", v, &v)
	v.isWorking = false
	v.name = v.name + " NewName"
}
