package greetings

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Luis")
	expected := "Hi, Luis. Welcome!"

	if result != expected {
		t.Fatalf("expected %q, got %q", expected, result)
	}
}
