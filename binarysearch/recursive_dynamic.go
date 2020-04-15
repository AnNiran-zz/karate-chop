package binarysearch

import (
	"fmt"
	"sort"
)

// RunRecursiveDynamic is called externally from cmd package
// generates the data, set up the target and runs the algorithm
func RunRecursiveDynamic(t int) int {
	generateData()
	target = t

	return rd()
}

// rd runs the logic implementation for binary search
// that uses separate slices of data from the generated one
// and cut parts of them during the search
func rd() int {
	// Do not start the logic if the generated data size
	// is one or zero
	if len(sData) == 0 {
		return -1
	}

	if len(sData) == 1 {
		if sData[0] == target {
			return 0
		}

		return -1
	}

	// Sort generated data
	sort.Ints(sData)

	// Get sData keys in separate slice, as well as 
	// copy the data inside dData that is going to be changed
	dData = sData
	getKeys(dData)

	po = len(sData)
	return getKeyrd()
}

// 
func getKeyrd() int {
	if po > 1 {
		po = po/2
	}

	// Check if data sized had became one and there 
	// is no need to proceed
	if len(dData) == 1 {
		if dData[0] == target {
			return keys[0]
		}
		return -1
	}

	// Check each case and calls recursively if needed
	switch {
	case target == dData[po]:
		return keys[po]
		
	case target < dData[po]:
		dData = dData[:po]
		keys  = keys[:po]
		return getKeyrd()

	case target > dData[po]:
		dData = dData[po+1:]
		keys  = keys[po+1:]
		return getKeyrd()

	default:
		return -1
	}

	return -1
}

// rdDescription contains description of the dynamic recursive implementation
// matches decription in doc.go file
func RdDescription() string {
	return fmt.Sprint(
`"recursive-dynamic" logic implements binary search algorithm by using 
separate data lists (slices) for keys of the originally generated data, as well as 
a copy of the data and a pointer

Initially the copy of the generated data contains all of it, as well as the keys slice 
contains all keys, the pointer is set at the half of the original data lenght position

- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -  keys
- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -  copy data
										|
										p

At each step when the value at the pointer position is compared to the target
If the target is greater than the value at p:
* the pointer is moved to towards the greater values
* length of the data and keys lists are cut with 50% of their lengths

                                        - - - - - - - - - - - - - - - - - - - - - - -  keys
                                        - - - - - - - - - - - - - - - - - - - - - - -  copy data
										|
										p ->                      |
																  p
											
If the target is smaller than the value at p:
* the pointer is moved towards the smaller values with half of its previous size
* data copy and keys parts after the previous pointer position are removed

                                        - - - - - - - - - - - - - -                    keys
                                        - - - - - - - - - - - - - -                    copy data
                                        |
                                        p ->                      |
						                            |          <- p
													p
													
After each step the size of keys and data copy decreases as well as the potential range values
for searcging
When the target is met - the value at the p from keys list is returned, otherwise -1`)
}
