package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Printf("Turtles Problem Sloving with GA\n")

	n := 0
	endInt := 0
	fmt.Printf("How many hexagons you want? ")
	fmt.Scanf("%d", &n)

	startTime := time.Now()

	listOfNumOfVerticesToAdd := newVerticesOfHexGenerator(n)

	i := 0
	for i < n {
		endInt += listOfNumOfVerticesToAdd[i]
		i++
	}
	myrand := random(endInt)

	randomRoundCheck, HexCheck := roundCheck(n)

	fmt.Println("size of List: ", len(listOfNumOfVerticesToAdd), "listOfNumOfVerticesToAdd: ", listOfNumOfVerticesToAdd)
	fmt.Println("round #", randomRoundCheck, "Hex #", HexCheck, "size: ", n, "( ~ ", endInt, ")\nlist: ", myrand)

	sumOfVerticesInHex := sumOfVerticesInHex(listOfNumOfVerticesToAdd, myrand, n)

	fmt.Println("sum of vertices in hexagonal: ", sumOfVerticesInHex)
	endTime := time.Now()

	totalTime := endTime.Sub(startTime)
	fmt.Println("COMPUTATION TIME: ", totalTime)
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

func roundCheck(currentHexNum int) (int, int) {
	an := 1
	roundNum := 1

	lastHexOfPrevRound := 0
	// lastHexOfCurrentRound := an
	for {
		lastHexOfCurrentRound := an + (roundNum * (6 * (roundNum - 1)) / 2)
		// fmt.Println("last Hex of Round #", roundNum, " is Hex #", lastHexOfCurrentRound)

		if currentHexNum > lastHexOfCurrentRound {
			lastHexOfPrevRound = lastHexOfCurrentRound
			roundNum++
		} else {
			break
		}
	}

	hexNumOfRound := currentHexNum - lastHexOfPrevRound

	return roundNum, hexNumOfRound
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
				// fmt.Println("Round #", roundNum, " Hex#", currentHexNum, "lastHex#", lastHexOfCurrentRound, "numOfVerticesToAdd: ", numOfVerticesToAdd)
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
			// fmt.Println("Round #", roundNum, " Hex#", currentHexNum, "lastHex#", lastHexOfCurrentRound, "numOfVerticesToAdd: ", numOfVerticesToAdd)
		}
		// fmt.Println("Round #", roundNum, " Hex#", currentHexNum, "lastHex#", lastHexOfCurrentRound, "numOfVerticesToAdd: ", numOfVerticesToAdd)
		listOfNumOfVerticesToAdd = append(listOfNumOfVerticesToAdd, numOfVerticesToAdd)

		currentHexNum++
	}
	return listOfNumOfVerticesToAdd
}

func getStartVertexForHex(listOfNumOfVerticesToAdd []int, hexNum int) int {
	i := 0
	startVertexForHex := 0
	for i < hexNum-1 {
		startVertexForHex += listOfNumOfVerticesToAdd[i]
		i++
	}
	return startVertexForHex
}

