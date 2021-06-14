package main

import (
	"html/template"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func main() {
	templateText := "Before loop: {{.}}\n{{range .}}In loop :{{.}}\n{{end}}After loop: {{.}}\n"
	executeTemplate(templateText, []string{"do", "re", "mi"})

	templateText = "Price:\n{{range .}}${{.}}\n{{end}}"
	executeTemplate(templateText, []float64{1.25, 0.99, 27})

}
