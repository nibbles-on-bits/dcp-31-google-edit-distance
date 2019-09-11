package main

import (
	"testing"
	"fmt"
)


func TestComputeEditDistance(t *testing.T) {


	got := ComputeEditDistance("kitten","sitten")
	if  got != 3 {
		t.Errorf(fmt.Sprintf("Expected 3, got %d\n", got))
	}
}

func TestFunc1(t *testing.T) {
	if !Func1() {
		t.Errorf("It failed homie")
	}

}
