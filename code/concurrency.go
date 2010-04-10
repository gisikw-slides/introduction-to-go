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

//#START:semaphore
func client(input chan int) {
  for true {
    <- input
    // Do something
    input <- 1
  }
}
//#END:semaphore

//#START:asynchronous_channels
func chatter(){
  output := make(chan string 10)
  go log(output)
  input <- "Hey. Hey. Listen"
  // Chatter terminates almost immediately
}
//#END:asynchronous_channels

//#START:thread_join
func spawner(){
  gui_dead := make(chan int)
  go gui()

  // Do some stuff
  <- gui_dead
}
//#END:thread_join
