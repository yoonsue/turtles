package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Printf("Turtles Problem Sloving with GA\n")

	n := 0
	endInt := 0
	fmt.Printf("How many hexagons you want? ")
	fmt.Scanf("%d", &n)

	listOfNumOfVerticesToAdd := newVerticesOfHexGenerator(n)

	i := 0
	for i < n {
		endInt += listOfNumOfVerticesToAdd[i]
		i++
	}
	myrand := random(endInt)

	randomRoundCheck := roundCheck(n)

	fmt.Println("size of List: ", len(listOfNumOfVerticesToAdd), "listOfNumOfVerticesToAdd: ", listOfNumOfVerticesToAdd)
	fmt.Println("round #", randomRoundCheck, "size: ", n, "( ~ ", endInt, ")\nlist: ", myrand)
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

func roundCheck(currentHexNum int) int {
	an := 1
	roundNum := 1

	for {
		lastHexOfCurrentRound := an + (roundNum * (6 * (roundNum - 1)) / 2)
		// fmt.Println("last Hex of Round #", roundNum, " is Hex #", lastHexOfCurrentRound)

		if currentHexNum > lastHexOfCurrentRound {
			roundNum++
		} else {
			break
		}
	}
	return roundNum
}

func newVerticesOfHexGenerator(hexN int) []int {
	var listOfNumOfVerticesToAdd []int

	currentHexNum := 1

	an := 1
	roundNum := 1
	lastHexOfCurrentRound := an + (roundNum * (6 * (roundNum - 1)) / 2)

	numOfVerticesToAdd := 0

	for currentHexNum <= hexN {
		tmpBetweenTwoVertexOfHex := roundNum - 2

		if currentHexNum == lastHexOfCurrentRound+1 { // first vertex for current ROUND
			numOfVerticesToAdd = 4
			roundNum++
			lastHexOfCurrentRound = an + (roundNum * (6 * (roundNum - 1)) / 2)
		} else if currentHexNum == lastHexOfCurrentRound { // last vertex for current ROUND
			if roundNum == 1 { // ROUND1 last vertex
				numOfVerticesToAdd = 6
			} else if roundNum == 2 { // ROUND2 last vertex
				numOfVerticesToAdd = 2
			} else { // last vertex (more than 2 ROUND)
				numOfVerticesToAdd = 1
			}
		} else {
			for tmpBetweenTwoVertexOfHex > 0 {
				numOfVerticesToAdd = 2
				listOfNumOfVerticesToAdd = append(listOfNumOfVerticesToAdd, numOfVerticesToAdd)
				fmt.Println("Round #", roundNum, " Hex#", currentHexNum, "lastHex#", lastHexOfCurrentRound, "numOfVerticesToAdd: ", numOfVerticesToAdd)
				tmpBetweenTwoVertexOfHex--
				currentHexNum++
				if currentHexNum == lastHexOfCurrentRound {
					break
				}
			}
			if currentHexNum != lastHexOfCurrentRound {
				numOfVerticesToAdd = 3
			} else {
				numOfVerticesToAdd = 1
			}
			fmt.Println("Round #", roundNum, " Hex#", currentHexNum, "lastHex#", lastHexOfCurrentRound, "numOfVerticesToAdd: ", numOfVerticesToAdd)
		}
		fmt.Println("Round #", roundNum, " Hex#", currentHexNum, "lastHex#", lastHexOfCurrentRound, "numOfVerticesToAdd: ", numOfVerticesToAdd)
		listOfNumOfVerticesToAdd = append(listOfNumOfVerticesToAdd, numOfVerticesToAdd)

		currentHexNum++
	}
	return listOfNumOfVerticesToAdd
}

func sumOfVerticesInHex(randList []int, currentHexNum int, incrementOfInt int) int {

	tmpIntList := make([]int, 0, 6)

	sum := 0

	if currentHexNum == 1 { // Round 1.
		tmpIntList = randList[:6]
	} else { // From Round 2.
		if incrementOfInt == 4 { // Round start-point
			tmpIntList = randList
		} else if incrementOfInt == 3 { // 6 vertices

		} else if incrementOfInt == 2 { // rest of edges

		} else if incrementOfInt == 1 { // Round end-point (Start at Round 3)

		}
	}

	for i, _ := range tmpIntList {
		sum += tmpIntList[i]
	}

	return sum

}
