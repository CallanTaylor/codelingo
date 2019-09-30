package main

import (
	"fmt"
)

type thing struct {
	name string
}

func readVulnerableField(t *thing) {
	fmt.Println(t.name)
}

func getSafeField(t thing) {
	fmt.Println(t.name)
}

func writeVulnerableField(t *thing) {
	t.name = "car"
}

// Here we perform some reads and writes to the data field of a pointer to
// a 'thing'. The Tenet should only catch attempts to write to the fields of
// a pointer that happen on seperate threads to the calling context.
func readAndWriteToPointer() {

	t := &thing{"Pointer"}

	readVulnerableField(t)

	go readVulnerableField(t)

	go writeVulnerableField(t) // Issue

	go func(t *thing) {
		fmt.Println(t.name)
	}(t)

	go func(t *thing) { // Issue
		t.name = "car"
	}(t)

}

// Here we perfor read and write opertations to a field of a 'thing', however
// here we are passing the 'thing' itself to the read and write methods which
// is safe to do concurrently to other operations on the calling context.
func readAndWriteToObject() {

	t2 := thing{"Not-pointer"}

	getSafeField(t2)

	go getSafeField(t2)

	go func(t thing) {
		fmt.Println(t.name)
	}(t2)

	go func(t thing) {
		t.name = "bus"
	}(t2)
}

func main() {
	readAndWriteToPointer()
	readAndWriteToObject()
}
