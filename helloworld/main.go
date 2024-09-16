package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	person := Person{
		Name: "Nico",
	}
	fmt.Println(person.PrtinName())

}

type Person struct {
	Name string
}

func (p *Person) PrtinName() string {
	return p.Name
}
