package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	MaxCapacity   = 30
	TotalStudents = 100
)

var startTime = time.Now()

func student(wg *sync.WaitGroup, lib chan int, id int) {
	defer wg.Done()

	rand.Seed(time.Now().UnixNano())
	readingTime := time.Duration(rand.Intn(4) + 1)

	lib <- id
	fmt.Printf("Time %d: Student %d begins reading at the library\n", int(time.Since(startTime).Seconds()), id)

	time.Sleep(readingTime * time.Second)

	<-lib
	fmt.Printf("Time %d: Student %d departs. Spent %d hours reading\n", int(time.Since(startTime).Seconds()), id, readingTime)
}

func main() {
	lib := make(chan int, MaxCapacity)
	var wg sync.WaitGroup

	for id := 1; id <= TotalStudents; id++ {
		wg.Add(1)
		go student(&wg, lib, id)
	}

	wg.Wait()

	fmt.Printf("Time %d: No more students. End of the day\n", int(time.Since(startTime).Seconds()))
	fmt.Printf("The library needs to be open for %d hours\n", int(time.Since(startTime).Seconds()))
}
