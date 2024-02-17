package variables

import "fmt"

func MuestroEnteros() {
	var intcomun int
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("intcomun será igual a: ", intcomun)
	fmt.Println("intde32 será igual a: ", intde32)
	fmt.Println("intde64 será igual a: ", intde64)
}
