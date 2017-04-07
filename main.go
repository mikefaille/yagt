package yagt

import (
	"fmt"
	"log"
	"regexp"
	"runtime"
	"time"
)

func TimeTrack(start time.Time) time.Duration {
	elapsed := time.Since(start)
	return elapsed
}

func Enter() {
	// Skip this function, and fetch the PC and file for its parent
	pc, _, _, _ := runtime.Caller(1)
	// Retrieve a Function object this functions parent
	functionObject := runtime.FuncForPC(pc)
	// Regex to extract just the function name (and not the module path)
	extractFnName := regexp.MustCompile(`^.*\.(.*)$`)
	fnName := extractFnName.ReplaceAllString(functionObject.Name(), "$1")
	fmt.Printf("Entering %s\n", fnName)
	log.Printf("Entering %s\n", fnName)
}

func Exit(start time.Time) {
	// Skip this function, and fetch the PC and file for its parent
	pc, _, _, _ := runtime.Caller(1)
	// Retrieve a Function object this functions parent
	functionObject := runtime.FuncForPC(pc)
	// Regex to extract just the function name (and not the module path)
	extractFnName := regexp.MustCompile(`^.*\.(.*)$`)
	fnName := extractFnName.ReplaceAllString(functionObject.Name(), "$1")
	fmt.Printf("Exiting  %s after %s \n", fnName, TimeTrack(start).String())
	log.Printf("Exiting  %s after %s \n", fnName, TimeTrack(start).String())
}
