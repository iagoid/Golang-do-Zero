package main

import (
	"fmt"
	"os"
)

func main() {
	// go build main.go
	// main Ola Iago

	fmt.Printf("%#v \n", os.Args)

	fmt.Printf("Path - %#v \n", os.Args[0])
	fmt.Printf("1 - %#v \n", os.Args[1])
	fmt.Printf("2 - %#v \n", os.Args[2])

}
