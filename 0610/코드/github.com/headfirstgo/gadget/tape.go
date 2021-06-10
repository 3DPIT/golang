package gadget

import "fmt"

type NoiseMaker interface {
	MakeSound()
}

func Try(noise NoiseMaker) {
	noise.MakeSound()
	noise2, ok := noise.(Horn)
	if ok {
		noise2.MakeNoise()
		fmt.Println("나는 Horn이야")
	}
}

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

func (h Horn) MakeNoise() {
	fmt.Println("Noise")
}
