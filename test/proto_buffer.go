package main

import (
	"log"
	"os"
)

var args []string

type Person struct {
	Id    int
	Name  string
	Email string
}

func (p *Person) WriteTo(output *os.File) error {
	// return os.ErrClosed
}

func main() {
	//Go code
	john := &Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
	}
	output, err := os.Create(args[0])
	if err != nil {
		log.Fatalln("Failed to open file", args[0])
	}
	err = john.WriteTo(output)
	if err != nil {
		log.Fatalln("Failed to write to file", args[0])
	}
}
