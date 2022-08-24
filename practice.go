package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func prac() {
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
	fmt.Printf("%+v\n", people)
	fmt.Printf("%#v\n", people)
	fmt.Printf("%#v\n", people)
	/**/
	var (
		int4 = 4
		int5 = 05
		int6 = 0xff
	)
	fmt.Printf("%d\n", int4)
	fmt.Printf("%o\n", int5)
	fmt.Printf("%x\n", int6)
	fmt.Printf("%X\n", int6)
	/**/
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	/**/
	var (
		c1 = 1 + 2i
		c2 = 2 + 3i
	)
	fmt.Println(c1)
	fmt.Println(c2)
	/**/
	string4 := `第一行
第二行
第三行`
	string5 := "Sstring5S"
	fmt.Println(string4)
	fmt.Println(len(string4))
	fmt.Println(strings.HasPrefix(string5, "S"))
	fmt.Println(strings.HasSuffix(string5, "S"))
	/**/
	string6 := "Hello云建信"
	for i := 0; i < len(string6); i++ { //byte 无法处理中日文
		fmt.Printf("%v(%c)", string6[i], string6[i])
	}
	fmt.Println()
	for _, r := range string6 { //rune 可以处理中日文
		fmt.Printf("%v(%c)", r, r)
	}
	fmt.Println()
	/**/
	string7 := "string7"
	byteS1 := []byte(string7)
	byteS1[0] = 'S'
	fmt.Println(string(byteS1))
	string8 := "白菜"
	runeS2 := []rune(string8)
	runeS2[0] = '生'
	fmt.Println(string(runeS2))
	/**/
	var int_7, int_8 = 3, 4
	var int_9 = int(math.Sqrt(float64(int_7*int_7 + int_8*int_8)))
	fmt.Println(int_9)
	/**/
	string9 := "hello百旺弘祥"
	count := 0
	for _, v := range string9 {
		if unicode.Is(unicode.Han, v) {
			count++
			fmt.Println(string(v))
		}
	}
	fmt.Println(count)
	/**/
	string10 := "string10"
	switch string10 {
	case "string10":
		fmt.Println(string10)
		fallthrough
	case "string9":
		fmt.Println("string9")
		goto gotoTag
	default:
		fmt.Println("...")
	}
gotoTag:
	fmt.Println("goto")
	/**/
	for i := 1; i < 10; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Printf("|%d * %d = %d", j, i, i*j)
			if j != 1 && (i*j)/10 == 0 {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

func foo() (int, string) {
	return 3, "string3"
}

type Human struct {
	Name string
}
