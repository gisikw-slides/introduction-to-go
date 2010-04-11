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

//#START:interface_embedding
type Spider interface {
  CreateWebs()
}

type Man interface {
  MissUncleBen()
}

type SpiderMan interface {
  Spider
  Man
}

func (m SpiderMan) FightGreenGoblin() {} // Target must CreateWebs() and MissUncleBen()
//#END:interface_embedding

//#START:struct_embedding
type Spider struct {
  CreateWebs()
}

type Man struct {
  MissUncleBen()
}

type SpiderMan struct {
  Spider
  Man
}

func (m Man) GetPaid() {} // Can pass a Man or a SpiderMan
//#END:struct_embedding
