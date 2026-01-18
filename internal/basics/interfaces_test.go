package basics

import "testing"

func TestInterface_ImplicitImplementation(t *testing.T) {
	dog := Dog{Name: "Buddy"}

	// Dog implicitly implements Speaker
	var s Speaker = dog

	if s.Speak() != "Woof!" {
		t.Errorf("expected Woof!, got %s", s.Speak())
	}
}

func TestInterface_Polymorphism(t *testing.T) {
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}

	// Both can be passed to MakeSpeak
	dogSpeak := MakeSpeak(dog)
	catSpeak := MakeSpeak(cat)

	if dogSpeak != "Woof!" {
		t.Errorf("expected Woof!, got %s", dogSpeak)
	}

	if catSpeak != "Meow!" {
		t.Errorf("expected Meow!, got %s", catSpeak)
	}
}

func TestInterface_MultipleInterfaces(t *testing.T) {
	cat := Cat{Name: "Whiskers"}

	// Cat implements both Speaker and Mover
	var s Speaker = cat
	var m Mover = cat

	if s.Speak() != "Meow!" {
		t.Errorf("expected Meow!, got %s", s.Speak())
	}

	if m.Move() != "Walking on four legs" {
		t.Errorf("expected Walking on four legs, got %s", m.Move())
	}
}

func TestInterface_ComposedInterface(t *testing.T) {
	cat := Cat{Name: "Whiskers"}

	// Cat implements SpeakerMover (composed interface)
	var sm SpeakerMover = cat

	if sm.Speak() != "Meow!" {
		t.Errorf("expected Meow!, got %s", sm.Speak())
	}

	if sm.Move() != "Walking on four legs" {
		t.Errorf("expected Walking on four legs, got %s", sm.Move())
	}
}

func TestInterface_PointerReceiver(t *testing.T) {
	robot := &Robot{ID: 1}

	// Robot implements Speaker with pointer receiver
	var s Speaker = robot

	expected := "Robot 1 speaking"
	if s.Speak() != expected {
		t.Errorf("expected %s, got %s", expected, s.Speak())
	}
}

func TestInterface_EmptyInterface(t *testing.T) {
	// Empty interface accepts any type
	result1 := EmptyInterfaceExample("hello")
	result2 := EmptyInterfaceExample(42)
	result3 := EmptyInterfaceExample(true)

	if result1 != "Received: hello (type: string)" {
		t.Errorf("unexpected result: %s", result1)
	}

	if result2 != "Received: 42 (type: int)" {
		t.Errorf("unexpected result: %s", result2)
	}

	if result3 != "Received: true (type: bool)" {
		t.Errorf("unexpected result: %s", result3)
	}
}

func TestInterface_TypeAssertion(t *testing.T) {
	// Successful type assertion
	var i interface{} = "hello"
	s, ok := TypeAssertion(i)

	if !ok {
		t.Error("expected type assertion to succeed")
	}

	if s != "hello" {
		t.Errorf("expected hello, got %s", s)
	}

	// Failed type assertion
	i = 42
	s, ok = TypeAssertion(i)

	if ok {
		t.Error("expected type assertion to fail")
	}

	if s != "" {
		t.Errorf("expected empty string, got %s", s)
	}
}

func TestInterface_TypeSwitch(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{"hello", "String: hello"},
		{42, "Int: 42"},
		{true, "Bool: true"},
		{3.14, "Unknown type: float64"},
	}

	for _, tt := range tests {
		result := TypeSwitch(tt.input)
		if result != tt.expected {
			t.Errorf("expected %s, got %s", tt.expected, result)
		}
	}
}

func TestInterface_NilInterface(t *testing.T) {
	s := NilInterface()

	// Nil interface
	if s != nil {
		t.Error("expected nil interface")
	}
}

func TestInterface_NilValueGotcha(t *testing.T) {
	s := InterfaceWithNilValue()

	// This is the gotcha: interface is not nil, but holds nil value
	if s == nil {
		t.Error("expected interface to not be nil (holds nil value)")
	}

	// This would panic if uncommented:
	// s.Speak() // panic: nil pointer dereference
}

func TestInterface_CheckNil(t *testing.T) {
	// Nil interface
	var s1 Speaker
	if !CheckNil(s1) {
		t.Error("expected nil interface to be nil")
	}

	// Interface with nil value (gotcha!)
	var d *Dog
	var s2 Speaker = d
	if CheckNil(s2) {
		t.Error("expected interface with nil value to not be nil")
	}
}

func TestInterface_Comparison(t *testing.T) {
	dog1 := Dog{Name: "Buddy"}
	dog2 := Dog{Name: "Buddy"}
	dog3 := Dog{Name: "Max"}

	var s1 Speaker = dog1
	var s2 Speaker = dog2
	var s3 Speaker = dog3

	// Same values
	if !InterfaceComparison(s1, s2) {
		t.Error("expected s1 and s2 to be equal")
	}

	// Different values
	if InterfaceComparison(s1, s3) {
		t.Error("expected s1 and s3 to be different")
	}
}

func TestInterface_AcceptAnything(t *testing.T) {
	results := AcceptAnything("hello", 42, true, 3.14)

	expected := []string{"hello", "42", "true", "3.14"}

	if len(results) != len(expected) {
		t.Errorf("expected %d results, got %d", len(expected), len(results))
	}

	for i, exp := range expected {
		if results[i] != exp {
			t.Errorf("expected %s, got %s", exp, results[i])
		}
	}
}

func TestInterface_ValueVsPointerReceiver(t *testing.T) {
	// Value receiver: both value and pointer satisfy interface
	dog := Dog{Name: "Buddy"}
	dogPtr := &Dog{Name: "Max"}

	var s1 Speaker = dog
	var s2 Speaker = dogPtr

	if s1.Speak() != "Woof!" {
		t.Errorf("expected Woof!, got %s", s1.Speak())
	}

	if s2.Speak() != "Woof!" {
		t.Errorf("expected Woof!, got %s", s2.Speak())
	}

	// Pointer receiver: only pointer satisfies interface
	robotPtr := &Robot{ID: 1}
	var s3 Speaker = robotPtr

	if s3.Speak() != "Robot 1 speaking" {
		t.Errorf("expected Robot 1 speaking, got %s", s3.Speak())
	}

	// This won't compile (value doesn't satisfy interface with pointer receiver):
	// robot := Robot{ID: 1}
	// var s4 Speaker = robot
}
