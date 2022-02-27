package main

import "fmt"

func main() {
	const (
		a = 6
		b = a << iota
	)
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	fmt.Printf("%d\t%b\t%#x", b, b, b)
}
