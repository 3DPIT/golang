package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func viewHandler(writer http.ResponseWriter, reques *http.Request) {
	placeholder := []byte("signture list goes here")
	_, err := writer.Write(placeholder)
	check(err)
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func main() {

	executeTemplate("Start {{if .}}Dot is true ! {{end}}finish\n", true)
	executeTemplate("Start {{if .}}Dot is true ! {{end}}finish\n", false)
}
