package main

import (
	"fmt"
	"math"
)

const s string = "constant"
func main(){
	fmt.Println(s)

	const n = 50000000
	const d = 3e20/n

	fmt.Println(d, int64(d), math.Sin(n))
}