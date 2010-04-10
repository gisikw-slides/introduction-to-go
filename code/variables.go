package main

import "fmt"

func main() {

  //#START:declarations
  // Declare a variable
  var s string = "";

  // Go infers the type
  var s2 = "";

  // Syntactic shorthand - initializing declaration
  s3 := "";
  //#END:declarations
  
  //#START:arrays
  var arrayOfInt [10]int
  //#END:arrays

  //#START:maps
  var m map = map[string] int{}
  m["price"] = 5
  //#END:maps

  fmt.Println(s)
  fmt.Println(s2)
  fmt.Println(s3)
  fmt.Println(arrayOfInt)

}
