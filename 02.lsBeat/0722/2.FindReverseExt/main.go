package main

import "github.com/huandu/xstrings"

func findExt(fileName string) string {
	var copyExtensions string
	flag := 1
	for index := len(fileName) - 1; index >= 0; index-- {
		if fileName[index] == '.' {
			flag = 0
			break
		}
		copyExtensions += string(fileName[index])
	}
	if flag == 1 { //확장자 없는 경우 에러 처리
		return ""
	}
	copyExtensions = xstrings.Reverse(copyExtensions)

	return copyExtensions
}

//go get github.com/huandu/xstrings 다운 진행
