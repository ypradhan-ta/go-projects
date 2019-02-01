/*
   File IO: Code to demonstrate working with files
*/

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	f, err := os.Create("./tmp/dat2")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)

	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
	d1 := []byte("hello\ngo\n")
	err = ioutil.WriteFile("./tmp/dat1", d1, 0644)
	check(err)

	dataObject := "{'name':'yash pradhan', 'username':'user1234', 'phone':'2145112345'}"
	byteStream := []byte(dataObject)
	err = ioutil.WriteFile("./tmp/data.txt", byteStream, 0644)
	check(err)

}
