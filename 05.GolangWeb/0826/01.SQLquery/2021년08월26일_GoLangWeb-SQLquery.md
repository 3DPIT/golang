## 2021년08월26일_GoLangWeb-SQLquery

- DB를 열고 Close를 해야하는데 이 인스턴스 생명주기 관리해야하는데 이패키지에서는 알 수 없음
- 이패키지 안에서 핸들러 언제 쓰고 끝나는지 몰라서 이것을 사용하는 쪽으로 권한을 넘겨야함

- dbHandler 부분에 close() 추가

```go
type DBHandler interface {
	GetTodos() []*Todo
	AddTodo(name string) *Todo
	RemoveTodo(id int) bool
	DompleteTodo(id int, complete bool) bool
	Close()
}

var handler DBHandler
```

- 외부 노출로 변경해준다.

```go
func GetTodos() []*Todo {
	return handler.getTodos()
}
func AddTodo(name string) *Todo {
	return handler.addTodo(name)
}

func RemoveTodo(id int) bool {
	return handler.removeTodo(id)
}

func CompleteTodo(id int, complete bool) bool {
	return handler.completeTodo(id, complete)
}
```

- 기능들 삭제

```go
//init

func NewDBHandler() DBHandler {
	//handler = newMemoryHandler()
	//todoMap = make(map[int]*Todo)
	return newSqliteHandler()
}
```

- 위와 같이 변경해주고 나머지도 다 맞춰준다.

### memoryHandler.go

```go
package model

import "time"

type memoryHandler struct {
	todoMap map[int]*Todo
}

func (m *memoryHandler) GetTodos() []*Todo {
	list := []*Todo{}
	for _, v := range m.todoMap {
		list = append(list, v)
	}
	return list
}

func (m *memoryHandler) AddTodo(name string) *Todo {
	id := len(m.todoMap) + 1
	todo := &Todo{id, name, false, time.Now()}
	m.todoMap[id] = todo
	return todo
}

func (m *memoryHandler) RemoveTodo(id int) bool {
	if _, ok := m.todoMap[id]; ok {
		delete(m.todoMap, id)
		return true
	}
	return false
}

func (m *memoryHandler) CompleteTodo(id int, complete bool) bool {
	if todo, ok := m.todoMap[id]; ok {
		todo.Completed = complete
		return true
	}
	return false
}
func newMemoryHandler() DBHandler {
	m := &memoryHandler{}
	m.todoMap = make(map[int]*Todo)
	return m
}

```

### sqliteHandler.go

```go
package model

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteHandler struct {
	db *sql.DB
}

func (s *sqliteHandler) GetTodos() []*Todo {
	return nil
}
func (s *sqliteHandler) AddTodo(name string) *Todo {
	return nil
}
func (s *sqliteHandler) RemoveTodo(id int) bool {
	return false
}
func (s *sqliteHandler) CompleteTodo(id int, complete bool) bool {
	return false
}

func (s *sqliteHandler) Close() {
	s.db.Close()
}
func newSqliteHandler() DBHandler {
	database, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	//테이블을 만들기 위한 스키마
	statement, _ := database.Prepare(
		`CREATE TABLE IF NOT EXISTS todos (
			id        INTEGER  PRIMARY KEY AUTOINCREMENT,
			name      TEXT,
			completed BOOLEAN,
			createdAt DATETIME
		)`)
	statement.Exec()
	return &sqliteHandler{db: database}
}

```

## app.go

```go
type AppHandler struct {
	http.Handler
  model.DBHandler
}
```

- 이부분도 close 권한 있어야함 패키지이기 때문에 하지만 저렇게 선언 하면 상속의 개념 처럼 isa 관계는 아니고 포함 관계임
- 암시적으로 안쓰임

## app.go

```go
package app

import (
	"net/http"
	"strconv"

	"web/web1/model"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rd *render.Render = render.New()

type AppHandler struct {
	http.Handler
	db model.DBHandler
}

func (a *AppHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func (a *AppHandler) getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := a.db.GetTodos()
	rd.JSON(w, http.StatusOK, list)
}

func (a *AppHandler) addTodoListHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	todo := a.db.AddTodo(name)
	rd.JSON(w, http.StatusCreated, todo)
}

type Success struct {
	Success bool `json:"success"`
}

func (a *AppHandler) removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ok := a.db.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}

}

func (a *AppHandler) completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"
	ok := a.db.CompleteTodo(id, complete)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}

}

func MakeHandler() *AppHandler {
	//todoMap = make(map[int]*Todo)

	r := mux.NewRouter()
	a := &AppHandler{
		Handler: r,
		db:      model.NewDBHandler(),
	}

	r.HandleFunc("/todos", a.getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", a.addTodoListHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", a.removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/complete-todo/{id:[0-9]+}", a.completeTodoHandler).Methods("GET")
	r.HandleFunc("/", a.indexHandler)

	return a
}

```

## GetTodo

```go
func (s *sqliteHandler) GetTodos() []*Todo {
	todos := []*Todo{}
	rows, err := s.db.Query("SELECT id, name, completed, createdAt FROM todos")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.ID, &todo.Name, &todo.Completed, &todo.CreatedAt)
		todos = append(todos, &todo)
	}
	return todos
}
```

## AddTodo

```go
func (s *sqliteHandler) AddTodo(name string) *Todo {
	stmt, err := s.db.Prepare("INSERT INTO todos (name, completed, createdAt) VALUES (?, ?,datetime('now'))")
	if err != nil {
		panic(err)
	}
	rst, err := stmt.Exec(name, false)
	if err != nil {
		panic(err)
	}
	id, _ := rst.LastInsertId()
	var todo Todo
	todo.ID = int(id)
	todo.Name = name
	todo.Completed = false
	todo.CreatedAt = time.Now()
	return &todo
}
```

## app_test.go

```go
func TestTodos(t *testing.T) {
	os.Remove("./test.db")
	assert := assert.New(t)
	ah := MakeHandler()
	defer ah.db.Close()
	ts := httptest.NewServer(ah)
	defer ts.Close()
```

- 이부분 이렇게 수정해야함
- DB를 삭제하고 제대로 되어야 에러가 안생김

## CompletedTodo

```go
func (s *sqliteHandler) CompleteTodo(id int, complete bool) bool {
	stmt, err := s.db.Prepare("UPDATE todos SET completed = ? WHERE id==?")
	if err != nil {
		panic(err)
	}
	rst, err := stmt.Exec(complete, id)
	if err != nil {
		panic(err)
	}

	cnt,_:=rst.RowsAffected()
	return cnt>0
}
```

- ​	cnt,_:=rst.RowsAffected() == 1이면 바뀐거 0 이면 안 바뀜

## RemoveTodo

```go
func (s *sqliteHandler) RemoveTodo(id int) bool {
	stmt, err := s.db.Prepare("DELETE FROM todos WHERE id =?")
	if err != nil {
		panic(err)
	}
	rst, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	cnt, _ := rst.RowsAffected()
	return cnt > 0
}
```

- 여기까지 제대로 실행됨을 확인

- filedb로만들면 db가 꺼져도 데이터가 그대로 남아 있다.

