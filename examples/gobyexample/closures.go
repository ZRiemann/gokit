package main

import (
	"fmt"
)

func intSeq() func() int{
	i := 0
	return func() int{
		i += 1
		return i
	}
}

func Singleton(i interface{}) func() interface{}{
	obj := i
	return func() interface{}{
		return obj;
	}
}

type Aaa struct{
	i int
}
func (Aaa) aaa(){
	
}

func main(){
	fmt.Println("closure example")
	nextInt := intSeq()
	fmt.Println(nextInt(), nextInt(), nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	aaa := Aaa{i: 2}
	singleton := Singleton(&aaa)
	j := singleton()
	fmt.Println(Aaa(j).i)
}