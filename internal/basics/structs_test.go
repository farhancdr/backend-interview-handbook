package basics

import "testing"

func TestStruct_Creation(t *testing.T) {
	p := NewPerson("Alice", 30)

	if p.Name != "Alice" {
		t.Errorf("expected Name to be Alice, got %s", p.Name)
	}

	if p.Age != 30 {
		t.Errorf("expected Age to be 30, got %d", p.Age)
	}
}

func TestStruct_Embedding(t *testing.T) {
	e := NewEmployee("Bob", 25, "123 Main St", "NYC", 50000)

	// Promoted fields from Person
	if e.Name != "Bob" {
		t.Errorf("expected Name to be Bob, got %s", e.Name)
	}

	if e.Age != 25 {
		t.Errorf("expected Age to be 25, got %d", e.Age)
	}

	// Promoted fields from Address
	if e.Street != "123 Main St" {
		t.Errorf("expected Street to be 123 Main St, got %s", e.Street)
	}

	if e.City != "NYC" {
		t.Errorf("expected City to be NYC, got %s", e.City)
	}

	// Employee's own field
	if e.Salary != 50000 {
		t.Errorf("expected Salary to be 50000, got %d", e.Salary)
	}
}

func TestStruct_PromotedMethods(t *testing.T) {
	e := NewEmployee("Charlie", 35, "456 Oak Ave", "LA", 60000)

	// Greet method is promoted from Person
	greeting := e.Greet()
	expected := "Hello, I'm Charlie"

	if greeting != expected {
		t.Errorf("expected %s, got %s", expected, greeting)
	}
}

func TestStruct_ValueReceiver(t *testing.T) {
	p := NewPerson("Alice", 30)

	// Value receiver doesn't modify original
	p.UpdateAge(40)

	if p.Age != 30 {
		t.Errorf("expected Age to still be 30, got %d", p.Age)
	}
}

func TestStruct_PointerReceiver(t *testing.T) {
	p := NewPerson("Alice", 30)

	// Pointer receiver modifies original
	p.UpdateAgePointer(40)

	if p.Age != 40 {
		t.Errorf("expected Age to be 40, got %d", p.Age)
	}
}

func TestStruct_Comparison(t *testing.T) {
	p1 := Person{Name: "Alice", Age: 30}
	p2 := Person{Name: "Alice", Age: 30}
	p3 := Person{Name: "Bob", Age: 25}

	// Identical structs are equal
	if !CompareStructs(p1, p2) {
		t.Error("expected p1 and p2 to be equal")
	}

	// Different structs are not equal
	if CompareStructs(p1, p3) {
		t.Error("expected p1 and p3 to be different")
	}
}

func TestStruct_Copy(t *testing.T) {
	original := Person{Name: "Alice", Age: 30}
	copy := CopyStruct(original)

	// Copy should be modified
	if copy.Age != 100 {
		t.Errorf("expected copy.Age to be 100, got %d", copy.Age)
	}

	// Original should be unchanged
	if original.Age != 30 {
		t.Errorf("expected original.Age to be 30, got %d", original.Age)
	}
}

func TestStruct_AnonymousStruct(t *testing.T) {
	point := AnonymousStruct()

	if point.X != 10 {
		t.Errorf("expected X to be 10, got %d", point.X)
	}

	if point.Y != 20 {
		t.Errorf("expected Y to be 20, got %d", point.Y)
	}
}

func TestStruct_ZeroValue(t *testing.T) {
	p := ZeroValueStruct()

	// Zero value for string is ""
	if p.Name != "" {
		t.Errorf("expected Name to be empty string, got %s", p.Name)
	}

	// Zero value for int is 0
	if p.Age != 0 {
		t.Errorf("expected Age to be 0, got %d", p.Age)
	}
}

func TestStruct_PointerVsValue(t *testing.T) {
	pPtr, pVal := StructPointerVsValue()

	// Pointer
	if pPtr.Name != "Alice" {
		t.Errorf("expected pointer Name to be Alice, got %s", pPtr.Name)
	}

	// Value
	if pVal.Name != "Bob" {
		t.Errorf("expected value Name to be Bob, got %s", pVal.Name)
	}

	// Modify via pointer
	pPtr.Age = 100
	if pPtr.Age != 100 {
		t.Errorf("expected pointer Age to be 100, got %d", pPtr.Age)
	}
}

func TestStruct_FieldAccess(t *testing.T) {
	e := Employee{
		Person: Person{
			Name: "Alice",
			Age:  30,
		},
		Address: Address{
			Street: "123 Main St",
			City:   "NYC",
		},
		Salary: 50000,
	}

	// Access via promoted fields
	if e.Name != "Alice" {
		t.Errorf("expected Name to be Alice, got %s", e.Name)
	}

	// Access via embedded struct
	if e.Person.Name != "Alice" {
		t.Errorf("expected Person.Name to be Alice, got %s", e.Person.Name)
	}

	// Both ways access the same field
	e.Name = "Bob"
	if e.Person.Name != "Bob" {
		t.Errorf("expected Person.Name to be Bob, got %s", e.Person.Name)
	}
}

func TestStruct_EmptyStruct(t *testing.T) {
	// Empty struct has size 0
	type Empty struct{}

	e := Empty{}
	_ = e

	// Empty structs are useful as signals in channels
	done := make(chan struct{})

	go func() {
		// Do work
		done <- struct{}{}
	}()

	<-done // Wait for signal
}
