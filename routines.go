/*
   go routines
   normal way: f()
   routines: go f()
*/

package main

import "fmt"

func f(name string) {
	for i := 0; i < 100; i++ {
		fmt.Println(name, ":", i)
	}
}

func main() {

	//demonstrates the concurrent go routines

	// f("main1")

	go f("go routine  1")
	// f("main2")
	go f("go routine     2")
	// f("main3")

	// go routines are executed concurrently
	// execution falls through here
	fmt.Scanln()
}
