package hello

import "testing"

func TestHelloWorld(t *testing.T) {
	got := Hello("GoCon")
	expected := "Hello GoCon!"
	if got != expected {
		t.Errorf("Got: %s, Expected: %s", got, expected)
	}
}
