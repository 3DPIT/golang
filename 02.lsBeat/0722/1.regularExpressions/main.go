package main

import (
	"regexp"
)

func regularExpressionFind(fileName string) string {
	r, _ := regexp.Compile("\\.[a-zA-Z0-9가-힣]+")
	return r.FindString(fileName)
}
