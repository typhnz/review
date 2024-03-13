package main

import "fmt"

func main() {
	fibonacci := fibonacciGenerator()
	for i := 0; i < 10; i++ {
		fmt.Print(fibonacci(), " ")
	}
}

func fibonacciGenerator() func() int {
	f1 := 0 /*first two fibonacci numbers we must return ourselves*/
	f2 := 1
	return func() int {
		f2, f1 = (f1 + f2), f2
		return f1
	}
}