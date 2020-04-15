package binarysearch

import (
	//"fmt"
	//"sort"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGenerateData(t *testing.T) {
	res := generateData()

	assert.Equal(t, dataLength, len(res))
}

func TestGetKeys(t *testing.T) {
	testData1 := []int{ 21, 546546, 54645, 4646, 235, 2345, 66767, 434, 899, 5757, 4643, 665}
	testData2 := []int{435, 6767, 45, 342, 66765, 121, 6456, 466, 9080, 78, 567, 16}

	getKeys(testData1)
	resKeysTest1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	getKeys(testData2)
	resKeysTest2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	assert.Equal(t, resKeysTest1, keys)
	assert.Equal(t, resKeysTest2, keys)
}

func generateTestData() {
	data := [][]int{
		{ 12,      5,    87,  987, 45335,  123,   677, 9786,   43, 4565, 67547,   9},
		{435,   6767,    45,  342, 66765,  121,  6456,  466, 9080,   78,   567,  16},
		{ 21, 546546, 54645, 4646,   235, 2345, 66767,  434,  899, 5757,  4643, 665},
	}

	sData = nil
	for _, s := range data {
		sData = append(sData, s...)
	}
}