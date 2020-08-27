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

func main() {
	var humanList1 [20]Human
	humanList1[0] = Human{name: "Elxan", isWorking: false}
	humanList1[1] = Human{name: "John", isWorking: true}
	humanList1[2] = Human{name: "Elthon", gender: "Male"}
	fmt.Println(humanList1, "\n\n")

	humanList2 := [5]int8{1, 2, 10, 12, 127}
	fmt.Println(humanList2, "\n\n")

	humanList3 := []int8{1, 10, 12, 127}
	fmt.Println(humanList3, "\n\n")

	size := 15
	humanList4 := make([]bool, size)
	fmt.Println(humanList4, "\n\n")
}
