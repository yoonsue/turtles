package main

import (
	"fmt"
	"math/rand"
)

func main() {
	sizeList := []int{
		6,
		4, 3, 3, 3, 3, 2,
		4, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 1,
		4, 2, 2, 3, 2, 2, 3, 2, 2, 3, 2, 2, 3, 2, 2, 3, 2, 1,
		4, 2, 2, 2, 3, 2, 2, 2, 3, 2, 2, 2, 3, 2, 2, 2, 3, 2, 2, 2, 3, 2, 2, 1,
	}

	fmt.Printf("Turtles Problem Sloving with GA\n")

	n := 0
	endInt := 0
	fmt.Printf("How many hexagons you want? ")
	fmt.Scanf("%d", &n)

	i := 0
	for i < n {
		endInt += sizeList[i]
		i++
	}
	myrand := random(endInt)

	fmt.Println("size: ", n, "( ~ ", endInt, ")\nlist: ", myrand)
}

func random(size int) []int {
	list := rand.Perm(size)
	for i, _ := range list {
		list[i]++
	}

	randList := make([]int, len(list))
	perm := rand.Perm(len(list))
	for i, v := range perm {
		randList[v] = list[i]
	}
	return randList
}
