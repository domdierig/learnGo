package refVsValue

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	person := person{
		name: name,
		age:  age,
	}
	return &person
}

func RefVsValue() {
	eddy := newPerson("Eddy", 9)

	fmt.Println("Eddy Variable:", eddy.name, eddy.age)

	percy := eddy
	percy.name = "Percy"

	fmt.Println("Eddy Variable:", eddy.name, eddy.age)
	fmt.Println("Percy Variable:", percy.name, percy.age)

	fmt.Println("==================================")

	eddy.name = "Eddy"
	eddyCopy := *eddy
	eddyCopy.name = "EddyCopy"

	fmt.Println("Eddy Variable:", eddy.name, eddy.age)
	fmt.Println("EddyCopy Variable:", eddyCopy.name, eddyCopy.age)

}
