package concurrency

import "testing"

func TestChannel_Unbuffered(t *testing.T) {
	result := UnbufferedChannel()

	if result != "message" {
		t.Errorf("expected 'message', got %s", result)
	}
}

func TestChannel_Buffered(t *testing.T) {
	result := BufferedChannel()

	if result != "message" {
		t.Errorf("expected 'message', got %s", result)
	}
}

func TestChannel_Blocking(t *testing.T) {
	result := ChannelBlocking()

	if !result {
		t.Error("expected true")
	}
}

func TestChannel_BufferedFull(t *testing.T) {
	results := BufferedChannelFull()

	if len(results) != 2 {
		t.Errorf("expected 2 results, got %d", len(results))
	}

	if results[0] != 1 || results[1] != 2 {
		t.Errorf("expected [1, 2], got %v", results)
	}
}

func TestChannel_Direction(t *testing.T) {
	result := ChannelDirection()

	if result != 42 {
		t.Errorf("expected 42, got %d", result)
	}
}

func TestChannel_Range(t *testing.T) {
	count := 5
	results := ChannelRange(count)

	if len(results) != count {
		t.Errorf("expected %d results, got %d", count, len(results))
	}

	for i := 0; i < count; i++ {
		if results[i] != i {
			t.Errorf("expected results[%d] to be %d, got %d", i, i, results[i])
		}
	}
}

func TestChannel_Select(t *testing.T) {
	result := ChannelSelect()

	// Should receive from one of the channels
	if result != "from ch1" && result != "from ch2" {
		t.Errorf("unexpected result: %s", result)
	}
}

func TestChannel_NonBlocking(t *testing.T) {
	result := ChannelNonBlocking()

	// Should use default case since channel is empty
	if result != "no message" {
		t.Errorf("expected 'no message', got %s", result)
	}
}

func TestChannel_OrDone(t *testing.T) {
	// Test normal completion
	done := make(chan bool)
	result := ChannelOrDone(done)

	if result != "work done" {
		t.Errorf("expected 'work done', got %s", result)
	}

	// Test cancellation
	done2 := make(chan bool)
	close(done2) // Signal cancellation

	result2 := ChannelOrDone(done2)
	if result2 != "cancelled" {
		t.Errorf("expected 'cancelled', got %s", result2)
	}
}

func TestChannel_Nil(t *testing.T) {
	result := NilChannel()

	if result != "nil channel blocks" {
		t.Errorf("expected 'nil channel blocks', got %s", result)
	}
}

func TestChannel_Capacity(t *testing.T) {
	capacity, length := ChannelCapacity()

	if capacity != 5 {
		t.Errorf("expected capacity 5, got %d", capacity)
	}

	if length != 2 {
		t.Errorf("expected length 2, got %d", length)
	}
}

func TestChannel_Pipeline(t *testing.T) {
	numbers := []int{1, 2, 3, 4}
	results := ChannelPipeline(numbers)

	expected := []int{1, 4, 9, 16}

	if len(results) != len(expected) {
		t.Errorf("expected %d results, got %d", len(expected), len(results))
	}

	for i := 0; i < len(expected); i++ {
		if results[i] != expected[i] {
			t.Errorf("expected results[%d] to be %d, got %d", i, expected[i], results[i])
		}
	}
}

func TestChannel_SendReceive(t *testing.T) {
	ch := make(chan int, 1)

	// Send
	ch <- 42

	// Receive
	val := <-ch

	if val != 42 {
		t.Errorf("expected 42, got %d", val)
	}
}

func TestChannel_MultipleValues(t *testing.T) {
	ch := make(chan int, 3)

	// Send multiple values
	ch <- 1
	ch <- 2
	ch <- 3

	// Receive in order
	if <-ch != 1 {
		t.Error("expected first value to be 1")
	}
	if <-ch != 2 {
		t.Error("expected second value to be 2")
	}
	if <-ch != 3 {
		t.Error("expected third value to be 3")
	}
}

func TestChannel_ZeroValue(t *testing.T) {
	var ch chan int // nil channel

	if ch != nil {
		t.Error("expected nil channel")
	}

	// len and cap of nil channel are 0
	if len(ch) != 0 {
		t.Errorf("expected len 0, got %d", len(ch))
	}

	if cap(ch) != 0 {
		t.Errorf("expected cap 0, got %d", cap(ch))
	}
}
