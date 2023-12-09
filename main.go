package main

import (
	"fmt"
	"github.com/godesdecero/variables"
)

func main() {
	//variables.ShowIntegers()
	//variables.RestVariables()
	state, text := variables.TextConvert(123)
	fmt.Println(state, text)
}
