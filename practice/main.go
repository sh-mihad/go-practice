package main

import "fmt"

type Person struct {
name string
age int
}

//receiver function
func (personObj Person)showPerson(){
 fmt.Println(personObj.name)
  fmt.Println(personObj.age)
}
func (p *Person) birthday() {
    p.age++
}
func main (){
 p1:= Person{name:"Sabbir",age:24}
 p1.showPerson() // calling the method
 p1.birthday()
  p1.showPerson() // calling the method
}