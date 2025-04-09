// problem- 1 :  Swap Two Numbers Using Pointer
// problem- 1 :  Update Struct Using Pointer

package main

import "fmt"

// func swap(a *int, b *int) {
//  c:= *a
//  *a = *b
//  *b = c
// }

// // problem - 1
// func main() {
//   x := 5
//   y := 10

//   swap(&x, &y)

//   fmt.Println("x =", x) // Output: x = 10
//   fmt.Println("y =", y) // Output: y = 5
// }

//problem - 2
// type Person struct {
//   name string
//   age  int
// }

// func birthday(p *Person) {
//   // age বাড়াও
//   p.age +=1
// }
// func main() {
//   p1 := Person{name: "Sabbir", age: 24}
//   birthday(&p1)
//   fmt.Println(p1.age) // Output: 25
// }

// problem - 3

func printValue(p *int) {
  // যদি p nil না হয় তাহলে মান প্রিন্ট করো, না হলে বলো "Nil pointer"
 if(p == nil){
  fmt.Println("Nil pointer")
 }else{
  fmt.Println(*p)
 }
}

func main() {
  var a *int
  var b int = 42

  printValue(a)    // Output: Nil pointer
  printValue(&b)   // Output: 42
}