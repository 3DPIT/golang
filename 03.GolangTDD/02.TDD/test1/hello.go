package hello

type Speaker interface {
	Speak(string)
}
func Hello(s Speaker, friend string) {
	s.Speak("Hello" + friend + "!")
}
