## 2021년 06월11일 golang, UnitTest 학습 내용
### 1. 완료조건
- [x] Head First Go 고루틴과 채널에 대한 학습과 이해
- [x] Head First Go 자동테스트의 구현과 이해
- [x] Wiki UnitTest Assert, Arrange, 라이프사이클 이해
### 2. 상세내용
- 직접 메소드등에 고루틴 적용하여 이해하고 실습 진행
- 직접 자동화 테스트 진행 하면서 이해하고 실습 진행
- Assert, Arrange, 단위테스트 함수 라이프사이클 실습 진행
### 3. 질문
- [ ] go언어에서 자동화 테스트를 진행 하면 채널에 그 데이터를 집어넣고 결과적으로 채널의 값을 집어넣은 순서대로 출력하는 것을 진행했는데 특정 채널의 의 데이터 값만 가져 올 수 있나 방법이 있나요?  소스는 아래와 같습니다.
```go
  package main

import "fmt"

func odd(channel chan int) {
	channel <- 1
	channel <- 3
}
func even(channel chan int) {
	channel <- 2
	channel <- 4
}
func main() {
	channelA := make(chan int)
	channelB := make(chan int)
	go odd(channelA)
	go even(channelB)
	fmt.Println(<-channelA)
	fmt.Println(<-channelA)
	fmt.Println(<-channelB)
	fmt.Println(<-channelB)
}
```
### 4. 참여자
[고형호](@HyungHo.Ko)