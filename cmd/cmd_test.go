package main

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

var cmdInit     = "./cmd"
var cmdString   = "binary-search"
var targetFl    = "-target"
var algorithmFl = "-algorithm"

var algNames = make([]string, len(algorithms))

//
func TestCmdBinarySearchUnknownAlgorithm(t *testing.T) {
	for k := range algorithms {
		algNames = append(algNames, k)
	}
	algNames = algNames[3:]
	testArgs := [][]string{
		{
			targetFl,
			"21",
			algorithmFl,
			fmt.Sprintf("%s-test", algNames[4]),
		},
		{
			targetFl,
			"12",
			algorithmFl,
			fmt.Sprintf("%s-test", algNames[5]),
		},
		{
			targetFl,
			"567",
			algorithmFl,
			fmt.Sprintf("%s-test", algNames[6]),
		},
	}

	for n, test := range testArgs {
		res := cmdBinarySearch(test)

		expectedRes := ErrUnknownAlgorithm(test[3])
		assert.Equal(t, expectedRes, res, fmt.Sprintf("Failed at test %d", n))
	}
}

//
func TestCmdBinarySearchProperAlgorithm(t *testing.T) {
	for k := range algorithms {
		algNames = append(algNames, k)
	}

	algNames = algNames[3:]
	testArgs := [][]string{
		{
			targetFl,
			"21",
			algorithmFl,
			algNames[0],
		},
		{
			targetFl,
			"12",
			algorithmFl,
			algNames[1],
		},
		{
			targetFl,
			"567",
			algorithmFl,
			algNames[2],
		},
	}

	// algorithms contains list of available algorithms to run from binarysearch package
	var algorithmsTest = make(map[string]algorithmFunc)

	for name, _ := range algorithms {
		algorithmsTest[name] = func(target int) int { return -1}
	}

	for n, test := range testArgs {
		res := cmdBinarySearch(test)
		assert.Equal(t, nil, res, fmt.Sprintf("Failed at test %d", n))
	}
}
