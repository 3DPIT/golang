//인터페이스 없이 구현
package main

import "fmt"

type Bread struct {
	val string
}

type StrawberryJam struct {
}
type SpoonOfStrawberryJam struct {
}

func (s *SpoonOfStrawberryJam) String() string {
	return " + strawberry"
}
func (j *StrawberryJam) GetOneSpoon() *SpoonOfStrawberryJam {
	return &SpoonOfStrawberryJam{}
}
func (b *Bread) PutJam(jam *StrawberryJam) {
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *Bread) String() string {
	return "bread" + b.val
}
func main() {
	bread := &Bread{}
	jam := &StrawberryJam{}
	bread.PutJam(jam)

	fmt.Println(bread)

}
