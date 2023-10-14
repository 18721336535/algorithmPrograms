package main

import "fmt"

type Student struct {
	name string

	age int

	score float64
}

func (std *Student) say() string {
	return fmt.Sprintf("student[name = %v, age =%v, score = %v]",
		std.name, std.age, std.score)
}

type Visitor struct {
	name string
	age  int
}

func (visitor *Visitor) getFee() {
	if visitor.age > 18 {
		fmt.Println("age  > 18, need fee 20")
	} else {
		fmt.Println("age  < 18, free")
	}
}

func Test() {

	var student = Student{
		name:  "tom",
		age:   100,
		score: 99.6,
	}

	str := student.say()

	fmt.Println("stdd--", str)

	var visitor Visitor
	for {
		fmt.Scanln(&visitor.name)
		if "exit" == visitor.name {
			fmt.Println("exit")
			break
		}
		fmt.Scanln(&visitor.age)
		visitor.getFee()

	}

}
