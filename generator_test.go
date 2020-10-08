package generator

import (
	"strings"
	"testing"
)

func TestGenerateString(t *testing.T) {
	length := 100
	generatedString, err := String(length)

	if err != nil || len(generatedString) != length {
		t.Fatalf("Test failed")
	}
}

func TestGenerateStringWithCharset(t *testing.T) {
	length := 100
	charset := Lowercase

	generatedString, _ := StringWithCharset(length, charset)
	generatedStringArray := []rune(generatedString)

	for _, r := range generatedStringArray {
		if !strings.Contains(charset, string(r)) {
			t.Fatalf("Test failed: String %s does not contain %c", generatedString, r)
		}
	}
}

func TestGenerateStringListWithCharset(t *testing.T) {
	length := 5
	charset := Lowercase
	number := 10000000

	list, _ := StringListWithCharset(length, charset, number)
	exist := make(map[string]bool)

	for i, s := range list {
		exists := exist[s]
		if !exists {
			exist[s] = true
			list[i] = s
		} else {
			t.Fatalf("Test failed: Duplicate value %s", s)
		}
	}
}

func TestGenerateStringWithInvalidInputLength(t *testing.T) {
	_, err := String(0)

	if err == nil {
		t.Fatalf("Test failed: No error was thrown")
	}
}
