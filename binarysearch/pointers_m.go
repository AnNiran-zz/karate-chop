package binarysearch

import (
	"fmt"
	"sort"
)

// RunMovingPointers is called externally from cmd package
// generates data and set up the target value, then runs the algorithm
func RunPointersMoving(t int) int {
	generateData()
	target = t

	return mp()
}

// mp runs the binary-search logic using 3 pointers:
// * starting pointer - initially set at the first list position
// * ending pointer - initially set at the last list position
// * searching pointer - initially set at the half of the list length position
// while searching each pointer moves towards a smaller potential range for searching the target
// description of the logic is inside doc.go file
func mp() int {
	sort.Ints(sData)

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

	// Set up starting pointer
	ps := 0

	// Set up ending pointer
	pe := len(sData)-1

	// Set up pointer
	p := len(sData)/2
	for len(sData[ps:pe]) > 1 {

		// First check if met value equals target
		if target == sData[p] {
			return p
		}

		// Check boundary conditions => there are 3 moving points across the slice => we need to check them
		// this check handles the condition where ps equals pe as well
		for pos := range []int{ps, pe} {
			if target == sData[pos] {
				return pos
			}

		}
		
		// If target is greater than half value -> move start pointer to step position
		// decrease step with 50%
		if target > sData[p] {
			ps = p

			// Cut next move size with 50%
			p = ps + len(sData[ps:pe])/2
			continue
		}

		if target < sData[p] {
			pe = p

			// Cut next move size with 50%
			p = ps + len(sData[ps:pe])/2
			continue
		}

		return -1
	}

	return -1
}

// mpDescription outputs the description of the logic for running the moving pointer binary search
// matches the text inside doc.go
func MpDescription() string {
	return fmt.Sprint(
`"pointers-move" implementation for binary search algorithm searches through the sorted 
array of values for the given target using three pointers:
starting, ending and searching pointer - that are used to move accross the list after each binary check 

At the first step:
* starting pointer is set at the first position of the list, ending pointer - 
at the last position and the searching pointer is set at the position at the half length of the list

- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
|                                   |                                  |
ps                                  p                                  pe

After comparing the list value at the searching pointer position with the target:
* if the target is greater - the searching pointer is moved towards the right with 
half of its previous value
* starting pointer is moved at the previous searching pointer position
* the ending pointer keeps its current position

- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
|                                   |                                    |
ps ->                               p ->                                 pe
									
									ps                 p                 pe
									   
* if the target is smaller than the searching pointer positioin
* the searching pointer is moved towards the smallest range with half of its previous value 
* the ending pointer is moved at the previous position of the searching pointer
* the starting pointer keeps its position

- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
|                                   |                                    |
ps ->                               p ->                                 pe
									
									ps                 p              <- pe

									ps        p        pe


At each loop the searching pointer key is cut with 50% and the range between starting and ending pointers decreases

If no value is found -1 is returned
`)
}