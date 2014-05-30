package GoDI

import (
  "reflect"
)

type Container struct {
  injMap map[string] interface{}
}

func New() *Container {
  c := Container{}
  c.injMap = make(map[string] interface{})
  return &c
}
func (c *Container) Register(obj interface{}) {
  v := reflect.ValueOf(obj)
  c.injMap[v.Type().Name()] = obj
}

func (c *Container) BindToInterface(s interface{}, objName string) interface{} {
  s = c.injMap[objName]
  return s
}
