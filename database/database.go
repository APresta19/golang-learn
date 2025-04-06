package main

import (
	"fmt"
	"time"
	"sync"
)

//simulates a database call to learn about Goroutines
var mutex = sync.RWMutex{} //avoid data corruption
var wg = sync.WaitGroup{} //waits for all goroutines to finish before moving to main
var dbData = []string{"id1", "id2", "id3", "id4", "id5"} //database information
var results = []string{} //appended results

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\n The results are %v", results)
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("\nThe result from the database is: ", dbData[i])
	
	save(dbData[i])
	log()

	wg.Done()
}

func save(str string) {
	mutex.Lock() // full lock - check is performed to see if a lock is already performed by another goroutine
	results = append(results, str)
	mutex.Unlock()
}

func log() {
	//read data can still be read by multiple goroutines but not write 
	mutex.RLock() //will run if no full locks are active
	fmt.Printf("\nThe results are %v", results)
	mutex.RUnlock()
}
