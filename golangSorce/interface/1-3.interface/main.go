//오렌지 잼 추가 에러 해결 인터페이스 제대로 구현
package main

import "fmt"

//인터페이스
type SpoonOfJam interface {
	String() string
}
type Jam interface { //관계만 정의
	GetOneSpoon() SpoonOfJam
}

//딸기 잼
type StrawberryJam struct {
}
type SpoonOfStrawberryJam struct {
}

func (s *SpoonOfStrawberryJam) String() string {
	return " + strawberry"
}
func (j *StrawberryJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{}
}

//오렌지 잼 추가
type OrangeJam struct {
}
type SpoonOfOrangeJam struct {
}

func (s *SpoonOfOrangeJam) String() string {
	return " + Orange"
}

func (j *OrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}
}

//빵
type Bread struct {
	val string
}

func (b *Bread) String() string {
	return "bread" + b.val
}
func (b *Bread) PutJam(jam Jam) {
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func main() {
	bread := &Bread{}
	//jam := &StrawberryJam{}
	jam := &OrangeJam{}
	bread.PutJam(jam)

	fmt.Println(bread)

}
