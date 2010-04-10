package main

func main() {
  //#START:goroutines
  go gimmeFive()

  go func() {
    time.Sleep(5)
    fmt.Println("Computer over. Virus = very yes")
  }()
  //#END:goroutines

}

func gimmeFive() {

}

//#START:channels
func chat() {
  output := make(chan string)
  go log(output)
  for true {
    input <- "Hey. Hey. Listen"
  }
}

func log(input chan string){
  for true {
    data := <- input
    fmt.Println(data)
  }
}
//#END:channels
