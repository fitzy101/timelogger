package timelogger

import (
	"fmt"
	"time"
	"runtime"
)

type Logger struct {
	start    time.Time     // The start time of the log.
	elapsed  time.Duration // The total elapsed time.
	logLevel int           // The log detail level.
	funcName string        // The name of the function being loggerd.
}

// New is the public interface with the timelogger package, which returns an
// instance of a logger starting at the current time.
func New(name string, logLevel int) *Logger {
	nL := new(Logger)
	nL.funcName = name
	nL.start = time.Now()
	
	// If the logLevel provided was 0, use a default level of 1. This determines
	// the detail for what information is included in the log.
	if logLevel == 0 {
		nL.logLevel = 1 
	} else {
		nL.logLevel = logLevel
	}
	return nL
}

// Finish calculates the elapsed time, and prints to the provided io.Writer.
func (l *Logger) Finish() string {
	l.setElapsed()
	return retLog(l)
}



// Current returns the current elapsed time in a formatted string.
func (l *Logger) Elapsed() string {
	l.setElapsed()
	return fmt.Sprintf("Total time for %v: %v", l.funcName, l.elapsed)
}

// Sets the elapsed time to the start minus now called before logging.
func (l *Logger) setElapsed() {
	l.elapsed = time.Now().Sub(l.start)
}

// Gets the number of current goroutines running.
func numRoutines() int {
	return runtime.NumGoroutine()
}

// writeLog finishes off the logger, and prints the output to stdout (depending
// on the log level).
func (l *Logger) WriteLog() {
	switch l.logLevel {
	case 1:
		fmt.Println(level1(l))
	case 2:
		fmt.Println(level2(l))
	}
	return
}

// retLog finishes off the logger, and returns the output as a string.
func retLog(l *Logger) string {
	switch l.logLevel {
	case 1:
		return level1(l)
	case 2:
		return level2(l)
	default:
		return level1(l)
	}
}

// This is the basic information for log level 1.
func level1(l *Logger) string {
	str := fmt.Sprintf("Total time for %v: %v\n", l.funcName, l.elapsed)
	return str
}

// This is the basic information for log level 2 - included level 1 plus 
// number of goroutines.
func level2(l *Logger) string {
	str := level1(l) + fmt.Sprintf("Number of current goroutines: %v\n", numRoutines())
	return str
}