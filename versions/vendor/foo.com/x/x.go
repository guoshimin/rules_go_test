package x

import "fmt"

func Foo1(x int) int {
	fmt.Println("directly vendored")
	return x
}

const foo = 2
