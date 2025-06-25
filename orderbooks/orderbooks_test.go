package orderbooks

import (
	"testing"
)

func TestOrderbooks(t *testing.T) {
	// STEP 1: Add Order to orderbooks
	ob := NewOrderBook()
	ob.AddOrder("BUY", 100, 10)
	ob.AddOrder("SELL", 101, 20)
	res := ob.GetAll()
	if len(res) != 2 {
		t.Errorf("Expected 2 orders, got %d", len(res))
	}
}
