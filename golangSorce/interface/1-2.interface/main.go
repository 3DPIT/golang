//오렌지 잼 추가, 에러 발생
package main

import "fmt"

type Jam interface { //관계만 정의
	GetOneSpoon() *SpoonOfStrawberryJam
}

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
func (b *Bread) PutJam(jam Jam) {
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *Bread) String() string {
	return "bread" + b.val
}

//오렌지 잼 추가
type OrangeJam struct {
}

func (s *OrangeJam) GetOneSpoon() *SpoonOfOrangeJam {
	return &SpoonOfOrangeJam{}
}

type SpoonOfOrangeJam struct {
}

func (s *SpoonOfOrangeJam) Stirng() string {
	return "+ Orange"
}
func main() {
	bread := &Bread{}
	//jam := &StrawberryJam{}
	jam := &OrangeJam{}
	bread.PutJam(jam)

	fmt.Println(bread)

}
