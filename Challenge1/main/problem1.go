package main

import (
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
)

func problem1() {

	log.Printf("problem1: started --------------------------------------------")
	var wg sync.WaitGroup
	var counter uint64 //this variable will be used as a concurrent counter between goroutines

	//
	// Todo:
	//
	// Quit all go routines after
	// a total of exactly 100 random
	// numbers have been printed.
	//
	// Do not change the 25 in loop!
	//

	for inx := 0; inx < 10; inx++ {
		go printRandom1(inx, &wg, &counter)
		wg.Add(1) //increment the waitgroup counter, 1 for each goroutine
	}
	wg.Wait() // block the execution until waitgroup counter is again 0

	//
	// Todo:
	//
	// Remove this quick and dirty sleep
	// against a synchronized wait until all
	// go routines are finished.
	//

	log.Printf("problem1: finished --------------------------------------------")
}

func printRandom1(slot int, wg *sync.WaitGroup, counter *uint64) {

	//
	// Do not change 25 into 10!
	//

	for inx := 0; inx < 25; inx++ {

		count := atomic.LoadUint64(counter) // get the concurrent counter value using atomic package
		if count < uint64(100) {
			atomic.AddUint64(counter, 1) // increment the counter by 1 concurrently
			log.Printf("problem1: slot=%03d count=%05d rand=%f", slot, inx, rand.Float32())
		} else {
			break // if the counter is >100, then break the loop and proceed to signal Done() for the goroutine
		}
	}
	wg.Done() // decrement the waitgroup counter by 1
}
