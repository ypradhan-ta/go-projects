/* maps in go */

package main

import "fmt"

func main() {
	myMap := make(map[string]int)

	myMap["key1"] = 10
	myMap["key2"] = 20

	fmt.Println("myMap:", myMap)
	val, present := myMap["key2"]

	fmt.Println("myMap:", myMap)
	if present {
		fmt.Println("myMap[k2]", val)
	}
	delete(myMap, "key2")
	fmt.Println("myMap:", myMap)

	val, present = myMap["k2"]

	if present {
		fmt.Println("myMap[k2]", val)
	}
}
