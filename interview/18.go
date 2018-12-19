//是否可以编译通过？如果通过，输出什么？
package main

import (
	"fmt"
)

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}

	fmt.Println("non-empty interface")
}

func main() {
	var x *int = nil
	Foo(x)
}
