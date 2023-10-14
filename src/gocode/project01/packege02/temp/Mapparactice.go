package main

import (
	"fmt"
	_ "math/rand"

	"sort"
	_ "time"
)

type Studx struct {
	Name    string
	Age     int
	pointer *int
	slice   []int
	mymap   map[string]string
}

func Test1() {

	//map
	mp := make(map[int]string)
	mp[1] = "1"
	mp[2] = "we"

	mp1 := make(map[string]map[string]string)
	mp1["std1"] = make(map[string]string)
	mp1["std1"]["name"] = "tome"
	mp1["std1"]["age"] = "10"
	mp1["std1"]["gender"] = "男"

	mp1["std2"] = make(map[string]string)
	mp1["std2"]["name"] = "adt"
	mp1["std2"]["age"] = "100"
	mp1["std2"]["gender"] = "v"

	fmt.Println("--", mp1)
	fmt.Println("--", len(mp1))

	//map slice
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)

	if monsters[0] == nil {
		monsters[0] = make(map[string]string)
		monsters[0]["name"] = "ss"
		monsters[0]["age"] = "100"
	}
	if monsters[1] == nil {
		monsters[1] = make(map[string]string)
		monsters[1]["name"] = "sss"
		monsters[1]["age"] = "1000"
	}
	fmt.Println("--", monsters)

	monste := map[string]string{
		"name": "d",
		"age":  "50",
	}
	// auto inc map slice
	monsters = append(monsters, monste)
	monsters = append(monsters, map[string]string{
		"name": "dttttttttt",
		"age":  "50",
	})
	fmt.Println("--", monsters)

	// map sort
	map01 := make(map[int]string)
	map01[6] = "x1"
	map01[3] = "x3"
	map01[5] = "x5"
	fmt.Println("-map01-", map01)
	var keys []int

	for key, _ := range map01 {
		keys = append(keys, key)

	}

	sort.Ints(keys)
	fmt.Println("-keys-", keys)

	for _, key := range keys {
		fmt.Println("-key-valeu-", map01[key])

	}

	stdmap := make(map[string]Stud)
	sd1 := Stud{"name1", 18, "sh"}
	sd2 := Stud{"name2", 28, "GZ"}

	stdmap["student1"] = sd1
	stdmap["student2"] = sd2

	fmt.Println("students:", stdmap)

	for key, v := range stdmap {
		fmt.Printf("\n student no:%v", key)
		fmt.Printf("\n student name:%v", v.name)
		fmt.Printf("\n student age:%v", v.age)
		fmt.Printf("\n student address:%v", v.address)
		fmt.Println("------------")

	}

	stdents2 := make(map[string]map[string]string)
	stdents2["n1"] = make(map[string]string)
	stdents2["n1"]["pwd"] = "555"
	stdents2["n1"]["nickname"] = "jk"

	fmt.Println("-----modifysd-first------", stdents2)

	modifysd(stdents2, "n1")
	fmt.Println("-----modifysd-------", stdents2)

	fmt.Println("-----******************------")

	var stx Studx
	fmt.Println("-----Studx-------", stx)
	stx.Age = 100

	fmt.Println("-----Studx-------", stx)
	intxx := 888
	stx.pointer = &intxx

	fmt.Println("-----Studx.pointer-------", stx)
	stx.slice = make([]int, 3)

	stx.slice[0] = 99
	fmt.Println("-----Studx-slice------", stx)
	stx.mymap = make(map[string]string)

	stx.mymap["name1"] = "cat"
	fmt.Println("-----Studx-mymap------", stx)
}

func modifysd(stds map[string]map[string]string, name string) {
	if stds[name] != nil {
		stds[name]["pws"] = "000"
	} else {
		stds[name] = make(map[string]string)
		stds[name]["pwd"] = "99"
		stds[name]["nkname"] = "andy"
	}
}

type Stud struct {
	name    string
	age     int
	address string
}
