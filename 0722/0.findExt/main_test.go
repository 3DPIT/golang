package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindExtension(t *testing.T) {
	assert := assert.New(t)

	testPath := []string{"/go/src/txt.txt", "go/src/test.txt", "go/test1.t", "go/t.SMF/do.txt", ".txt"}
	//test1 일반적인 경우
	assert.Equal(".txt", findExtension(testPath[0]))
	//test2 일반적인 경우2
	assert.Equal(".txt", findExtension(testPath[1]))
	//test3 일반적인 확장자가 아닌경우도 . 이후로 가져오는지 확인
	assert.Equal(".t", findExtension(testPath[2]))
	//test4 폴더명이 .으로 끝나는경우 걸러지는 아닌지 확인
	assert.Equal(".txt", findExtension(testPath[3]))
	//test5 경로없이 바로 파일이 있는 경우
	assert.Equal(".txt", findExtension(testPath[4]))

}
