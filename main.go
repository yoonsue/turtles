package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"time"
)

func main() {

	fmt.Printf("Turtles Problem Sloving with GA\n")

	n := 0
	totalVertex := 0
	fmt.Printf("How many hexagons you want? ")
	fmt.Scanf("%d", &n)

	populationNum := 200
	replaceNum := 20
	// generationNum := 10
	mutationRate := 5

	startTime := time.Now()

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	listOfNumOfVerticesToAdd := newVerticesOfHexGenerator(n)

	i := 0
	for i < n {
		totalVertex += listOfNumOfVerticesToAdd[i]
		i++
	}

	population := makePopulation(populationNum, totalVertex)
	// fmt.Println("population: ", population)

	count := 0
	// for count <= generationNum {
	for {
		count++

		j := 0
		for j < replaceNum {
			tmpN1 := r1.Intn(populationNum)
			parent1 := population[tmpN1]

			tmpN2 := r1.Intn(populationNum)
			for tmpN2 == tmpN1 {
				tmpN2 = r1.Intn(populationNum)
			}
			parent2 := population[tmpN2]
			// fmt.Println("parent1: ", parent1)
			// fmt.Println("parent2: ", parent2)

			children := cycleCrossover(parent1, parent2, totalVertex)
			// fmt.Println("children: ", children)

			children, _ /* isMutationOccur :*/ = mutation(mutationRate, children, totalVertex)
			// if isMutationOccur == true {
			// 	fmt.Println("mutatedChild: ", children)
			// }

			population = replacement(population, tmpN1, tmpN2, children, listOfNumOfVerticesToAdd, n)

			j++
		}

		bestLoc, _ := getBestWorstGene(population, listOfNumOfVerticesToAdd, n)

		if count%5000 == 0 {
			fmt.Printf("GENERATION #%d best gene: %v\n", count, population[bestLoc])
		}

		if isSolution(listOfNumOfVerticesToAdd, population[bestLoc], n) {
			fmt.Println("SOLUTION FOUND!!!")
			answerShow(listOfNumOfVerticesToAdd, population[bestLoc], n)

			endTime := time.Now()
			totalTime := endTime.Sub(startTime)
			fmt.Println("COMPUTATION TIME: ", totalTime)
			break
		}

	}

}

func isSolution(listOfNumOfVerticesToAdd []int, gene []int, n int) bool {
	i := 1
	sumOfVerticesInFirstHex := sumOfVerticesInHex(listOfNumOfVerticesToAdd, gene, i)
	for i <= n {
		if sumOfVerticesInFirstHex != sumOfVerticesInHex(listOfNumOfVerticesToAdd, gene, i) {
			return false
		}
		i++
	}
	return true
}

func getBestWorstGene(population [][]int, listOfNumOfVerticesToAdd []int, n int) (int, int) {
	i := 0

	bestFitness := fitness(listOfNumOfVerticesToAdd, population[i], n)
	worstFitness := fitness(listOfNumOfVerticesToAdd, population[i], n)

	bestLoc := 0
	worstLoc := 0

	for i < len(population) {
		fitnessOfGene := fitness(listOfNumOfVerticesToAdd, population[i], n)
		if fitnessOfGene < bestFitness {
			bestFitness = fitnessOfGene
			bestLoc = i
		}
		if fitnessOfGene > worstFitness {
			worstFitness = fitnessOfGene
			worstLoc = i
		}
		i++
	}

	return bestLoc, worstLoc
}

func replacement(population [][]int, parent1Loc int, parent2Loc int, children []int, listOfNumOfVerticesToAdd []int, n int) [][]int {
	newPopulation := population

	fitnessOfParent1 := fitness(listOfNumOfVerticesToAdd, population[parent1Loc], n)
	fitnessOfParent2 := fitness(listOfNumOfVerticesToAdd, population[parent2Loc], n)
	fitnessOfChildren := fitness(listOfNumOfVerticesToAdd, children, n)

	// changedGeneNum := 0
	if fitnessOfChildren < fitnessOfParent1 {
		newPopulation[parent1Loc] = children
		// changedGeneNum = parent1Loc
	} else if fitnessOfChildren < fitnessOfParent2 {
		newPopulation[parent2Loc] = children
		// changedGeneNum = parent2Loc
	} else {
		/// OPTION1: randomly change gene
		// randNum := rand.Intn(len(population))
		// newPopulation[randNum] = children
		// changedGeneNum = randNum

		/// OPTION2: worst gene changed
		_, worstLoc := getBestWorstGene(population, listOfNumOfVerticesToAdd, n)
		newPopulation[worstLoc] = children
	}
	// fmt.Println("changedGeneNum: ", changedGeneNum)
	return newPopulation
}

// mutationRate is multiplied by 100 --> So, its type is integer
func mutation(mutationRate int, children []int, totalVertex int) ([]int, bool) {
	flag := false
	i := 0
	for i < totalVertex {
		if rand.Intn(100) < mutationRate {
			j := rand.Intn(totalVertex)
			tmp := children[i]
			children[i] = children[j]
			children[j] = tmp
			flag = true
		}
		i++
	}

	return children, flag
}

