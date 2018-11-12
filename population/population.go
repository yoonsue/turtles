package population

import (
	"math/rand"
	"strconv"
)

func generateAWord(length int) string {
	result := ""
	for i := 0; i < length; i++ {
		letter := strconv.Itoa(97 + int(26*rand.Intn(100)))
		result += letter
	}
	return result
}

func generateFirstPopulation(sizePopulation int, password string) []string {
	var population []string
	for i := 0; i < sizePopulation; i++ {
		population = append(population, generateAWord(len(password)))
	}
	return population
}
