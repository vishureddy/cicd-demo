package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	// emptyResult := hello("")

	// //fmt.println(emptyResult)

	// if emptyResult != "Hello Dude" {
	// 	t.Errorf("hello(\"\")failed , expected %v, got %v", "Hello Dude", emptyResult)
	// }

	result := hello("Vishu")

	if result != "Hello Vishu" {
		t.Errorf("hello(\"\")failed , expected %v, got %v", "Hello Dude", result)
	}
}
