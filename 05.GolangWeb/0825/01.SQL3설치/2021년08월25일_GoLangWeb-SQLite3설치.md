## 2021년08월25일_GoLangWeb-SQLite3설치

- RDBMS 는 관계형 데이터 베이스라고 함
- 데이터베이스의 관계를 정할수 있는 DB , SQL Query를 사용할 수 있는 것

- MysQL , MsSQL, Oracle DB, Amazon(aurora) 등 있음

- SQLite 는 파일로 만 동작하는 DB고 RDBMS, SQL query제공해서 간단하게 사용 가능하기 좋음

## go에서 사용하는 go-sqlLite3

- CGo를 사용 이것은 C-Library 사용하게 해줌
- 표준-c 컴파일러가 필요한데 window에는 없는데 , gcc, clang을 사용하지만 
- 고랭은 gnu c compile , gcc를 사용
- ms window는 지원을 하지 않음 ㅜㅜ
- 그래서 사람들이 ms window에서도gcc 동작 할 수 있게 mingW라는걸 만듦 여기서 gcc를 돌려야함
- 복잡하긴하다.

## tdm-gcc 설치 

[설치링크](https://jmeubank.github.io/tdm-gcc/)