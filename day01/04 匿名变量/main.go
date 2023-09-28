package main

import "fmt"

func foo() (int, int) {
	return 1, 2
}
func main() {
	var a, _ = foo()
	fmt.Println(a)
}
