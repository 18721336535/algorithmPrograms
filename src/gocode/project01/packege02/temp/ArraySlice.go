package main

import (
	"fmt"
	"math/rand"
	"time"
)

func testary(ary *[3]int) {
	(*ary)[0] = 99
	fmt.Printf("\ntestary --%p, %p , %p", &ary, ary, &*ary)
}

func Test() {

	var ary [5]int
	for i := 0; i < 3; i++ {
		// fmt.Scanln(&ary[i])
		fmt.Printf("please input next value\n")
	}

	for j := 0; j < len(ary); j++ {
		fmt.Printf("ary[%d] = %v\t", j, ary[j])
	}

	// var ary2 = [...]int{2,3,5,8,9}

	// ary3 := [...]int{2,3,5,8,9}

	ary3 := [3]string{"f", "r", "j"}
	for index, value := range ary3 {
		fmt.Printf("ary3[%d] = %v\t", index, value)
	}

	ary01 := [3]int{11, 33, 55}

	testary(&ary01)
	fmt.Println("\nmain ary01 ", ary01)
	fmt.Printf("\n%p ", &ary01)

	ary11 := [...][2]int{{1, 3}, {5, 6}}
	fmt.Printf("ary11=%v", ary11[1])

	for i := 0; i < len(ary11); i++ {
		for j := 0; j < len(ary11[i]); j++ {
			fmt.Printf("\nary11=%v", ary11[i][j])
		}
	}

	for i, v := range ary11 {
		for j, v2 := range v {
			fmt.Printf("\nar11[%v][%v] = %v", i, j, v2)
		}
	}

	//binary search
	fmt.Printf("\n")
	ay := [8]int{5, 6, 7, 88, 90, 100, 200, 500}
	// fmt.Scanln(&expectVal)
	binary(&ay, 0, len(ay)-1)

	scores := [...][2]int{{66, 77}, {88, 89}, {99, 60}}
	// for i := 0; i< len(scores) ; i++{
	// 	for j := 0; j < len(scores[i]) ; j++ {
	// 		fmt.Printf("please input class %v,student %v score\n",i,j)
	// 		fmt.Scanln(&scores[i][j])
	// 	}
	// }
	// sumscores := [3]int{0,0}
	total := 0
	for dx, v := range scores {
		sumscores := 0
		for dx1, v2 := range v {
			// sumscores[dx] += v2
			fmt.Printf("scores[%v][%v]= %v", dx, dx1, v2)
			sumscores += v2
		}
		// fmt.Printf("\n class %v avrage score %v",dx,float32(sumscores[dx])/float32(len(v)))
		fmt.Printf("\n class %v avrage score %v", dx,
			float32(sumscores)/float32(len(v)))

		total += sumscores
	}

	// fmt.Printf("\n sum scores %v",sumscores)
	fmt.Printf("\n total scores %v", total)

	aryExercise()

}

var expectVal int = 0

func binary(ay *[8]int, leftIdex int, rightIdex int) {
	mindleIdx := (leftIdex + rightIdex) / 2
	if leftIdex > rightIdex {
		fmt.Println("\n not find value")
		return
	}
	if ay[mindleIdx] == expectVal {
		fmt.Printf("find value %v", expectVal)
	}
	if ay[mindleIdx] < expectVal {
		binary(ay, (mindleIdx + 1), rightIdex)
	}
	if ay[mindleIdx] > expectVal {
		binary(ay, leftIdex, (mindleIdx - 1))
	}

}

func aryExercise() {
	rand.Seed(time.Now().UnixNano())
	var nums [10]int
	for i := 0; i < 10; i++ {
		nums[i] = rand.Intn(101)
	}
	fmt.Printf("\n %v", nums)

	for i := len(nums) - 1; i >= 0; i-- {
		fmt.Printf("\n reverse prit %v", nums[i])
	}

	min := nums[0]
	mindex := 0
	maxdex := 0
	max := nums[0]
	for dx, v := range nums {
		if v < min {
			min = v
			mindex = dx
		}
		if v > max {
			max = v
			maxdex = dx
		}
	}
	fmt.Printf("%v, %v, %v, %v", min, mindex, max, maxdex)

	//

}
