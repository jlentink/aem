package output

import (
	"fmt"
	"strings"
	"time"
)

const (
	// NORMAL Output in normal mode
	NORMAL = iota

	// VERBOSE Output in verbose mode
	VERBOSE
)

var logLevel = NORMAL

// SetLevel Set log level (NORMAL, VERBOSE)
func SetLevel(level int) {
	logLevel = level
}

// SetVerbose set verbose level
func SetVerbose(v bool) {
	if v {
		SetLevel(VERBOSE)
	}
}

// Print formats using the default formats for its operands and writes to standard output.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
// Only displays when log level is high enough
func Print(l int, o string) (int, error) {
	if logLevel >= l {
		return fmt.Print(o)
	}
	return 0, nil
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
// Only displays when log level is high enough
func Printf(l int, o string, a ...interface{}) (int, error) {
	if logLevel >= l {
		return fmt.Printf(o, a...)
	}
	return 0, nil
}

// PrintListing formats a array into a list
func PrintListing(key, valuesString string, trail bool) {
	values := strings.Split(valuesString, ",")
	fmt.Printf("- %s:\n", key)
	for i, value := range values {
		fmt.Printf("\t %d) %s\n", i+1, strings.TrimSpace(value))
	}
	if trail {
		fmt.Printf("\n")
	}
}

// UnixTime Format time
func UnixTime(timestamp int64) *time.Time {
	if timestamp == 0 {
		return nil
	}
	tm := time.Unix(timestamp/1000, 0)
	theTime := tm.UTC()
	return &theTime
}
