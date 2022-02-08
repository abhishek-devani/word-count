package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "sample code"

	check := strings.Contains(str, "sample")
	fmt.Println(check)
}
