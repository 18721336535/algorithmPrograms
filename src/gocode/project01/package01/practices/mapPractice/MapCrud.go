package mapPractice

import (
	"fmt"
	"strconv"
)

var i int

func TestMap01(name string) {
	mp1 := make(map[string]map[string]string)
	mp1["Andy"] = make(map[string]string)
	mp1["Andy"]["id"] = "001"
	mp1["Andy"]["age"] = "18"
	mp1["Andy"]["phone"] = "18xxxx"

	if mp1[name] != nil {
		mp1[name]["age"] = "19"
	} else {
		mp1[name] = make(map[string]string)
		i++
		mp1[name]["id"] = "name" + strconv.Itoa(i)
		mp1[name]["age"] = "20"
	}
	fmt.Println(mp1)
}
