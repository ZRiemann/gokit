package main

import (
	"fmt"
)

type rect struct{
	width, height float64
}

func (r *rect) area() float64{
	return r.width * r.height
}

func (r *rect) perim() float64{
	return 2 *r.width + 2 * r.height
}

type geometry interface{
	area() float64
	perim() float64
}

func measure(g geometry){
	fmt.Println(g, g.area(), g.perim())
}


func main(){
	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area(), "perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area(), "perim:", rp.perim())

	measure(&r)
}