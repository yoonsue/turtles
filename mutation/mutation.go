package mutation

import (
	"math/rand"
	"strconv"
)

func mutateWord(word string) string {
	indexModification := int(rand.Int() * len(word))
	if indexModification == 0 {
		word = strconv.Itoa(97+int(26*rand.Int())) + word[1:]
	} else {
		word = word[:indexModification] + strconv.Itoa(97+int(26*rand.Int())) + word[indexModification+1:]
	}
	return word
}

func mutatePopulation(population []string, chanceOfMutation int) []string {
	for i := 0; i < len(population); i++ {
		if rand.Int()*100 < chanceOfMutation {
			population[i] = mutateWord(population[i])
		}
	}
	return population
}
