//html 태그 추가하기
package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) IsOld() bool {
	return u.Age > 30
}
func main() {
	user := User{Name: "3dpit", Email: "3dpit@naver.com", Age: 20}
	user2 := User{Name: "4dpit", Email: "4dpit@naver.com", Age: 31}
	tmpl, err := template.New("Tmpl1").ParseFiles("template/templ1.tmpl")
	if err != nil {
		panic(err)
	}
	tmpl.ExecuteTemplate(os.Stdout, "templ1.tmpl", user)
	tmpl.ExecuteTemplate(os.Stdout, "templ1.tmpl", user2)
}
