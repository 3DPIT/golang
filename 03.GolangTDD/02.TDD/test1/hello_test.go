package hello

import "testing"

type fakeSpeaker struct {
	CallCount int
}

func (f *fakeSpeaker) Speak(string) {
	f.CallCount++
}
func TestHelloWorld(t *testing.T) {
	t.Run("calls Specker", func(t *testing.T) {
		f := &fakeSpeaker{}
		Hello(f, "GoCon")

		actual := f.CallCount
		expected := 1
		if actual != expected {
			t.Errorf("Got: %v, Expected: %v", actual, expected)
		}
	})
}
