package main

import (
	"fmt"
	"math/rand"
	"time"
	t "timelogger"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(200)
	for i := 1; i <= 200; i++ {
		randName := "func" + fmt.Sprintf("%v", i)
		go tester(randName, i, &wg)
	}
	wg.Wait()
}

func tester(name string, i int, wg *sync.WaitGroup) {
	timelog := t.New(name, 2)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
	timelog.Finish()
	wg.Done()
}
