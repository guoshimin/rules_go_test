package x

import "fmt"

func Foo2(x string) string {
	fmt.Println("nested in the vendor")
	return x
}

const foo = 1
