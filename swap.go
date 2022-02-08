package main

import "fmt"

func main() {
	list := []int{1, 2, 3}
	result := swap(list)
	fmt.Println(result)
}

func swap(l []int) []int {
	var new []int
	for i := len(l) - 1; i >= 0; i-- {
		new = append(new, l[i])
	}
	return new
}
