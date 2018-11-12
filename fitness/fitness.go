package fitness

import (
	"fmt"
)

type Fitness struct {
	fit float32
}

func GetFitness(password string, testWord string) float32 {
	if len(testWord) != len(password) {
		fmt.Printf("length difference between test word and password\n")
		return -1
	} else {
		score := 0
		for i := range password {
			if password[i] == testWord[i] {
				score += 1
			}
		}
		return float32(score * 100 / len(password))
	}
}
