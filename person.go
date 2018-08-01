package main

import (
	"fmt"
	"math"
)

type Person struct {
	name string
	age  int
}

func newPerson() Person {
	return Person{"Diego", 50}
}

func (d Person) AddAge(x, y int) int {
	return x + y
}

func (d Person) SubAge(x, y int) int {
	return x - y
}

func (d Person) GetGoodWeightByHeight(weight, height float64) (float64, float64) {
	result := weight / math.Pow(2, height)
	return height, result
}

func main() {
	p := newPerson()
	//fmt.Println(p.SubAge(p.AddAge(42, 13), 10))
	fmt.Println(p.GetGoodWeightByHeight(70.0, 1.80))
}
