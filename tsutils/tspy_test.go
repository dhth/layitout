package tsutils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetPyData(t *testing.T) {
	cases := []struct {
		name         string
		fileContents []byte
		expected     []string
		err          error
	}{
		// SUCCESSES
		{
			name: "a function with no parameters or return type hint",
			fileContents: []byte(`
def say_something():
    print("hola")
`),
			expected: []string{"def say_something()"},
		},
		{
			name: "a function with parameters but no return type hint",
			fileContents: []byte(`
def say_something(x):
    print("hola " + x)
`),
			expected: []string{"def say_something(x)"},
		},
		{
			name: "a function with parameters with types but no return type hint",
			fileContents: []byte(`
def say_something(x: str, y: str):
    print(f"hola, {x} and {y}")
`),
			expected: []string{"def say_something(x: str, y: str)"},
		},
		{
			name: "a function with return type hint but no parameters",
			fileContents: []byte(`
def say_something() -> str:
    return "hola"
`),
			expected: []string{"def say_something() -> str"},
		},
		{
			name: "a function with parameters and return type hints",
			fileContents: []byte(`
def say_something(x: str, y: str) -> str:
    return f"hola, {x} and {y}"
`),
			expected: []string{"def say_something(x: str, y: str) -> str"},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getPyData(tt.fileContents)
			fmt.Println(got)
			fmt.Println(tt.expected)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got: %#v, expected: %#v", got, tt.expected)
			}
			if err != tt.err {
				t.Errorf("error mismatch; got: %v, expected: %v", err, tt.err)
			}
		})
	}

}
