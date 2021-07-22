package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegularExpressionCheck(t *testing.T) {
	assert := assert.New(t)
	fileName1 := "test1.txt"
	fileName2 := "test1.001"
	fileName3 := "test1.SMF"
	fileName4 := "test1.SMF.txt"

	assert.Equal(".txt", regularExpressionFind(fileName1))
	assert.Equal(".001", regularExpressionFind(fileName2))
	assert.Equal(".SMF", regularExpressionFind(fileName3))
	assert.Equal(".txt", regularExpressionFind(fileName4)) // .txt를 못찾고 .SMF를 찾음

}
