package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {

	person := &Person{
		Name: "Tarcisio Rocha",
		Age:  41,
	}

	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Println("Tamanho: ", len(data))

	newPerson := &Person{}
	err = proto.Unmarshal(data, newPerson)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Println("Idade: ", newPerson.GetAge())
	fmt.Println("Nome: ", newPerson.GetName())

}
