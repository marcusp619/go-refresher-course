package main

import "fmt"

func main() {
	m := map[string]int{
		"James": 32,
		"Mark":  34,
	}

	fmt.Print(m)

	//validate if value exists in map
	v, ok := m["Mark"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["James"]; ok {
		fmt.Println("This will print", v)
	}
}
