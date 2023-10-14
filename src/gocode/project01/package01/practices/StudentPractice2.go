package practices

import (
	"fmt"
	"gocode/project01/package01/model"
)

func Cest01() {
	var stud = model.NewStudent("sj", 88.9)
	fmt.Println("std=", *stud)

}
