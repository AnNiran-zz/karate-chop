package binarysearch

import (
	"fmt"
	"sort"
)

// RunPointersMovingInt is called externally from cmd package
// generates the data and set up the target value, then runs
// the pointers-moving implmentation with interfaces
func RunPointersMovingInt(t int) int {
	generateData()
	target = t

	return int(mpi())
}

// updateSearchPointer function updates the pointer
// type
func (p *pointer) updateSearchPointer() {
	*p = ps+pointer(len(sData[ps:pe])/2)
}

// mpi runs the implementation of binary search using moving 
// pointers across the list of values and using interface
func mpi() pointer {
	// Do not proceed with logic if data length is zero or one
	if len(sData) == 0 {
		return -1
	}

	if len(sData) == 1 {
		if sData[0] == target {
			return 0
		}

		return -1
	}

	// Sort list of values
	sort.Ints(sData)

	// Set up start pointer value position
	ps = 0

	// Set up end pointer value position
	pe = pointer(len(sData) - 1)

	// Set up searching pointer position
	p = pointer(pe/2)

	return getKeympi()
}

//
func getKeympi() pointer {

	if len(sData[ps:pe]) > 1 {
		// Check boundary conditions => there are 3 moving points across the slice => we need to check them
		// this check handles the condition where ps equals pe as well
		for pos := range []pointer{ps, pe} {
			if target == sData[pos] {
				return pointer(pos)
			}
		}
		
		if target == sData[p] {
			return p
		}

		if target < sData[p] {
			pe = p
			p.updateSearchPointer()
				
			return getKeympi()
		}

		if target > sData[p] {
			ps = p
			p.updateSearchPointer()
				
			return getKeympi()
		}
	}

	return -1
}

// mpiDescription outputs description for moving-pointers implementation
// with interface; matches the doc.go file
func MpiDescription() string {
	return fmt.Sprint(
`moving-pointers implementation with interface functions in the 
same way as the "moving-pointers", except that is uses custom defined
types - pointer for values of the three pointers, and a in interface
that the type implements
this is done only for prepresentational purposes and is not considerably important
for the code performance and resources optimisation
`)
}