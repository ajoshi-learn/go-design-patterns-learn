package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()

	if counter1 == nil {
		t.Error("Expected pointer to singleton, not nil")
	}

	expectedCounter := counter1

	currentCount := counter1.AddOne()

	if currentCount != 1 {
		t.Error("After first calling it must be 1")
	}

	counter2 := GetInstance()

	if counter2 != expectedCounter {
		t.Error("Expected same instance")
	}

	currentCount = counter2.AddOne()

	if currentCount != 2 {
		t.Error("After second calling it must be 2")
	}
}
