package main

import (
	"fmt"
)

func plus(a int, b int) int{
	return a + b
}

func mult_plus(a int, b int) (int, bool){
	return a + b, true
}

func sum(nums ...int){
	fmt.Print(nums, " ")

	total := 0
	for _, num := range nums{
		total += num
	}
	fmt.Println(total)
}

func main(){
	plus(1, 2)
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}