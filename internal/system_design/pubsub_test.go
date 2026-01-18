package systemdesign

import (
	"testing"
	"time"
)

func TestPubSub(t *testing.T) {
	ps := NewPubSub()
	topic := "news"

	// 1. Subscribe
	ch1 := ps.Subscribe(topic)
	ch2 := ps.Subscribe(topic)

	// 2. Publish
	msg := "Breaking News"
	ps.Publish(topic, msg)

	// 3. Verify
	select {
	case received := <-ch1:
		if received != msg {
			t.Errorf("ch1 expected %s, got %s", msg, received)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("ch1 timed out")
	}

	select {
	case received := <-ch2:
		if received != msg {
			t.Errorf("ch2 expected %s, got %s", msg, received)
		}
	case <-time.After(100 * time.Millisecond):
		t.Error("ch2 timed out")
	}
}
