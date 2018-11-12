package selection

import (
	"math/rand"

	"github.com/yoonsue/turtles/fitness"
)

func computePerfPopulation(population []string, password string) {
	var populationPerf []float32
	for i, individual := range population {
		populationPerf[i] = fitness.GetFitness(password, individual)
	}
}

func selectFromPopulation(populationSorted [][]string, bestSample []string, luckyFew []string) []string {
	var nextGeneration []string
	for i, _ := range bestSample {
		nextGeneration = append(nextGeneration, populationSorted[i][0])
	}
	for _, _ = range luckyFew {
		nextGeneration = append(nextGeneration, populationSorted[rand.Intn(len(populationSorted))][0])
	}
	// TO BE IMPLEMENTED
	// Population shuffle reversely
	rand.Shuffle(len(nextGeneration), func(i, j int) {
		nextGeneration[i], nextGeneration[j] = nextGeneration[j], nextGeneration[i]
	})
	return nextGeneration
}
