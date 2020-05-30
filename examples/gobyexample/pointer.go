package main

import (
	"fmt"
)

func zeroval(val int) int{
	val = 0
	return val
}

func zeroptr(iptr *int) *int{
	*iptr = 0
	return iptr
}
func main(){
	fmt.Println("pointer example")
	i := 1
	zeroval(i)
	fmt.Println(i)
	zeroptr(&i)
	fmt.Println(i, &i)
}