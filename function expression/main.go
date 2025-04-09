package main

import "fmt"

const p = 10;
var a = 100

func outerFunction () func(){
    money := 100
   age :=30
   show := func ()  {
    fmt.Println(age)
    money = money + p + a
   }

   return show
}

func call()  {
    print1 := outerFunction()
    print1()
    print1()

    print2 := outerFunction()
    print2()
    print2()
}

func main()  {
    call()
}

func init()  {
    fmt.Println("----Bank App---")
}