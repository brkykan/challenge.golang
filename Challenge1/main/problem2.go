package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func problem2() {

	var wg sync.WaitGroup
	tick := time.Tick(1 * time.Second) // create a channel of type time.Time that will send for every 1 second
	log.Printf("problem2: started --------------------------------------------")

	//
	// Todo:
	//
	// Throttle all go subroutines in a way,
	// that every one second one random number
	// is printed.
	//

	for inx := 0; inx < 10; inx++ {
		wg.Add(1)
		go printRandom2(inx, &wg, tick)

	}
	wg.Wait() // block the execution until waitgroup counter is again 0
	//
	// Todo:
	//
	// Remove this quick and dirty sleep
	// against a synchronized wait until all
	// go routines are finished.
	//
	// Same as problem1...
	//

	//time.Sleep(5 * time.Second)

	log.Printf("problem2: finished -------------------------------------------")
}

func printRandom2(slot int, wg *sync.WaitGroup, tick <-chan time.Time) {

	for inx := 0; inx < 10; inx++ {
		<-tick // block the execution until it receives from Time channel `tick`
		log.Printf("problem2: slot=%03d count=%05d rand=%f", slot, inx, rand.Float32())
	}
	wg.Done() // decrement the waitgroup counter by 1
}
