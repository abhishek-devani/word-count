package main

import "fmt"

func main() {
	fmt.Println(fibo(7))
}

func fibo(n int) int {
	a, b := 0, 1
	c := 0
	if n < 2 {
		return n
	} else {
		for i := 2; i < n+1; i++ {
			c = a + b
			a = b
			b = c
		}
	}
	return c
}
