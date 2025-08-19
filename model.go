package main

import "fmt"

type Human struct {
	Name string
	Age  int
}
type Action struct {
	Human
	Role string
}

func (h *Human) SayHello() {
	fmt.Println("Привет! Меня зовут", h.Name)
}

func (a *Action) Do() {
	fmt.Println("я", a.Name, ",выполняю работу как", a.Role)
}
