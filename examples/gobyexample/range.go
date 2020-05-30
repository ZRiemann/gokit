package main

import (
	"fmt"
)

func main(){
	nums := []int{2, 3, 4}
	sum := 0
	for idx, num := range nums{
		sum += num
		fmt.Printf("[%d:%d] + ", idx, num)
	}
	fmt.Println("sum: ", sum)

}