package GoDI

import (
  "reflect"
)

type Container struct {
  injMap map[string] interface{}
  factoryMethod map[string] (func() interface{})
}

func New() *Container {
  c := Container{}
  c.injMap = make(map[string] interface{})
  c.factoryMethod = make(map[string] (func() interface{}))
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

func (c *Container) BindFactoryMethod(name string, f func() interface{}) {
  c.factoryMethod[name] = f
}

func (c *Container) CallFactoryMethod(name string) interface{} {
  return c.factoryMethod[name]()
}
