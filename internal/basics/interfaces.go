package basics

import "fmt"

// Why interviewers ask this:
// Interfaces are central to Go's type system and enable polymorphism. Understanding
// implicit implementation, empty interfaces, and interface satisfaction is crucial.
// This is one of the most commonly asked topics in Go interviews.

// Common pitfalls:
// - Not understanding implicit interface implementation
// - Confusion about nil interfaces vs interfaces holding nil values
// - Thinking you need to declare "implements" (you don't)
// - Not knowing that empty interface (interface{}) accepts any type
// - Forgetting that interfaces are satisfied by value OR pointer receivers

// Key takeaway:
// Interfaces are implemented implicitly. Any type that has the required methods
// satisfies the interface. Empty interface (any in Go 1.18+) accepts any value.

// Speaker interface defines a contract
type Speaker interface {
	Speak() string
}

// Mover interface defines another contract
type Mover interface {
	Move() string
}

// SpeakerMover combines multiple interfaces
type SpeakerMover interface {
	Speaker
	Mover
}

// Dog implements Speaker (implicitly)
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

// Cat implements both Speaker and Mover
type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

func (c Cat) Move() string {
	return "Walking on four legs"
}

// Robot implements Speaker with pointer receiver
type Robot struct {
	ID int
}

func (r *Robot) Speak() string {
	return fmt.Sprintf("Robot %d speaking", r.ID)
}

// MakeSpeak demonstrates polymorphism via interfaces
func MakeSpeak(s Speaker) string {
	return s.Speak()
}

// EmptyInterfaceExample demonstrates interface{} (any)
func EmptyInterfaceExample(v interface{}) string {
	return fmt.Sprintf("Received: %v (type: %T)", v, v)
}

// TypeAssertion demonstrates type assertion
func TypeAssertion(i interface{}) (string, bool) {
	s, ok := i.(string) // Type assertion with safety check
	return s, ok
}

// TypeSwitch demonstrates type switch
func TypeSwitch(i interface{}) string {
	switch v := i.(type) {
	case string:
		return "String: " + v
	case int:
		return fmt.Sprintf("Int: %d", v)
	case bool:
		return fmt.Sprintf("Bool: %t", v)
	default:
		return fmt.Sprintf("Unknown type: %T", v)
	}
}

// NilInterface demonstrates nil interface behavior
func NilInterface() Speaker {
	var s Speaker // nil interface
	return s
}

// InterfaceWithNilValue demonstrates interface holding nil value
func InterfaceWithNilValue() Speaker {
	var d *Dog // nil pointer
	return d   // interface is not nil, but holds nil value
}

// CheckNil demonstrates the nil interface gotcha
func CheckNil(s Speaker) bool {
	return s == nil
}

// InterfaceComparison demonstrates interface comparison
func InterfaceComparison(s1, s2 Speaker) bool {
	return s1 == s2
}

// AcceptAnything demonstrates empty interface usage
func AcceptAnything(values ...interface{}) []string {
	results := []string{}
	for _, v := range values {
		results = append(results, fmt.Sprintf("%v", v))
	}
	return results
}
