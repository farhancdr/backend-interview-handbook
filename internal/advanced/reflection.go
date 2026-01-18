package advanced

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// WalkStruct prints the field names and values of a struct.
// It handles nested structs and pointers.
// This demonstrates deep inspection using reflection.
func WalkStruct(v interface{}, depth int) []string {
	val := reflect.ValueOf(v)

	// Handle pointers
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return []string{fmt.Sprintf("%s<nil>", strings.Repeat("  ", depth))}
		}
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return []string{fmt.Sprintf("%s%v", strings.Repeat("  ", depth), val)}
	}

	var results []string
	t := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := t.Field(i)
		fieldVal := val.Field(i)

		// Check if it's exported (pkg path is empty for exported fields)
		if field.PkgPath != "" {
			continue
		}

		line := fmt.Sprintf("%sField: %s, Type: %s, Value: %v", strings.Repeat("  ", depth), field.Name, field.Type, fieldVal)
		results = append(results, line)

		// Recursive walk if it's a struct or pointer to struct
		if fieldVal.Kind() == reflect.Struct || (fieldVal.Kind() == reflect.Ptr && fieldVal.Elem().Kind() == reflect.Struct) {
			nested := WalkStruct(fieldVal.Interface(), depth+1)
			results = append(results, nested...)
		}
	}
	return results
}

// ValidateStruct checks if fields tagged with `required:"true"` are non-zero.
// This demonstrates reading struct tags.
func ValidateStruct(v interface{}) error {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return errors.New("input must be a struct")
	}

	t := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("required")

		if tag == "true" {
			fieldVal := val.Field(i)
			if isZero(fieldVal) {
				return fmt.Errorf("field '%s' is required", field.Name)
			}
		}
	}
	return nil
}

func isZero(v reflect.Value) bool {
	return v.IsValid() && v.IsZero()
}
