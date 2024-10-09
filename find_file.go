// Communicating sequential processes
// Don't communicate by sharing memory; share memory by communicating.

package main

import (
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

var query = "git"
var matches int
var workerCount = 0
var maxWorkerCount = 32

var mutex = &sync.Mutex{}

// Bidirection, with block
var searchRequest = make(chan string)
var workerDone = make(chan bool)
var foundResult = make(chan bool)

func main() {
	start_t := time.Now()
	workerCount = 1
	go search("/home/brianlee/", true)
	waitWorker()
	fmt.Println(matches, "matches")
	fmt.Println(time.Since(start_t))
}

func waitWorker() {
	for {
		select {
		case path := <-searchRequest:
			mutex.Lock()
			workerCount++
			mutex.Unlock()
			go search(path, true)
		case <-foundResult:
			matches++
		case <-workerDone:
			workerCount--
			// fmt.Printf("worker left is: %d\n", workerCount)
			if workerCount == 0 {
				return
			}
		}
	}
}

func search(path string, master bool) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		if master {
			workerDone <- true
		}
		return
	}
	for _, file := range files {
		name := file.Name()
		if name == query {
			foundResult <- true
		}
		if file.IsDir() {
			mutex.Lock()
			if workerCount < maxWorkerCount {
				mutex.Unlock()
				searchRequest <- path + name + "/"
			} else {
				mutex.Unlock()
				search(path+name+"/", false)
			}
		}
	}
	if master {
		workerDone <- true
	}
}
