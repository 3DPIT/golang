package main

import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep Boop")
}

func (r Robot) Walk() {
	fmt.Println("Powering legs")
}

type NoiseMaker interface { // MakeSound 메서드를 가진 모든 타입 표현
	MakeSound()
}

func play(n NoiseMaker) {
	n.MakeSound()
}

func main() {

	var noiseMaker NoiseMaker = Robot("iron man")
	noiseMaker.MakeSound()
	var robot Robot = Robot{}
}
