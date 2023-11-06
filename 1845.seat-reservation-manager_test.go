package main

import (
	"testing"
)

func TestSeatManger(t *testing.T) {
	temp := Constructor(5)
	sm := &temp

	expectReserve(t, sm, 1)
	expectReserve(t, sm, 2)

	sm.Unreserve(2)

	expectReserve(t, sm, 2)
	expectReserve(t, sm, 3)
	expectReserve(t, sm, 4)
	expectReserve(t, sm, 5)

	sm.Unreserve(2)
}

func expectReserve(t *testing.T, sm *SeatManager, want int) {
	if got := sm.Reserve(); got != want {
		t.Errorf("Reserve() = %v, want %v", got, want)
	}
}
