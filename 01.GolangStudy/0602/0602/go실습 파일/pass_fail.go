//pass_fail 프로그램은 성적의 합격 여부를 알려 줍니다.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")          //사용자에게 성적 입려하라고 알려줌
	reader := bufio.NewReader(os.Stdin)   //키보드로 부터 텍스트 읽어오기
	input, err := reader.ReadString('\n') //엔터키가 눌리면 입력된 문자 반환
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(input)               // 사용자가 입력한 값 출력
	input = strings.TrimSpace(input) //공백 제거
	score, err := strconv.ParseFloat(input, 64)
	var status string
	if score >= 60 {
		status = "pass"
	} else {
		status = "fail"
	}
	fmt.Println(status)
}
