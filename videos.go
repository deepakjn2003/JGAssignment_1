package main

import (
	"fmt"
	"sync"
	"time"
)

type viewMap struct {
	viewCountMap map[string]int
	mutex        sync.Mutex
}

// package level variable
var myLibrary viewMap

// method to print map in "/"
func (v *viewMap) String() string {

	var resultantString string
	for k, v := range v.viewCountMap {
		resultantString += fmt.Sprintf("-> Video id : %v, count : %v\n", k, v)
	}

	return resultantString
}

// fnnction to populate map
func createMap() {
	var lib = &myLibrary
	lib.viewCountMap = make(map[string]int)
	lib.viewCountMap["vid01"] = 1
	lib.viewCountMap["vid02"] = 5
	lib.viewCountMap["vid03"] = 10
}

// method to check view count for a given video id
func (v *viewMap) getViewCount(videoId string) int {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	count := v.viewCountMap[videoId]
	return count
}

// method to increment viewcount for a given video id by 1
func (v *viewMap) incrementViewCount(videoId string) {
	v.mutex.Lock()
	v.viewCountMap[videoId] += 1
	v.mutex.Unlock()

	// introducing some delay to demonstrate goroutine
	time.Sleep(5 * time.Second)
}
