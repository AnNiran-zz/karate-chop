package main

import (
	"fmt"
	"flag"
	"time"
)

// cmdBinarySearch accepts arguments for starting binary search algorithms
// returns found target key or -1 if not found
func cmdBinarySearch(args []string) error {	
	// Define command
	binSrchCmd := flag.NewFlagSet("binary-search", flag.ExitOnError)

	// Define flags
	info      := binSrchCmd.String("info", "", "Displays description for each algorithm")
	algorithm := binSrchCmd.String("algorithm", "", "Algorithm to be used for the search")
	target    := binSrchCmd.Int("target", 0, "Target value to be found")

	// Set up -help output text
	binSrchCmd.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), `
Please enter the type of algorithm for binary search,
max value for set of values to be seached
and the required value

Run example:
binary-search -algorithm <algorithm-name> -target <value-to-search>

Display descriptions example:
binary-search -info <algorithm-name>`,
	)
		binSrchCmd.PrintDefaults()
	}

	binSrchCmd.Parse(args)

	// We have checked arguments count, no we need to check content
	if *info != "" {
		// Remove all non-alphanumeric characters from algorithm name
		trimmedd, err := clearNonAlphNumChars(*info)
		if err != nil {
			return err
		}

		algDesc, ok := aldescriptions[trimmedd]
		if !ok {
			return ErrUnknownAlgorithm(trimmedd)
		}

		fmt.Println(algDesc())
	}

	if *algorithm != "" {
		// Remove all non-alphanumeric characters from algorithm name
		trimmeda, err := clearNonAlphNumChars(*algorithm)
		if err != nil {
			return err
		}

		// Check if provided algorithm name corresponds to any existing algorithm
		alg, ok := algorithms[trimmeda]
		if !ok {
			return ErrUnknownAlgorithm(trimmeda)
		}

		fmt.Println(run(alg, *target))
	}

	return nil
}

// 
func run(a algorithmFunc, target int) int {
	// Set up time for measuring execution
	start := time.Now()

	// Call corresponding algorithm function after all checks are performed
	res := a(target)

	// Print the executed time in milliseconds
	fmt.Println(measureExecTime(start))
	return res
}
