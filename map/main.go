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

	//this is called a comma "ok" idiom
	if v, ok := m["James"]; ok {
		fmt.Println("This will print", v)
	}

	//add to a map
	m["todd"] = 99

	//range over a map
	for k, v := range m {
		fmt.Println("key: ", k)
		fmt.Println("value: ", v)
	}

	//delete in a map
	delete(m, "todd")

	//range over a map
	xi := []int{4, 5, 6, 9, 7, 31}
	for i, v := range xi {
		fmt.Println("index: ", i)
		fmt.Println("value: ", v)
	}

	exercise5()
}

func exercise1() {
	//composite literal
	a := [5]int{1, 3, 4, 5, 10}

	fmt.Printf("type: %T\n", a)

	for _, v := range a {
		fmt.Println("values: ", v)
	}
}

func exercise2() {
	a := []int{0, 1, 2, 3, 5, 6, 7, 8, 9, 10}

	fmt.Printf("type: %T\n", a)

	for i, v := range a {
		fmt.Println("index: ", i)
		fmt.Println("value: ", v)
	}
}

func exercise3() {
	a := []int{0, 1, 2, 3, 5, 6, 7, 8, 9, 10}
	fmt.Println(a[:6])
	fmt.Println(a[6:])
	fmt.Println(a[0:2])
}

func exercise4() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	a = append(a, 52)
	fmt.Println(a)

	a = append(a, 53, 54, 55)
	fmt.Println(a)

	x := []int{56, 57, 58, 59, 60}
	a = append(a, x...)
	fmt.Println(a)
}

func exercise5() {
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	a = append(a[:3], a[6:]...)
	fmt.Println(a)
}
