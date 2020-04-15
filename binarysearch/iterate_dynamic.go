package binarysearch

import (
	"fmt"
	"sort"
)
// RunIterateDynamic is called externally from cmd package
// generates the data, set up the target values and runs the dynamic iteration logic
func RunIterateDynamic(t int) int {
	generateData()
	target = t
	
	return id()
}

// id runs the dynamic iteration implementation logic of a binary search
// uses offset as a value to move around the list and changes the size of the sData variable
func id() int {
	// Do not enter loop if data length is 0 or 1
	if len(sData) == 0 {
		return -1
	}
			
	if len(sData) == 1 {
		if sData[0] == target {
			return 0
		}
		return -1
	}

	// sort the generated data
	sort.Ints(sData)

	offset := 0

	// While the sData lingth is larger than 0 - we have data to cover
	for len(sData) > 0 {

		p := len(sData) / 2

		if sData[p] == target {
			return p + offset
		}

		if sData[p] > target {
			sData = sData[:p]
		} else {
			sData = sData[p+1:]
			offset += p + 1
		}
	}

	return -1
}

// idDescription outputs the description for the dynamic iteration implementation
// matches the doc.go file text description
func IdDescription() string {
	return fmt.Sprint(
`"resursive-dynamic" implementation of the binary search algoritm uses minimal values settings:
an offset for calculating next range for iteration

At first the offset value is set at 0 and after each iteration, the sData content that is not to be checked
is removed and the offset is updated:
* if the value at the p (pointer) position is smaller than the target - offset is updated to cover half of the 
data plus one position for the remaining part of it
* if the target value is smaller than the value at the p position - the offset keeps its size

sData is dynamically cut at each step while its length decreases => that is how the potential range is smaller each time
and target key is found 
`)
}