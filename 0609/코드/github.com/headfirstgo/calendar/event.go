package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	Data  //데이트타입을 입베딩
}

func (e *Event) Title() string { //접근자 메서드
	return e.title
}
func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}
