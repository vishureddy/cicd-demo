package main

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

func main() {
	message := hello("")
	fmt.Println(aurora.Blue(message))
}
