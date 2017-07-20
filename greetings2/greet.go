package main

import (
	"fmt"
	"os"

	"github.com/vladimirvivien/getting-started-with-go/greetlib"
)

func main() {
	lang := "English"
	if len(os.Args) >= 2 {
		lang = os.Args[1]
	}
	fmt.Println(greetlib.GreetIn(lang))
}
