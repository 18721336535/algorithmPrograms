package temp

import "fmt"

func Addup() func(int) int {
	var n int = 10
	var str = "hello"
	return func(x int) int {
		n = n + x
		str += string(36)
		fmt.Println(str)
		return n
	}
}

func Test() {
	ft := Addup()
	fmt.Println(ft(3))
	fmt.Println(ft(3))
	fmt.Println("hello world!!!")
}
