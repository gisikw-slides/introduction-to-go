package main

import "fmt"

func main() {

  my_map := map[string] string{}

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