func randInit() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// better when result value is similar to 0
func answerShow(listOfNumOfVerticesToAdd []int, gene []int, n int) {
	i := 1
	for i <= n {
		averageOfVerticesInHex := (float64)(sumOfVerticesInHex(listOfNumOfVerticesToAdd, gene, i)) / 6
		fmt.Printf("%2.1f\t", averageOfVerticesInHex)
		i++
	}
	fmt.Println()

	fmt.Println("SOLUTION for ", n, " hexagonal ", gene)

}

// better when result value is similar to 0
func fitness(listOfNumOfVerticesToAdd []int, gene []int, n int) float64 {

	// sumOfVerticesInHex := 0
	averageOfVerticesInHex := 0.0
	i := 1
	for i <= n {
		averageOfVerticesInHex += (float64)(sumOfVerticesInHex(listOfNumOfVerticesToAdd, gene, i)) / 6
		i++
	}
	averageOfVerticesInHex = averageOfVerticesInHex / float64(n)

	dispersionOfData := 0.0
	for i <= n {
		dispersionOfData += math.Pow(float64(sumOfVerticesInHex(listOfNumOfVerticesToAdd, gene, i))-averageOfVerticesInHex, 2)
		i++
	}
	dispersionOfData = dispersionOfData / float64(n)

	// deviation := math.Sqrt(dispersionOfData)

	// absFitness := abs(averageOfVerticesInHex - deviation*math.Sqrt(float64(n)))
	return dispersionOfData
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// average sum of vertex in all hexagonal
func averageSumOfHex(n int) float64 {
	averageOfSum := 0.0

	i := 1
	for i <= n {
		averageOfSum += float64(i)
		i++
	}

	averageOfSum /= float64(n)

	return averageOfSum
}

func selection(population [][]int, listOfNumOfVerticesToAdd []int, currentHexNum int) []int {

	populationNum := len(population)

	// point := 0 // TO BE CHANGED
	bestFitness := fitness(listOfNumOfVerticesToAdd, population[0], currentHexNum)
	worstFitness := bestFitness
	bestSelectedNum := 0
	// worstSelectedNum := 0
	i := 0
	for i < populationNum {
		tmpFitness := fitness(listOfNumOfVerticesToAdd, population[i], currentHexNum)
		if tmpFitness < bestFitness {
			bestFitness = tmpFitness
			bestSelectedNum = i
		} else if tmpFitness > worstFitness {
			worstFitness = tmpFitness
			// worstSelectedNum = i
		}
		i++
	}
	return population[bestSelectedNum]
}

func cycleCrossover(parent1 []int, parent2 []int, totalVertexNum int) []int {

	children := make([]int, totalVertexNum)

	i := 0
	for i < totalVertexNum {
		checkNum := parent1[i]

		if parent1[i] != parent2[i] {
			j := 0
			for j < totalVertexNum {
				if checkNum == parent2[j] {
					if children[j] == checkNum {
						break
					}
					children[j] = checkNum
					checkNum = parent1[j]
					j = 0
				} else {
					j++
				}
			}
			break
		}
		i++
	}

	i = 0
	for i < totalVertexNum {
		if children[i] == 0 { // not created
			children[i] = parent1[i]
		}
		i++
	}

	return children
}

// 1 point crossover
func crossover(parent1 []int, parent2 []int, totalVertexNum int) []int {

	cuttingPoint := rand.Intn(totalVertexNum)

	children := []int{}

	children = append(children, parent1[:cuttingPoint]...)
	children = append(children, parent2[cuttingPoint:]...)

	fmt.Println("cuttingPoint: ", cuttingPoint)

	fmt.Println(children)

	return children
}

/// Make number of population
func makePopulation(populationNum int, totalVertexNum int) [][]int {
	population := [][]int{}

	i := 0
	for i < populationNum {
		gene := random(totalVertexNum)
		j := 0
		for j < i {
			if reflect.DeepEqual(population[j], gene) {
				gene = random(totalVertexNum)
				j--
			}
			j++
		}
		population = append(population, gene)
		// fmt.Println("gene #", i, population[i])
		i++
	}

	return population
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
	sumInt -= listOfNumOfVerticesToAdd[i-1]

	tmpIntList := make([]int, 6)
	sumVertexList := make([]int, 6)

	sum := 0
	incrementVertex := listOfNumOfVerticesToAdd[currentHexNum-1]

	if currentHexNum == 1 { // Round 1.
		// fmt.Println("ROUND 1")
		tmpIntList = randList[:6]

		sumVertexList[0] = currentVertexForHex
		sumVertexList[1] = currentVertexForHex + 1
		sumVertexList[2] = currentVertexForHex + 2
		sumVertexList[3] = currentVertexForHex + 3
		sumVertexList[4] = currentVertexForHex + 4
		sumVertexList[5] = currentVertexForHex + 5
	} else if roundNum == 2 {
		// fmt.Println("ROUND 2, incrementVertex: ", incrementVertex)
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

	// fmt.Println("sumVertexList: ", sumVertexList)
	for i, _ := range tmpIntList {
		sum += tmpIntList[i]
	}

	return sum

}
