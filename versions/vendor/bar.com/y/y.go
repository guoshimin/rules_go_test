package y

import "foo.com/x"
import "fmt"

func Bar() {
	fmt.Println(x.Foo2("abc"))
}
