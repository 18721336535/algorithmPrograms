package main

import (
	"fmt"
	"unsafe"
)

var z1, z2, z3 = "sre", 10, 11.01

var (
	y1         = 0
	f1, f2, f3 = "sre", 10, 11.01
)

func f10(n1 int, args ...int) int {
	sum := n1
	for i := n1; i < len(args); i++ {
		sum += args[i]

	}
	fmt.Println("sum=", sum)
	return sum

}
func f11(int1 int) {

}

func Test() {
	var x1, x2, x3 int
	var z1, z2, z3 = "sre", 10, 11.01
	t1, t2, t3 := 100, "bt", 10.11
	fmt.Printf("t1=%v,t2 = %v,t3=%v\n", x1, x2, x3)
	fmt.Printf("t1=%v,t2 = %v,t3=%v\n", z1, z2, z3)
	fmt.Printf("t1=%v,t2 = %v,t3=%v\n", t1, t2, t3)

	var j = 0
	var sum = 0
	for i := 0; i < 100; i++ {
		if i%9 == 0 {
			j++
			sum += i
		}

	}
	var n int = 6
	for i1 := -1; i1 < n; i1++ {
		fmt.Println("\t hello \n")

		fmt.Printf("%v + %v = %v \n", i1, n-i1, n)
	}

	fmt.Printf("count=%T:sum = %d", j, unsafe.Sizeof(sum))

	for i := 1; i < 5; i++ {
		fmt.Printf("\n%d", i)
	}
	var xstr string = "Hello中国"
	for i := 0; i < len(xstr); i++ {
		fmt.Printf("%c", xstr[i])
	}

	for index, value := range xstr {
		fmt.Printf("\nindex=%d,value= %c", index, value)

	}

	var xstr1 string = "Hello中国"
	xstr2 := []rune(xstr1)
	for i := 0; i < len(xstr2); i++ {
		fmt.Printf("\n%d,%c", i, xstr2[i])
	}

	f10(10, 21, 1, 1, 1, 1, 1, 3, 2, 2, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 1, 2)

	var n20, n30 int64
	fmt.Printf("%T,%T", n20, n30)

	// 外层循环9次， 内层循环：每一次内层循环 多执行一次
	for i := 1; i < 10; i++ {
		fmt.Println("")
		for j := 1; j <= i; j++ {
			fmt.Printf("%v + %v = %v ,", i, j, i+j)
		}
	}

	// 外层循环9次， 内层循环：每一次内层循环 多执行一次
	for i := 1; i < 10; i++ {
		fmt.Println("")
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v ,", i, j, i*j)
		}
	}

}
