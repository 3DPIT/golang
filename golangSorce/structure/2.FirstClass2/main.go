package main

import "fmt"

type Student struct {
	name   string
	class  int
	record Record
}
type Record struct {
	name  string
	grade string
}

func (s Student) ViewRecord() {
	fmt.Println(s.record)
}

func ViewRecord(s Student) { //그냥 함수
	fmt.Println(s.record)
}
func (s *Student) InputRecord(name string, grade string) {
	s.record.name = name
	s.record.grade = grade
}
func main() {
	var s Student
	s.name = "홍길동"
	s.class = 1
	s.record.name = "SW"
	s.record.grade = "A"
	s.ViewRecord()

	s.InputRecord("WS", "C")
	s.ViewRecord()
}
