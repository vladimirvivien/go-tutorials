package main

import (
	"fmt"

	"github.com/vladimirvivien/hellolib"
)

func main() {
	greetings := hellolib.Hello("English", "Kreyol", "Swahili")
	fmt.Printf("greetings in English=%s and Shwahili=%s\n", greetings[0], greetings[1])
}
