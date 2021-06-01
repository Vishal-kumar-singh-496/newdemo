package main

import (
	"fmt"
	"math"
)

type Student struct {
	avgmarks float64
	name     string
	roll     int
}

func main() {

	var x []int = []int{56, 89, 96, 95, 38, 45}

	var y []int = x[1:3]

	avg(&x)

	avg(&y)

}

func avg(x *[]int) {

	var u float64 = 0
	var l float64 = 0

	for i, value := range *x {

		u = u + float64(value)

		i++
	}

	l = u / float64(len(*x))

	l = math.Round(l)

	fmt.Println("closest no to avergae is", l)

	Structrial(l)
}

func Structrial(a float64) {

	var s1 Student = Student{a, "xyz", 1}

	fmt.Println(s1)

}
