package main

import "fmt"

func main() {

	name := "Александр"

	greeting := makeGreeting(name)
	fmt.Println(greeting)
}

func makeGreeting(name string) string {
	return "Привет, меня зовут " + name
}
