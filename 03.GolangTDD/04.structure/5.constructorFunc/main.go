package main

type dict struct {
	data map[int]string
}

//생성자 함수 정의
func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d
}
func main() {
	dic := newDict()
	dic.data[1] = "A"
	dic.data[2] = "B"
}
