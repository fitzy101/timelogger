## Summary
Timelogger is a package designed to log the amount of time a function call
takes to run.

There are 2 log levels.
Level 1 states the time from start to finish.
Level 2 states level 1 plus the number of current go routines running (at the time of the Finish() call).

## Installation

```sh
go get github.com/fitzy101/timelogger
```

Below are a few simple examples:

### Level 1
### Level 2