package main

import (
	"os"
	"fmt"
)

func main() {
	// Manually handle arguments number
	if err := checkArgsCount(os.Args); err != nil {
		os.Exit(2)
	}

	run, ok := commands[os.Args[1]]
	if !ok {
		fmt.Fprintf(os.Stderr, ErrUnknownCmd(os.Args[1]))
		fmt.Fprintf(os.Stderr, AvailCmdsOutput())
		os.Exit(2)
	}

	// Proper command is called, proceed with it
	if err := run(os.Args[2:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
