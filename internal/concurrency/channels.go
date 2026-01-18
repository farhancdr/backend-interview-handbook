package concurrency

// Why interviewers ask this:
// Channels are Go's primary mechanism for communication between goroutines.
// Understanding buffered vs unbuffered channels, blocking behavior, and proper
// usage is absolutely critical for Go interviews. This is asked in 90%+ of interviews.

// Common pitfalls:
// - Sending to unbuffered channel without receiver causes deadlock
// - Not understanding that unbuffered channels are synchronous
// - Forgetting that buffered channels block when full
// - Sending on closed channel causes panic
// - Not knowing when to use buffered vs unbuffered

// Key takeaway:
// Unbuffered channels (make(chan T)) are synchronous - sender blocks until receiver ready.
// Buffered channels (make(chan T, n)) are asynchronous up to capacity n.
// Channels enable safe communication between goroutines without explicit locks.

// UnbufferedChannel demonstrates unbuffered channel behavior
func UnbufferedChannel() string {
	ch := make(chan string) // Unbuffered

	go func() {
		ch <- "message" // Blocks until receiver ready
	}()

	return <-ch // Receives the message
}

// BufferedChannel demonstrates buffered channel behavior
func BufferedChannel() string {
	ch := make(chan string, 1) // Buffered with capacity 1

	ch <- "message" // Doesn't block, buffer has space

	return <-ch // Receives the message
}

// ChannelBlocking demonstrates blocking behavior
func ChannelBlocking() bool {
	ch := make(chan int)

	go func() {
		<-ch // Receiver ready
	}()

	ch <- 42 // Blocks until receiver ready
	return true
}

// BufferedChannelFull demonstrates buffered channel filling up
func BufferedChannelFull() []int {
	ch := make(chan int, 2) // Capacity 2

	ch <- 1 // OK, buffer has space
	ch <- 2 // OK, buffer has space
	// ch <- 3 would block here (buffer full)

	results := []int{<-ch, <-ch}
	return results
}

// ChannelDirection demonstrates send-only and receive-only channels
func ChannelDirection() int {
	ch := make(chan int, 1)

	// Send-only channel
	send := func(c chan<- int) {
		c <- 42
	}

	// Receive-only channel
	receive := func(c <-chan int) int {
		return <-c
	}

	send(ch)
	return receive(ch)
}

// ChannelRange demonstrates ranging over a channel
func ChannelRange(count int) []int {
	ch := make(chan int, count)

	// Send values
	go func() {
		for i := 0; i < count; i++ {
			ch <- i
		}
		close(ch) // Must close to end range
	}()

	// Receive via range
	var results []int
	for val := range ch {
		results = append(results, val)
	}

	return results
}

// ChannelSelect demonstrates select statement with channels
func ChannelSelect() string {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	ch1 <- "from ch1"
	ch2 <- "from ch2"

	select {
	case msg := <-ch1:
		return msg
	case msg := <-ch2:
		return msg
	}
}

// ChannelNonBlocking demonstrates non-blocking receive with select
func ChannelNonBlocking() string {
	ch := make(chan string)

	select {
	case msg := <-ch:
		return msg
	default:
		return "no message"
	}
}

// ChannelOrDone demonstrates channel coordination pattern
func ChannelOrDone(done <-chan bool) string {
	ch := make(chan string)

	go func() {
		ch <- "work done"
	}()

	select {
	case msg := <-ch:
		return msg
	case <-done:
		return "cancelled"
	}
}

// NilChannel demonstrates nil channel behavior
func NilChannel() string {
	var ch chan string // nil channel

	// Reading from nil channel blocks forever
	// Sending to nil channel blocks forever
	// This is useful in select statements to disable a case

	select {
	case <-ch:
		return "received" // Never happens
	default:
		return "nil channel blocks"
	}
}

// ChannelCapacity demonstrates checking channel capacity and length
func ChannelCapacity() (int, int) {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2

	return cap(ch), len(ch) // capacity=5, length=2
}

// ChannelPipeline demonstrates channel pipeline pattern
func ChannelPipeline(numbers []int) []int {
	// Stage 1: Generate numbers
	gen := func(nums []int) <-chan int {
		out := make(chan int)
		go func() {
			for _, n := range nums {
				out <- n
			}
			close(out)
		}()
		return out
	}

	// Stage 2: Square numbers
	square := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				out <- n * n
			}
			close(out)
		}()
		return out
	}

	// Build pipeline
	c1 := gen(numbers)
	c2 := square(c1)

	// Collect results
	var results []int
	for result := range c2 {
		results = append(results, result)
	}

	return results
}
