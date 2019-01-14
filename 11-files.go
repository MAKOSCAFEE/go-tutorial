package main

import ("io/ioutil"
		"fmt"
		"os"
)

/* Basic operations.
 * Open- get handle for access
 * Read - read bytes into []bytes
 * Write - Write []byte into file
 * Close - Release handle
 * Seek - move read/write head
 */


func main() {

	// data is []byte filled with contents of the entire file. using Readfile explicit open/close are not needed
	// not feasible for large file

	data, _ := ioutil.ReadFile("LICENCE")

	dat := []byte("Hello, World")

	_ =ioutil.WriteFile("out.txt", dat, 0777)

	_ =ioutil.WriteFile("outed.txt", data, 0777)

	// using os to open and reading 

	f, err := os.Open("out.txt")
	barr := make([]byte, 10)
	if err != nil {
		nb, _ := f.Read(barr)
		fmt.Println(nb)
	}

	f.Close()
}