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
  ClimbWalls()
}

type Man interface {
  TakePhotos()
  MissUncleBen()
}

type SpiderMan interface {
  Spider
  Man
}
//#END:interface_embedding

//#START:struct_embedding
struct Spider interface {
  CreateWebs()
  ClimbWalls()
}

struct Man interface {
  TakePhotos()
  MissUncleBen()
}

struct SpiderMan interface {
  Spider
  Man
}

func (m Man) GetPaid() {} // Can pass a Man or a SpiderMan
//#END:struct_embedding
