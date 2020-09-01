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
)

func main() {
	//
	//var arrInt [10]int
	//i := 8
	////println(i)
	//
	//for {
	//	//for i := 0; i < ; i++ {
	//	//}
	//	//if len(arrInt) < cap(arrInt) {
	//	//	arrInt[i] = i
	//	//} else {
	//		println("array length is equal to capacity")
	//
	//		time.Sleep(5)
	//	//}
	//}

	//fmt.Println(arrInt)

	/*	var arrayOfStr [5]string
		fmt.Printf("\n\n\nArray => Length: %d | Capacity: %d \n", len(arrayOfStr), cap(arrayOfStr))

		arrayOfStr[0] = "a"
		arrayOfStr[1] = "b"
		arrayOfStr[2] = "c"
		arrayOfStr[3] = "d"
		arrayOfStr[4] = "e"

		//nil vs empty
		//len() vs cap()

		var slicesOfStr []string = arrayOfStr[1:3]
		fmt.Println(slicesOfStr)

		var slice2 []string //nil
		fmt.Printf("Slice1 => Length: %d | Capacity: %d \n", len(slice2), cap(slice2))
		slice2 = append(slice2, "a")
		fmt.Printf("Slice2 => Length: %d | Capacity: %d \n", len(slice2), cap(slice2))
		fmt.Println(slice2)
		slice2 = append(slice2, "b")
		fmt.Printf("Slice3 => Length: %d | Capacity: %d \n", len(slice2), cap(slice2))
		fmt.Println(slice2)
		slice2 = append(slice2, "c")
		fmt.Println(slice2)
		println("=========================================================================================")
		fmt.Printf("Slice4 => Length: %d | Capacity: %d \n", len(slice2), cap(slice2))
		slice2 = append(slice2, "d", "e")
		fmt.Printf("Slice5 => Length: %d | Capacity: %d \n", len(slice2), cap(slice2))
		fmt.Println(slice2)
		slice2 = append(slice2, "f")
		fmt.Println(slice2)
	*/
	var cars []string
	println(cap(cars))
	println(len(cars))

	lastCapacity := cap(cars)
	for i := 1; i <= 1100; i++ {
		v := fmt.Sprintf("ID: %d", i)
		cars = append(cars, v)

		// checking if the capacity change
		currentCapacity := cap(cars)
		if lastCapacity != currentCapacity {
			capacityChange := (float64(currentCapacity-lastCapacity) / float64(lastCapacity)) * 100
			fmt.Printf("Address: %p \t\t Index: %d \t\t\t Capacity: %d -> %d% \n", &cars, i, currentCapacity, int(capacityChange))
			lastCapacity = currentCapacity
		}

	}
}
