package main

import (
	"fmt"
	"sync"
	"time"
)

// var mr = sync.RWMutex{}
var m = sync.Mutex{}
var wg = sync.WaitGroup{} // wait group
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	// CONCURRENT, routines
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i) // since we are using goroutine we can run multiple db calls concurrently and not waiting for the previous one to finish
	}
	wg.Wait() // wait for all the goroutines to finish
	fmt.Printf("\nTotal execution time: %v\n", time.Since(t0))
	fmt.Printf("\nThe results are %v\n", results)
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
	m.Lock()
	results = append(results, dbData[i]) // collect the database excecution sequence
	m.Unlock()
	wg.Done() // signal that the goroutine is done
}

// multiple process writing to a same memory location can cause some unexpected results such as missings...The results are [id5 id3 id1 id4]
//we can use mutex for that

//func save(result string) {
//	mu.RLock() // the lock is set to read only not for writing
//	results = append(results, result)
//	mu.RUnlock()
//
//}
