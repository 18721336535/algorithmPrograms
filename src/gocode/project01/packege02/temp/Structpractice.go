package main

import (
	"encoding/json"
	"fmt"
)

type Person0 struct {
	Name    string
	Age     int
	Address string
}

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func testA(a Person) {
	fmt.Println(a.Name)
}
func (a Person) testB() {
	fmt.Println("testB--", a.Name)
}

func (a *Person) testC() {
	fmt.Println("testB--", a.Name)
	(*a).Name = "cattt"
	fmt.Println("testB--", (*a).Name)
}

func (a *Person) String() string {
	return fmt.Sprintf("[---name=%v,age=%v]", a.Name, a.Age)
}

func Test03() {
	person0 := Person0{"ay0", 5001, "SH+"}

	strbytes0, error0 := json.Marshal(person0)
	if error0 != nil {
		fmt.Println("-json exist error-", error0)
	}
	fmt.Println("--", string(strbytes0))

	person := Person{"ay", 500, "SH"}
	fmt.Println("--", person)
	strbytes, error := json.Marshal(person)
	if error != nil {
		fmt.Println("-json exist error-", error)
	}
	fmt.Println("--", string(strbytes))
	testA(person)

	person.testB()

	fmt.Println("-mod--person-", person)
	person.testC() //等价下面
	(&person).testC()
	fmt.Println("-mod--person-", person)

	//如果实现了Person 类型的 string 方法 Printf会自动调用String 方法
	fmt.Println(&person)

	var cal Caculator
	cal.Num1 = 1.0
	cal.Num2 = 2.0

	fmt.Printf("-caculator add :%v", fmt.Sprintf("%.2f", cal.addNum()))

	fmt.Printf("-caculator add :%v", fmt.Sprintf("%.2f", cal.getResult('+')))

}

type Caculator struct {
	Num1     float64
	Num2     float64
	Operator byte
}

func (cal *Caculator) addNum() float64 {
	return cal.Num1 + cal.Num2

}

func (cal *Caculator) subNum() float64 {
	return cal.Num1 - cal.Num2

}

func (cal *Caculator) mulNum() float64 {
	return cal.Num1 * cal.Num2

}

func (cal *Caculator) divNum() float64 {
	return cal.Num1 / cal.Num2

}
func (cal *Caculator) getResult(Operator byte) float64 {
	res := 0.0
	switch Operator {
	case '+':
		res = cal.Num1 + cal.Num2
	case '/':
		res = cal.Num1 / cal.Num2
	default:
		fmt.Printf("error iprraor")
	}
	return res

}
