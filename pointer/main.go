package main

import "fmt"

func main (){
	x:= 10

	a := &x

	fmt.Println("memory address", a)
	fmt.Println("value of x", *a)
}