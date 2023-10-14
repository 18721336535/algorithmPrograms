package model

type student struct{
	name string

	age int

	score float64
}

func NewStudent(n string, score float64) *student{
	return &student{
		name : n,
		score :score,
	}
}
