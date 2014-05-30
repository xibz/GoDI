package UnitTest

import (
  "testing"
  "GoDI/Injector"
  //"reflect"
)

type Animal interface {
  Speak() string
}

type Dog struct {
  Sound string
}

func (d Dog) Speak() string {
  return d.Sound
}

type Cat struct {
  Sound string
}

func (c Cat) Speak() string {
  return c.Sound
}

func Test_Injector(t *testing.T) {
  cnt := GoDI.New()
  d := Dog{"WOOF"}
  c := Cat{"MEOW"}
  cnt.Register(d)
  cnt.Register(c)
  s := cnt.BindToInterface(new(Animal), "Dog")
  r := cnt.BindToInterface(new(Animal), "Cat")
  if s.(Animal).Speak() != "WOOF" && r.(Animal).Speak() != "MEOW" {
    t.Errorf("Injector did not produce right result")
  }
}
