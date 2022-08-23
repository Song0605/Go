package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	fmt.Println()
	/**/
	var str1, int1, bool1 = "string1", 1, true
	fmt.Println(str1)
	fmt.Println(int1)
	fmt.Println(bool1)
	fmt.Println()
	/**/
	int2 := 2
	string2 := "string2"
	fmt.Println(int2)
	fmt.Println(string2)
	fmt.Println()
	/**/
	int3, _ := foo()
	_, string3 := foo()
	fmt.Println(int3)
	fmt.Println(string3)
	fmt.Println()
	/**/
	const (
		pi = 3.14
		e  = 2.7182
		e1
		e2
	)
	fmt.Println(pi)
	fmt.Println(e)
	fmt.Println(e1)
	fmt.Println(e2)
	/**/
	const (
		n1 = iota
		n2
		n3
		n4
	)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)
	/**/
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
	)
	fmt.Printf("%+v\n", KB)
	fmt.Printf("%+v\n", MB)
	/**/
	people := Human{Name: "zhangsan"}
	fmt.Printf("%+v", people)
	fmt.Printf("%#v", people)
	fmt.Printf("%#v", people)
}

func foo() (int, string) {
	return 3, "string3"
}

type Human struct {
	Name string
}
