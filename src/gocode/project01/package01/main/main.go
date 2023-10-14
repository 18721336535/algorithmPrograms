package main

import (
	"fmt"
	"gocode/project01/package01/model"
	"gocode/project01/package01/practices"
	"gocode/project01/package01/practices/mapPractice"
)

func escape() *int {
	var a int = 1
	return &a
}

func escapexx() {

	// var a int = 1
	var slice1 = make([]int, 1, 2)
	slice1[0] = 9

	fmt.Println("--", slice1)
}

func main() {
	mapPractice.TestMap01("Tidy")
	model.NewStudent("1", 1.0)
	practices.Test()
	escape()
	escapexx()

	var slice = make([]int, 0, 0)
	slice = append(slice, 100)
	slice = append(slice, 101)
	slice = append(slice, 102)
	slice = append(slice, 103)
	for i, _ := range slice {
		fmt.Printf("--%v,%v", i, slice[i])
	}

}
