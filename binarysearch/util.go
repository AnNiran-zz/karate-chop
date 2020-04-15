package binarysearch

import (
	"math/rand"
	//"sort"

	//"os"
)

// generateData returns a slice of randomly created integers
// max level is 100 000 000, so positions might not be 100% unique
// this is handled in the insertion logic
func generateData() []int {
	sData = make([]int, dataLength)

	for i := range sData {
		sData[i] = rand.Intn(dataRandMax)
	}

	return sData
}


// getKeys gets the keys of the generated data 
// and sets them inside a slice that is updated suring the search
func getKeys(data []int) {
	keys = []int{}

	for i, _ := range data {
		keys = append(keys, i)
	}
}
