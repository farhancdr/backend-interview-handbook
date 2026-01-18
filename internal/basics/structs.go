package basics

// Why interviewers ask this:
// Structs are fundamental to Go's type system. Understanding struct composition,
// embedding, and the difference from inheritance is crucial. Interviewers want
// to see if you understand Go's "composition over inheritance" philosophy.

// Common pitfalls:
// - Expecting inheritance-like behavior (Go doesn't have inheritance)
// - Not understanding promoted fields in embedding
// - Confusion about struct comparison and equality
// - Forgetting that structs are value types (copied on assignment)
// - Not knowing when to use pointers vs values

// Key takeaway:
// Structs are value types. Use embedding for composition, not inheritance.
// Promoted fields make embedded types' fields/methods accessible directly.

// Person represents a basic struct
type Person struct {
	Name string
	Age  int
}

// Address represents an address
type Address struct {
	Street string
	City   string
}

// Employee demonstrates struct embedding (composition)
type Employee struct {
	Person  // Embedded struct (promoted fields)
	Address // Another embedded struct
	Salary  int
}

// NewPerson creates a new Person
func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

// NewEmployee creates a new Employee
func NewEmployee(name string, age int, street, city string, salary int) Employee {
	return Employee{
		Person: Person{
			Name: name,
			Age:  age,
		},
		Address: Address{
			Street: street,
			City:   city,
		},
		Salary: salary,
	}
}

// Greet is a method on Person
func (p Person) Greet() string {
	return "Hello, I'm " + p.Name
}

// UpdateAge demonstrates value receiver (doesn't modify original)
func (p Person) UpdateAge(newAge int) {
	p.Age = newAge // Modifies copy, not original
}

// UpdateAgePointer demonstrates pointer receiver (modifies original)
func (p *Person) UpdateAgePointer(newAge int) {
	p.Age = newAge // Modifies original
}

// CompareStructs demonstrates struct comparison
func CompareStructs(p1, p2 Person) bool {
	return p1 == p2 // Structs can be compared if all fields are comparable
}

// CopyStruct demonstrates that structs are copied
func CopyStruct(p Person) Person {
	copy := p // Creates a copy
	copy.Age = 100
	return copy // Original p is unchanged
}

// AnonymousStruct demonstrates anonymous struct usage
func AnonymousStruct() struct {
	X int
	Y int
} {
	return struct {
		X int
		Y int
	}{X: 10, Y: 20}
}

// StructWithTags demonstrates struct tags (used for JSON, DB, etc.)
type StructWithTags struct {
	PublicField  string `json:"public_field"`
	privateField string // Not exported, won't be marshaled
	OmitEmpty    string `json:"omit_empty,omitempty"`
}

// ZeroValueStruct demonstrates zero value behavior
func ZeroValueStruct() Person {
	var p Person // Zero value: Name="", Age=0
	return p
}

// StructPointerVsValue demonstrates the difference
func StructPointerVsValue() (*Person, Person) {
	p1 := &Person{Name: "Alice", Age: 30} // Pointer
	p2 := Person{Name: "Bob", Age: 25}    // Value
	return p1, p2
}
