package main

import (
	"fmt"
)

func main(){
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map: ", m, "k1->", m["k1"])

	delete(m, "k2")
	fmt.Println("delete k2 map: ", m)
	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)
}