package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindExt(t *testing.T) {
	assert := assert.New(t)

	fileName1 := "1.txt"
	fileName2 := "2.SMF"
	fileName3 := "3.001"
	fileName4 := "4.txt.SMF.001"
	fileName5 := "txt"

	assert.Equal("txt", findExt(fileName1))
	assert.Equal("SMF", findExt(fileName2))
	assert.Equal("001", findExt(fileName3))
	assert.Equal("001", findExt(fileName4))
	assert.Equal("", findExt(fileName5))

}
