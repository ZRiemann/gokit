package main

import (
	"fmt"
)

func fact(n int) int{
	fmt.Printf("%d->", n)
	if n == 0{
		return 1
	}
	return n * fact(n - 1)
}
func main(){
	fmt.Println("recursion example")
	fmt.Println("\n", fact(7))
}