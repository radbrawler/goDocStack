package main

import (
	"fmt"
	"sync"
)

type MapIntString map[int]string

type RWMap struct {
	sync.RWMutex
	rwm MapIntString
}

func main() {
	var m MapIntString

	fmt.Println("\n==== Declare And Initialize ====")
	m = declareAndInitialize()

	fmt.Println("\n==== Counting map elements ====")
	elementCount(m)

	fmt.Println("\n==== Check for a element in a map ====")
	key := 1 // should from declareAndInitialize()
	findElement(m, key)

	fmt.Println("\n==== Zero value of a map ====")
	zeroValue()

	fmt.Println("\n==== Iterating elements over a map ====")
	iterateElements(m)

	fmt.Println("\n==== Concurrent Access of maps ====")
	concurrentAccess(m)

	fmt.Println("\n==== Delete Map Element ====")
	deleteMapElement(m, 4)

	fmt.Println("\n==== Getting the Array of Keys and Values from a Map ====")
	getKeyArray(m)
}

func declareAndInitialize() MapIntString {
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

func elementCount(m MapIntString) {
	fmt.Println(" length of map:", len(m))
}

func findElement(m MapIntString, key int) {
	fmt.Println(" Finding Key ...")
	if val, err := m[key]; !err {
		fmt.Println(" key not found")
	} else {
		fmt.Println(" key value pair is ", key, ":", val)
	}
}

func zeroValue() {
	var zeroMap map[int]string
	fmt.Print(" Zero map => ")
	elementCount(zeroMap)
	fmt.Println(" Value of Expression (zeroMap == nil): ", zeroMap == nil)

	var nonZero MapIntString
	nonZero = make(MapIntString)
	fmt.Print(" Non Zero map => ")
	elementCount(nonZero)

	// Writing to non zero map
	nonZero[5] = "sand king"
	fmt.Println(nonZero[5])
}

func iterateElements(m MapIntString) {
	fmt.Print(" Printing All Elements")
	for key, val := range m {
		fmt.Println(" key: ", key, "value: ", val)
	}
}

func concurrentAccess(m MapIntString) {
	// Init
	fmt.Println(" Initializing new map struct...")
	counter := RWMap{rwm: make(MapIntString)}
	counter.rwm = m

	// Get a read lock
	fmt.Println(" Acquiring read lock...")
	counter.RLock()
	fmt.Println(" Accessing variable... \t value: ", counter.rwm[1])
	fmt.Println(" Releasing read lock...")
	counter.RUnlock()

	//Get write lock
	fmt.Println(" Acquiring R/W lock...")
	counter.Lock()
	fmt.Print(" Changing map value... \t")
	counter.rwm[1] = "a new hero"
	fmt.Print("New Value: ", counter.rwm[1], "\n")
	counter.Unlock()
}

func deleteMapElement(m MapIntString, key int) {
	fmt.Println(" Deleting element {", key, ":", m[key], "} ... ")
	delete(m, key)
	fmt.Println(" Element deleted ")
}

func getKeyArray(m MapIntString) {
	keys := make([]int, len(m))
	i := 0
	for key := range m {
		keys[i] = key
		i++
	}
	fmt.Print("\n Keys are: ", keys)

	vals := make([]string, len(m))
	i = 0
	for _, val := range m {
		vals[i] = val
		i++
	}
	fmt.Print("\n Values are: ", vals)
}

func mapAsSet(m MapIntString) {

}
