package main

//#START:interface
type Duck interface {
  Quack()
  Waddle()
  Swim()
}

func (i int) Quack() {}
func (i int) Waddle() {}
func (i int) Swim() {}

func (d Duck) WinAtLife() {} // We can pass this an int
//#END:interface
