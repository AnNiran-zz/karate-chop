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
	offset := 0

	for step >= 1 {
		
		if target == sData[p] {
			fmt.Println("Found position: ", sData[p])
			return p
		}

		// Check if boundary conditions occured
		if p == 0 || p == len(sData)-1 {
			if sData[p] == target {
				return p
			}
			return -1
		}

		if step > 1 {
			if step%2 == 1 {
				offset = 1
			} else {
				offset = 0
			}

			step = step/2
		} 
		
		if step == 1 {
			offset = 0
		}

		// if target is greater than half value -> move start pointer to step position
		// decrease step with 50%
		if target > sData[p] {
			p += (step+offset)
			continue
		}

		if target < sData[p] {
			p -= (step+offset)
			continue
		}
	}

	return -1
}


// fsDescription outputs the description for "flip-step" binary search implementation
// matches the doc.go file
func FsDescription() string {
	return fmt.Sprint(
`
Flip-step implementation of binary search logic uses a pointer and a stpe with changing size to
search through the list 
At the begining the pointer is set at the middles of the list position and the step size
is set to half length of the list

After comparing the target with the value at position p of the list:
* if the target is smaller than the value at p - the step size is "flipped" to the smaller range
of values in the list
* if the target is greater than the value at p - the step is "flipped" to the range with larger values

at each step the step size is decreased with 50% which decreases the potential range of values
`)
}