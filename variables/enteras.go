package variables

import "fmt"

func ShowIntegers() {
	var commonInt int
	intDe32 := int32(10)
	intDe64 := int64(100)

	fmt.Println("Common Integer", commonInt)
	fmt.Println("Common Integer 32", intDe32)
	fmt.Println("Common Integer 64", intDe64)
}
