package main

import (
	"fmt"
	"strconv"
	"strings"
	"errors"
	"github.com/AnNiran/karate-chop/binarysearch"
)

type commandFunc func (a []string) error

// commands map contains available commands
var commands = map[string]commandFunc {
	"binary-search": cmdBinarySearch,
	// 
}

type algorithmFunc func (t int) int

// algorithms contains list of available algorithms to run from binarysearch package
var algorithms = map[string]algorithmFunc {
	"flip-step":         binarysearch.RunFlippingStep, // fix
	"pointers-move":     binarysearch.RunPointersMoving,
	"pointers-move-int": binarysearch.RunPointersMovingInt,
	"recursive-dynamic": binarysearch.RunRecursiveDynamic,
	"iterate-dynamic":   binarysearch.RunIterateDynamic,
	"parallel-dynamic":  binarysearch.RunParallelDynamic,
	//
}

type algorithmDescFunc func () string

// aldescriptions contain functions calls to each algorithm implementation description
// from the binarysearch package
var aldescriptions = map[string]algorithmDescFunc {
	"flip-step":         binarysearch.FsDescription, // fix
	"pointers-move":     binarysearch.MpDescription,
	"pointers-move-int": binarysearch.MpiDescription,
	"recursive-dynamic": binarysearch.RdDescription,
	"iterate-dynamic":   binarysearch.IdDescription,
	"parallel-dynamic":  binarysearch.PdsDescription,
}

// Errors and outputs
var (
	ErrWrongArgsNum = func(num int) string {
		return fmt.Sprintf("Wrong number of arguments used: %s\n\n", strconv.Itoa(num))
	}
	ErrUnknownCmd = func(cmd string) string {
		return fmt.Sprintf("Unknown command %q\n", cmd)
	}
	ErrUnknownAlgorithm = func(algorithm string) error {
		return errors.New(fmt.Sprintf("Unknown algorithm name provided: %s", algorithm))
	}

	AvailCmdsOutput = func() string {
		return fmt.Sprintf("Available commands are:\n\t%s\n\n", strings.Join(availableCmds(), "\n"))
	}
	RunHelpCmdOutput = func (cmd string) string {
		return fmt.Sprintf("Run '%s <command> -help' to learn more about each command.\n", cmd)
	}
)
