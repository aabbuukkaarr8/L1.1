package main

func main() {

	a := Action{
		Human: Human{Name: "Макс", Age: 20},
		Role:  "Программист",
	}
	a.SayHello()
	a.Do()

}
