package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	//contactInfo contactInfo
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	james := person{firstName: "James", lastName: "Bond"}
	fmt.Println(james) // {James Bond}

	var rooney person // assigns zero value
	rooney.firstName = "Wayne"
	rooney.lastName = "Rooney"
	fmt.Println(rooney)       //{Wayne Rooney}
	fmt.Printf("%+v", rooney) //{firstName:Wayne lastName:Rooney}

	reus := person{
		firstName: "Marco",
		lastName:  "Reus",
		contactInfo: contactInfo{
			email:   "marco@gmail.com",
			zipCode: 12, //Every single line must have a comma, even if it is a last line in decoration
		}, // same here
	}
	fmt.Printf("%+v", reus) //{firstName:Marco lastName:Reus contactInfo:{email:marco@gmail.com zipCode:12}}
	reus.print()
	reusPointer := &reus

	reusPointer.updateName("Jaden")
	println("Changed Reus's name to Jaden!")
	reusPointer.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}

/**
About Pointers
&variable : Give me the memory address of the value this variable is pointing at.
It will fetch us the memory address in RAM where the variable is presently residing

*pointer : Give me the value this memory address is pointing at
It will fetch the actual struct/type/data  in the memory block where the pointer is currently pointing. Phewwww
*/