func sumOfVerticesInHex(listOfNumOfVerticesToAdd []int, randList []int, currentHexNum int) int {

	roundNum, hexNum := roundCheck(currentHexNum)

	edgeNum := 0
	vertexOfEdge := 0
	if roundNum != 1 {
		edgeNum = (hexNum - 1) / (roundNum - 1)
		vertexOfEdge = (hexNum - 1) % (roundNum - 1)
	}

	currentVertexForHex := getStartVertexForHex(listOfNumOfVerticesToAdd, currentHexNum)

	previousRoundVertexNum := 1
	if hexNum != 1 {
		currentVertexForHex--
	}

	previousRoundVertexNum += currentVertexForHex - edgeNum*2 - 6*(1+2*(roundNum-2))
	if vertexOfEdge != 0 {
		previousRoundVertexNum -= 2
	}

	twoRoundBeforeEndHexNum := 0
	previousRoundEndHexNum := 0
	currentRoundEndHexNum := 1

	if roundNum > 2 {
		currentRoundEndHexNum += (roundNum * (6 * (roundNum - 1)) / 2)
		previousRoundEndHexNum = currentRoundEndHexNum - (roundNum-1)*6
		twoRoundBeforeEndHexNum = previousRoundEndHexNum - (roundNum-2)*6
	}

	sumInt := 0
	i := 0
	previousRoundStartVertex := 0
	for i < currentHexNum {
		sumInt += listOfNumOfVerticesToAdd[i]
		if i < twoRoundBeforeEndHexNum {
			previousRoundStartVertex += listOfNumOfVerticesToAdd[i]
		}
		i++
	}

	j := 0
	currentRoundStartVertex := 0
	for j < roundNum-1 {
		currentRoundStartVertex += 6 * (1 + 2*j)
		j++
	}
	previousRoundEndVertex := currentRoundStartVertex - 1
	fmt.Println("previousRoundStartVertex: ", previousRoundStartVertex, "currentRoundStartVertex: ", currentRoundStartVertex, "previousRoundEndVertex: ", previousRoundEndVertex)
	sumInt -= listOfNumOfVerticesToAdd[i-1]

	tmpIntList := make([]int, 6)
	sumVertexList := make([]int, 6)

	sum := 0
	incrementVertex := listOfNumOfVerticesToAdd[currentHexNum-1]

	if currentHexNum == 1 { // Round 1.
		fmt.Println("ROUND 1")
		tmpIntList = randList[:6]

		sumVertexList[0] = currentVertexForHex
		sumVertexList[1] = currentVertexForHex + 1
		sumVertexList[2] = currentVertexForHex + 2
		sumVertexList[3] = currentVertexForHex + 3
		sumVertexList[4] = currentVertexForHex + 4
		sumVertexList[5] = currentVertexForHex + 5
	} else if roundNum == 2 {
		fmt.Println("ROUND 2, incrementVertex: ", incrementVertex)
		if incrementVertex == 4 { // Round start-point
			tmpIntList[0] = randList[currentVertexForHex]
			tmpIntList[1] = randList[currentVertexForHex+1]
			tmpIntList[2] = randList[currentVertexForHex+2]
			tmpIntList[3] = randList[currentVertexForHex+3]

			tmpIntList[4] = randList[previousRoundVertexNum]
			tmpIntList[5] = randList[previousRoundVertexNum+1]

			sumVertexList[0] = currentVertexForHex
			sumVertexList[1] = currentVertexForHex + 1
			sumVertexList[2] = currentVertexForHex + 2
			sumVertexList[3] = currentVertexForHex + 3
			sumVertexList[4] = previousRoundVertexNum
			sumVertexList[5] = previousRoundVertexNum + 1

		} else if incrementVertex == 3 { // 6 vertices
			tmpIntList[0] = randList[currentVertexForHex]
			tmpIntList[1] = randList[currentVertexForHex+1]
			tmpIntList[2] = randList[currentVertexForHex+2]
			tmpIntList[3] = randList[currentVertexForHex+3]

			tmpIntList[4] = randList[previousRoundVertexNum]
			tmpIntList[5] = randList[previousRoundVertexNum+1]

			sumVertexList[0] = currentVertexForHex
			sumVertexList[1] = currentVertexForHex + 1
			sumVertexList[2] = currentVertexForHex + 2
			sumVertexList[3] = currentVertexForHex + 3
			sumVertexList[4] = previousRoundVertexNum
			sumVertexList[5] = previousRoundVertexNum + 1
		} else if incrementVertex == 2 { // Round2 end-point
			tmpIntList[0] = randList[currentVertexForHex]
			tmpIntList[1] = randList[currentVertexForHex+1]
			tmpIntList[2] = randList[currentVertexForHex+2]

			tmpIntList[3] = randList[previousRoundStartVertex]
			tmpIntList[4] = randList[previousRoundStartVertex+1]
			tmpIntList[5] = randList[currentRoundStartVertex]

			sumVertexList[0] = currentVertexForHex
			sumVertexList[1] = currentVertexForHex + 1
			sumVertexList[2] = currentVertexForHex + 2
			sumVertexList[3] = previousRoundVertexNum
			sumVertexList[4] = previousRoundVertexNum + 1
			sumVertexList[5] = currentRoundStartVertex
		}

	} else { // From Round 3.

		if incrementVertex == 4 { // Round start-point
			tmpIntList[0] = randList[currentVertexForHex]
			tmpIntList[1] = randList[currentVertexForHex+1]
			tmpIntList[2] = randList[currentVertexForHex+2]
			tmpIntList[3] = randList[currentVertexForHex+3]

			tmpIntList[4] = previousRoundVertexNum
			tmpIntList[5] = previousRoundVertexNum + 1

			sumVertexList[0] = currentVertexForHex
			sumVertexList[1] = currentVertexForHex + 1
			sumVertexList[2] = currentVertexForHex + 2
			sumVertexList[3] = currentVertexForHex + 3
			sumVertexList[4] = previousRoundVertexNum
			sumVertexList[5] = previousRoundVertexNum + 1

		} else if incrementVertex == 3 { // 6 vertices
			tmpIntList[0] = randList[currentVertexForHex]
			tmpIntList[1] = randList[currentVertexForHex+1]
			tmpIntList[2] = randList[currentVertexForHex+2]
			tmpIntList[3] = randList[currentVertexForHex+3]

			tmpIntList[4] = previousRoundVertexNum
			tmpIntList[5] = previousRoundVertexNum + 1

			sumVertexList[0] = currentVertexForHex
			sumVertexList[1] = currentVertexForHex + 1
			sumVertexList[2] = currentVertexForHex + 2
			sumVertexList[3] = currentVertexForHex + 3
			sumVertexList[4] = previousRoundVertexNum
			sumVertexList[5] = previousRoundVertexNum + 1

		} else if incrementVertex == 2 { // rest of edges
			tmpIntList[0] = randList[currentVertexForHex]
			tmpIntList[1] = randList[currentVertexForHex+1]
			tmpIntList[2] = randList[currentVertexForHex+2]

			tmpIntList[3] = randList[previousRoundVertexNum]
			tmpIntList[4] = randList[previousRoundVertexNum+1]
			tmpIntList[5] = randList[previousRoundVertexNum+2]

			sumVertexList[0] = currentVertexForHex
			sumVertexList[1] = currentVertexForHex + 1
			sumVertexList[2] = currentVertexForHex + 2
			sumVertexList[3] = previousRoundVertexNum
			sumVertexList[4] = previousRoundVertexNum + 1
			sumVertexList[5] = previousRoundVertexNum + 2

		} else if incrementVertex == 1 { // Round end-point (Start at Round 3)
			tmpIntList[0] = randList[currentVertexForHex]
			tmpIntList[1] = randList[currentVertexForHex+1]

			tmpIntList[2] = randList[previousRoundStartVertex]
			tmpIntList[3] = randList[previousRoundStartVertex+1]
			tmpIntList[4] = randList[previousRoundEndVertex]
			tmpIntList[5] = randList[currentRoundStartVertex]

			sumVertexList[0] = currentVertexForHex
			sumVertexList[1] = currentVertexForHex + 1
			sumVertexList[2] = previousRoundStartVertex
			sumVertexList[3] = previousRoundStartVertex + 1
			sumVertexList[4] = previousRoundEndVertex
			sumVertexList[5] = currentRoundStartVertex
		}
	}

	fmt.Println("sumVertexList: ", sumVertexList)

	for i, _ := range tmpIntList {
		sum += tmpIntList[i]
	}

	return sum

}
