package binarysearch

// Values used by all functions running binary search algorithm
var dataRandMax = 90000
var dataLength  = 160

// sData is used to keep slice of integers to work on
var sData []int

// dData (dynamic data) and keys slices are used by the logic implementations
// that are dynamically changing the lists of values when searching
var dData []int
var keys  []int

// target is used to hold the target set for the running function
var target int

// pointer used across functions
var po int 

// ***
// Define pointers that will search through the data range
// used by the two implementations using the moving pointers
var p  pointer
var ps pointer
var pe pointer

// Type pointer and pointerUpdater interface are used 
// by implementation of moving pointers with interface 
type pointer int

type pointerUpdater interface {
	updateSearchPointer(pointer)
}