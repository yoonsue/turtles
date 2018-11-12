package breeding

import "math/rand"

func createChild(individual1 string, individual2 string) string {
	child := ""
	for i := 0; i < len(individual1); i++ {
		if int(100*rand.Int()) < 50 {
			child += string(individual1[i])
		} else {
			child += string(individual2[i])
		}
	}
	return child
}

func createChildren(breeders []string, numberOfChild int) []string {
	var nextPopulation []string
	for i := 0; i < (len(breeders) / 2); i++ {
		for j := 0; j < numberOfChild; j++ {
			nextPopulation = append(nextPopulation, createChild(breeders[i], breeders[len(breeders)-1-i]))
		}
	}
	return nextPopulation
}
