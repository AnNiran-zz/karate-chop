package binarysearch

import (
	"fmt"
	"sort"
)

func RunFlippingStep(t int) int {
	// Generate data
	generateData()
	
	// Set up target value
	target = t

	return flipStep(target)
}

func flipStep(target int) int {
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

	// Set up pointer position that will be moved across the list
	p := len(sData)/2
	
	// Set up step size that will be resized according to each 
	// pointer comparison with values from the data list
	step := len(sData)/2

	for step >= 1 {
		if step > 1 {
			step = step/2
		}
		
		if step == 1 {
			// Check for conditions of internal loop every time step is 1
				
			// We need a few conditions to be present in order infinite looping to occur
			// Step must be 1
			// p is moving with 1 between two values and non of them equals target
			// Since we reached this point value at p does not equal target
			// since p value might be changed in previous cases we need 
			// to check if p reached the 0-th or the len(sData)-1 position as well
			if p == 0 {
				if target > sData[p] && target < sData[p+1] {
					return -1
				}
			}

			if p == len(sData)-1 {
				if target < sData[p] && target > sData[p-1] {
					return -1
				}
			}

			if p > 0 && p < len(sData)-1 {
				fmt.Println(sData[p])
				if target > sData[p] && target < sData[p+1] {
					return -1
				}

				if target < sData[p] && target > sData[p-1] {
					return -1
				}
			}
		}

		switch {
		case target == sData[p]:
			return p

		case target > sData[p]:
			if p == len(sData)-1 {
				return -1
			}

			p += step
		
		case target < sData[p]:
			if p == 0 {
				return -1
			}

			p -= step
		default:
			return -1
		}
	}

	return -1
}


// fsDescription outputs the description for "flip-step" binary search implementation
// matches the doc.go file
func FsDescription() string {
	return fmt.Sprint(
`
"flip-step" implementation of binary search logic uses a pointer and a step with changing size to
search through the list
At the begining the pointer is set at the middle list position and the step size is set to half length of the list

After comparing the target with the value at position p of the list:
* if the target is smaller than the value at p:
the step size is cut with 50% if it is larger than 1
step is "flipped" to the smaller range of values in the list

p is moved at the new position - decreasing its key with the new step size

* if the target is greater than the value at p:
step size is cut with 50% if it is larger than 1
the step is "flipped" to the range with larger values in the list

At each step the step size is decreased with 50% which decreases the potential range of values
to be searched
`)
}