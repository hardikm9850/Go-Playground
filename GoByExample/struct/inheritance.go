package main

import "fmt"

type Animal struct {
	Name string
}

type Dog struct {
	Animal Animal
	Breed  string
	Sound  string
}

func (d Dog) speak() {
	fmt.Printf("Breed %s, Sound %s\n", d.Breed, d.Sound)
}

func (d *Dog) changeSound(newSound string) {
	d.Sound = newSound
}

func main() {
	lab := Dog{
		Animal: Animal{
			Name: "Dog",
		},
		Breed: "Labrador",
		Sound: "Woof Woof",
	}

	lab.speak()
	println("Changing the sound of dog")
	lab.changeSound("Bark Bark")
	lab.speak()
}
