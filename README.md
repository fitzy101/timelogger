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
```go
package main

import (
	"fmt"
	t "github.com/fitzy101/timelogger"
	"time"
	"math/rand"
)

func main() {
	// Create a new timelogger instance to record the function duration.
	timelog := t.New("foo", 1)
	
	// Sleep for a random amount of time.
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	
	// Call the Finish function on the timelogger to set the elapsed time to
	// be completed now. This returns the log in a string.
	fmt.Println(timelog.Finish())
}
```

```sh
go run main.go
# => Total time for foo: 83.152045ms
```

### Level 2
```go
package main

import (
	"fmt"
	t "github.com/fitzy101/timelogger"
	"time"
	"math/rand"
)

func main() {
	// Create a new timelogger instance to record the function duration.
	timelog := t.New("foo", 2)
	
	// Sleep for a random amount of time.
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	
	// Call the Finish function on the timelogger to set the elapsed time to
	// be completed now. This returns the log in a string.
	fmt.Println(timelog.Finish())
}
```

```sh
go run main.go
# => Total time for foo: 82.820102ms
# => Number of current goroutines: 5
```
