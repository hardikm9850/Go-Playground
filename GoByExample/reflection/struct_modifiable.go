package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    p := Person{"Alice", 30}
    fmt.Println("Before:", p)

    pVal := reflect.ValueOf(&p).Elem()
    nameField := pVal.FieldByName("Name")
    ageField := pVal.FieldByName("Age")

    if nameField.CanSet() {
        nameField.SetString("Bob")
    }
    if ageField.CanSet() {
        ageField.SetInt(40)
    }

    fmt.Println("After:", p)
}

// ------- o/p ------
Before: {Alice 30}
After: {Bob 40}
