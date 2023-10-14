package main

import (
	"fmt"
	"gocode/project01/packege02/temp/objects"
)

func main() {
	var stud = objects.NewStudent("sj", 88.9)
	fmt.Println("std=", *stud)

}
