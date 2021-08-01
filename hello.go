package main

import (
	"fmt"
	"log"

	"github.com/wataru-nakanishi/go-graphql/greetings"
)

func main() {
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
    fmt.Println(messages)
}