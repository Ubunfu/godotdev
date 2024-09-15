package greetings

import (
	"regexp"
	"testing"
)

// Calls Hello with a name, checking the return value
func TestHelloName(t *testing.T) {
	name := "Ryan"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("%v") = %q, %v, want match for %#q, nil`, name, msg, err, want)
	}
}

// Calls Hello with an empty string and checks for an error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("" = %q, %v, want "", error`, msg, err)
	}
}
