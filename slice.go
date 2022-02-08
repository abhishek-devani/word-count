package main

import (
	"bytes"
	"fmt"
)

func main() {

	s1 := []byte{1, 2, 3}
	s2 := []byte{1, 2, 3}
	var s3 []byte
	copy(s3, s1)

	result := bytes.Compare(s1, s2)
	fmt.Println(result)

	if result == 1 {
		fmt.Println("Equal Slices")
	} else {
		fmt.Println("Unequal Slices")
	}

	// fmt.Println("x")

	// go func() {
	// 	fmt.Println("Y")
	// }()

	// time.Sleep(1 * time.Second)

	// go func() {
	// 	fmt.Println("Z")
	// }()

	// time.Sleep(1 * time.Second)

	// fmt.Println("W")
}
