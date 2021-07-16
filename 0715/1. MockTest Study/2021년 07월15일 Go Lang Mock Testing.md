## 2021년 07월15일 Go Lang Mock Testing
- golang의 gomark라는 패키지를 사용 - 테스트에 사용되는 매우 인기 있는 모듈  
- 필수모듈 두개 gomark와 gomark 관련 필수 모듈(mock체인) 설치 
- 스텁 코드를 위한 것
## 필요 명령
```go
go get github.com/golang/mock/gomock
go get github.com/golang/mock/mockgen
```
## 폴더 파일 구성  
![image-20210715092923824](2021년 07월15일 Go Lang Mock Testing.assets/image-20210715092923824.png)
## IUser.go  
```go
package IUser

type IUserInterface interface {
	AddUser(int, string) error
}
```
## User.go  
```go
package User

import (
	"fileMonitoring/Golang.MockTest/IUser"
)

type User struct {
	IUser IUser.IUserInterface
}

func (u *User) Use() {
	u.IUser.AddUser(1, "sample test")
}
```
- 우선 go get github.com/golang/mock/gomock 와 go get github.com/golang/mock/mockgen 이것을 두개 설치
```go
mockgen -destination=mocks/mockRunner.go -package=mocks fileMonitoring/Golang.MockTest/IUser IUserInterface
```
-  위의 형식은 
mockgen -destination=**<원하는 폴더 이름>**/**<원하는 파일이름.go>** -package=**<원해당 인터페이스있는패키지위치>** **<인터페이스 이름>**
go test -v moduleName testFile
- 위명령어를 해주면 해당 인터페이스 있는 위치에 mocks라는 폴더에 mockRunner.go가 생성됨 
![image-20210715100243415](2021년 07월15일 Go Lang Mock Testing.assets/image-20210715100243415.png)
- 여기까지 하면 준비가 완료된 상태
- 스텁 메소드를 생성한 것 
- 우리가 생성한곳에 보면 새 모크 인터페이스 생성하는 법이 있다.
![image-20210715100904844](2021년 07월15일 Go Lang Mock Testing.assets/image-20210715100904844.png)
## User_test.go  
```go
package User_test

import (
	"testing"

	"fileMonitoring/Golang.MockTest/User"
	"fileMonitoring/Golang.MockTest/mocks"

	"github.com/golang/mock/gomock"
)

func TestUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUser := mocks.NewMockIUserInterface(mockCtrl)
	testUser := &User.User{IUser: mockUser}
	mockUser.EXPECT().AddUser(1, "sample test").Return(nil).Times(1)
	testUser.Use()

}
```
- 필요한것은 인터페이스로 구현을 해놓으면 그것을 mock로 만들어서 저렇게 자동화 테스트를 진행할 수 있음 
## 참고 사이트 
[Go Lang Mock Testing 참고동영상](https://www.youtube.com/watch?v=pzAgivQ9DHI)

