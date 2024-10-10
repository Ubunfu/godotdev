package main

import (
	"testing"
	"unicode/utf8"
)

// test name syntax is important TestXxx
func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}

	for _, tc := range testcases {
		rev, err := Reverse(tc.in)
		if err != nil {
			t.Skip()
		}
		if rev != tc.want {
			t.Errorf("Reversed output is: %q, but expected %q", rev, tc.want)
		}
	}
}

// test name syntax is important FuzzXxx
func FuzzReverse(f *testing.F) {
	testCases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testCases {
		f.Add(tc) // seeds the fuzz test? "seed corpus"?
	}
	f.Fuzz(func(t *testing.T, original string) {
		rev, err1 := Reverse(original)
		if err1 != nil {
			return
			// or t.Skip()
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
			// or t.Skip()
		}
		if original != doubleRev {
			t.Errorf("Expected original: %q, to match double-reverse: %q\n", original, doubleRev)
		}
		if utf8.ValidString(original) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string: %q\n", rev)
		}
	})
}
