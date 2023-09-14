package main

type HasName interface {
	GetName() string
}

func getName(hasName HasName) {
	println("my name is", hasName.GetName())
}

//implement interface

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	var nama Person
	nama.Name = "saleh"
	getName(nama)
}
