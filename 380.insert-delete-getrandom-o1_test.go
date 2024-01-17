package main

import "testing"

func TestRandomizedSet(t *testing.T) {
	randomizedSet := Constructor()
	res := randomizedSet.Insert(1)
	if res != true {
		t.Errorf("Expected true, got %v", res)
	}
	res = randomizedSet.Remove(2)
	if res != false {
		t.Errorf("Expected false, got %v", res)
	}
	res = randomizedSet.Insert(2)
	if res != true {
		t.Errorf("Expected true, got %v", res)
	}
	val := randomizedSet.GetRandom()
	if val != 1 && val != 2 {
		t.Errorf("Expected 1 or 2, got %v", val)
	}
	res = randomizedSet.Remove(1)
	if res != true {
		t.Errorf("Expected true, got %v", res)
	}
	res = randomizedSet.Insert(2)
	if res != false {
		t.Errorf("Expected false, got %v", res)
	}
	val = randomizedSet.GetRandom()
	if val != 2 {
		t.Errorf("Expected 2, got %v", val)
	}
}
