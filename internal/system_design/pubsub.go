package systemdesign

import (
	"sync"
)

// Why interviewers ask this:
// Pub-Sub is the backbone of event-driven architectures. Implementing a simple version
// demonstrates understanding of channels, goroutines, and safe slice manipulation.
// It also touches on decoupling producers from consumers.

// Common pitfalls:
// - Blocking the publisher if a subscriber is slow (use buffered channels or goroutines)
// - Writing to a closed channel
// - Not providing a way to unsubscribe (memory leak)

// Key takeaway:
// Use a map of `topic -> []chan string` to store subscribers.
// Protect the map with a Mutex.
// When publishing, iterate through the list and send the message (non-blocking preferred).

type PubSub struct {
	mu     sync.RWMutex
	topics map[string][]chan string
}

func NewPubSub() *PubSub {
	return &PubSub{
		topics: make(map[string][]chan string),
	}
}

// Subscribe returns a channel that receives messages for a topic
func (ps *PubSub) Subscribe(topic string) <-chan string {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch := make(chan string, 10) // Buffered to prevent immediate blocking
	ps.topics[topic] = append(ps.topics[topic], ch)

	return ch
}

// Publish sends a message to all subscribers of a topic
func (ps *PubSub) Publish(topic string, msg string) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	subscribers, ok := ps.topics[topic]
	if !ok {
		return
	}

	for _, ch := range subscribers {
		// Non-blocking send or spin up goroutine
		// For simplicity/safety in this example, we use a non-blocking select
		// to avoid freezing the publisher if a consumer is dead.
		select {
		case ch <- msg:
		default:
			// Subscriber slow/full, dropped message (implementation choice)
		}
	}
}

// CloseTopic closes all subscriber channels for a topic (cleanup)
func (ps *PubSub) CloseTopic(topic string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if subscribers, ok := ps.topics[topic]; ok {
		for _, ch := range subscribers {
			close(ch)
		}
		delete(ps.topics, topic)
	}
}
