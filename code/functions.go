package main

import "fmt"

func main() {

  my_map := map[string] string{}

  //#START:method
  type Mongoose int

  func (m Mongoose) Pluck() string {
    return "The mongoose was plucked"
  }
  //#END:method

  //#START:call_gimme_five
  x, message := gimmeFive()
  //#END:call_gimme_five

  //#START:why_multiple
  if text, err := readFile('foo.txt'); err == nil {
    // Read the file
  } else {
    // Handle the error
  }

  for key, value := range my_map {

  }
  
  // ...or
  for _, value := range my_map { // Discard the first returned value

  }
  //#END:why_multiple

  fmt.Println(x,message)
}

//#START:multiple_return
func gimmeFive() (int, string) {
  return 5, "thanks for asking!"
}
//#END:multiple_return

func readFile(name string) (out string, err string) {
  out = "File read"
  return
}

//#START:named_results
func gimmeSix() (value int, err string) {
  // value and err are set to nil
  value = 5
  return // implicitly returns 5, nil
}
//#END:named_results
