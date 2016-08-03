package main

import (
	"fmt"
	"sync"
)

type GoMap map[int]string

type RWMap struct {
	sync.RWMutex
	rwm GoMap
}

func main() {
	var m GoMap

	// Declare And Initialize
	m = declareAndInitialize()

	// Counting map elements
	elementCount(m)

	// Check for a element in a map
	key := 1 // should from declareAndInitialize()
	findElement(m, key)

	// Zero value of a map
	zeroValue()

	// Iterating elements over a map
	iterateElements(m)
}

func declareAndInitialize() GoMap {
	// var m1, m2 can't be used directly, they'll cause panic
	// Keys are ints, values are ints
	// var m1 map[int]int

	// Keys are strings, values are int
	// var m2 map[string]int

	// Using the function overiding we can set initial capacity of maps
	// m := make(map[string]int)
	m := map[int]string{1: "invoker", 2: "beastmaster", 3: "mirana", 4: "timbersaw"}
	return m
}

func elementCount(m GoMap) {
	fmt.Println("length of map:", len(m))
}

func findElement(m GoMap, key int) {
	fmt.Println("Finding Key ...")
	if val, err := m[key]; !err {
		fmt.Println("key not found")
	} else {
		fmt.Println("key value pair is ", key, ":", val)
	}
}

func zeroValue() {
	var zeroMap map[int]string
	fmt.Print("Zero map => ")
	elementCount(zeroMap)
	fmt.Println("Value of Expression (zeroMap == nil): ", zeroMap == nil)

	var nonZero GoMap
	nonZero = make(GoMap)
	fmt.Print("Non Zero map => ")
	elementCount(nonZero)

	// Writing to non zero map
	nonZero[5] = "sand king"
	fmt.Println(nonZero[5])
}

func iterateElements(m GoMap) {
	fmt.Print("Printing All Elements")
	for key, val := range m {
		fmt.Println("key: ", key, "value: ", val)
	}
}

func concurrentAccess(m GoMap) {
	// Init
	fmt.Println(" Initializing new map struct...")
	counter := RWMap{rwm: make(GoMap)}
	counter.rwm = m

	// Get a read lock
	fmt.Println(" Acquiring read lock...")
	counter.RLock()
	fmt.Println(" Accessing variable... \n value: ", counter.rwm[1])
	fmt.Println(" Releasing read lock...")
	counter.RUnlock()

	//Get write lock
	fmt.Println(" Acquiring R/W lock...")
	counter.Lock()
	fmt.Println(" Changing map value...")
	counter.rwm[1] = "a new hero"
	counter.Unlock()

}
