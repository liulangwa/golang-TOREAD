package main

import (
	"fmt"
	"reflect"
)

func main() {
	var k *float64 = nil
	var t2 interface{} = k
	fmt.Println("a1:", t2 == k)   //true
	fmt.Println("a1:", k == nil)  //true
	fmt.Println("a1:", t2 == nil) //false

	fmt.Println(reflect.TypeOf(t2).Kind())

}
