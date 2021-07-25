//목표 숫자 찾기 게임 만들기
//1. 플레이어가 추측한 숫자 난수로 생성해서 저장
//2. 플레이어가 추측할 수 있게 입력받고 하는것 만들기
/*3. 플레이어가 입력한 숫자 목표값보다 낮으면 : "Oops. Your guess was LOW"
높으면 : "Oops. Your guess was High*/
//4. 플레이어는 10까지 예측 가능 추측할때마다 남은 횟수 알려주기
//5. 플레이어 목표값 찾으면 "Good job! You guessed it!" 종료
//6. 추측횟수 안에 못맞추면 "Sorry. You didn't guess my number. It was: [target]" 출력
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix() //날짜 정수값으로 변환
	rand.Seed(seconds)           //시드생성
	target := rand.Intn(100) + 1 //난수 생성
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	fmt.Println(target)

	reader := bufio.NewReader(os.Stdin) //키보드 입력값 저장
	fmt.Print("Make a guess: ")
	input, err := reader.ReadString('\n') //엔터를 기준으로 입력 종료
	if err != nil {                       //에러 체크
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)  //공백 제거
	guess, err := strconv.Atoi(input) //입력 문자열 정수값 반환
	if err != nil {                   //에러 체크
		log.Fatal(err)
	}

}
