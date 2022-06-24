package main

import "fmt"

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

type Pet interface {
	Dog | Cat
}

func NewPet[T Pet](name string) T {
	a := new(T)
	i := any(a)

	switch i.(type) {
	case *Dog:
		v := Dog{name}
		return any(v).(T)
	default:
		v := Cat{name}
		return any(v).(T)
	}
}

func main() {
	d := NewPet[Cat]("Sam")

	fmt.Println(d.Name)
}
