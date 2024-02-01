package main

import "fmt"

type Tools struct{}

func main() {
	tool := Tools{}
	fmt.Println("Generating 12 character long random string  ", tool.RandomString(12))

}
