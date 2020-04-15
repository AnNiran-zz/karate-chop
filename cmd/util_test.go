package main

import (
	"os"
	"io"
	"fmt"
	"log"
	"sync"
	"bytes"

	"testing"
	"github.com/stretchr/testify/assert"
)

// Test case when wrong number of arguments are provided
func TestCheckArgsCountWrongNum(t *testing.T) {
	testArgsWrongNum := [][]string{
		{
			"Test wrong number of arguments 1",
			cmdInit,
			"test-args",
			"help",
		},
		{
			"Test wrong number of arguments 2",
			cmdInit,
			cmdString,
			targetFl,
			algorithmFl,
			"algorithm-test-name",
		},
		{
			"Test wrong number of arguments 3",
			cmdInit,
			cmdString,
			algorithmFl,
			"algorithm-test-name",
			targetFl,
			"test-value",
			"test-value",
		},
	}

	expectedOutput := func(args []string) {
		fmt.Printf(ErrWrongArgsNum(len(args)-1))
		fmt.Printf(AvailCmdsOutput())
		fmt.Printf(RunHelpCmdOutput(args[0]))
	}

	for _, test := range testArgsWrongNum {
		expOutput := captureOutput(func() {
			expectedOutput(test[1:])
		})

		expErrorText := ErrWrongArgsNum(len(test)-2)

		var errOutput error
		resOuput := captureOutput(func() {
			errOutput = checkArgsCount(test[1:])
		})

		assert.Equal(t, expOutput, resOuput, fmt.Sprintf("Failed at: %s", test[0]))
		assert.Equal(t, expErrorText, errOutput.Error(), fmt.Sprintf("Failed at: %s", test[0]))
	}
}

// Test case when -help flag should be used
func TestCheckArgsCountHelpFlag(t *testing.T) {
	testArgsHelpFlag := [][]string{
		{
			"Test arguments with help flag 1",
			cmdInit,
			"test-args",
			"-help",
		},
		{
			"Test arguments with help flag 2",
			cmdInit,
			"test-args",
			"-help",
		},
	}

	for _, test := range testArgsHelpFlag {
		var errOutput error
		resOuput := captureOutput(func() {
			errOutput = checkArgsCount(test[1:])
		})

		assert.Equal(t, "", resOuput, fmt.Sprintf("Failed at: %s", test[0]))
		assert.Equal(t, nil, errOutput, fmt.Sprintf("Failed at: %s", test[0]))
	}
}

// Test case when correct number of arguments is provided
func TestCheckArgsCountCorrectNum(t *testing.T) {
	testArgs := [][]string{
		{
			"Test arguments 1",
			cmdInit,
			cmdString,
			algorithmFl,
			"test-alogirthm-name",
			targetFl,
			"test-target-value-int",
		},
		{
			"Test arguments 2",
			cmdInit,
			"test-command",
			"test-flag-algorithm",
			"test-algorithm-name",
			"test-flag-target",
			"test-target-value-int",
		},
		{
			"Test arguments 3",
			cmdInit,
			cmdString,
			"-info",
			"algorithm-test-name",
		},
	}

	for _, test := range testArgs {
		var errOutput error
		resOuput := captureOutput(func() {
			errOutput = checkArgsCount(test[1:])
		})

		assert.Equal(t, "", resOuput, fmt.Sprintf("Failed at: %s", test[0]))
		assert.Equal(t, nil, errOutput, fmt.Sprintf("Failed at: %s", test[0]))
	}
}

// Test cases for characters that need to be cleared
func TestClearNonAlphNumCharsClear(t *testing.T) {
	testInput := []string{
		"a!@#$^&*()_$&346", 
		"test-very-long-string-input58934^&*%$75389457", 
		"!@#$^&target-test-input*()_$&346",
		"a!@#$^&*()_$&346-test-input",
		"--test-input",
	}
	
	resInput, err := clearNonAlphNumChars(testInput[0])
	expectedTestInput1 := "a346"

	assert.Equal(t, expectedTestInput1, resInput, fmt.Sprint("Failed at testInput 1"))
	assert.Equal(t, nil, err, fmt.Sprint("Failed at testInput 1"))

	resInput, err = clearNonAlphNumChars(testInput[1])
	expectedTestInput2 := "test-very-long-string-input5893475389457"
	assert.Equal(t, expectedTestInput2, resInput, fmt.Sprint("Failed at testInput 2"))
	assert.Equal(t, nil, err, fmt.Sprint("Failed at testInput 2"))

	resInput, err = clearNonAlphNumChars(testInput[2])
	expectedTestInput3 := "target-test-input346"
	assert.Equal(t, expectedTestInput3, resInput, fmt.Sprint("Failed at testInput 3"))
	assert.Equal(t, nil, err, fmt.Sprint("Failed at testInput 3"))

	resInput, err = clearNonAlphNumChars(testInput[3])
	expectedTestInput4 := "a346-test-input"
	assert.Equal(t, expectedTestInput4, resInput, fmt.Sprint("Failed at testInput 4"))
	assert.Equal(t, nil, err, fmt.Sprint("Failed at input 4"))

	resInput, err = clearNonAlphNumChars(testInput[4])
	expectedTestInput5 := "test-input"
	assert.Equal(t, expectedTestInput5, resInput, fmt.Sprint("Failed at testInput 5"))
	assert.Equal(t, nil, err, fmt.Sprint("Failed at input 5"))

}

// Test cases for no characters to be cleared from input
func TestClearNonAlphNumCharsNotClear(t *testing.T) {
	testInput := []string{
		"algorithm-test-input",
		"algorithm-test-input-9458074968", 
	}
	
	resInput, err := clearNonAlphNumChars(testInput[0])
	assert.Equal(t, testInput[0], resInput, fmt.Sprint("Failed at tesInput 1"))
	assert.Equal(t, nil, err, fmt.Sprint("Failed at testInput 1"))

	resInput, err = clearNonAlphNumChars(testInput[1])
	assert.Equal(t, testInput[1], resInput, fmt.Sprint("Failed at testInput 2"))
	assert.Equal(t, nil, err, fmt.Sprint("Failed at testInput 2"))
}

// Captute output from os.Stderr that needs to be checked across the package functions
// for testing
func captureOutput(f func()) string {
	// Initialize os.Pipeline - creates a pipe btewee nreader and writer with *os.File type
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	// Replace standard os output and err with pipeline
	// and set the output for a log package
	// log.Setouput changes the address of the output object
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()

	os.Stdout = writer
	os.Stderr = writer 
	log.SetOutput(writer)

	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)

	// We create a new goroutine because read and write cannot stay in one
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()

	writer.Close()
	return <- out
}
