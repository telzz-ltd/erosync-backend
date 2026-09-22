package main

import (
	"html/template"
	"log"
	"os"
	"path"
	"time"
)

func main() {
	cwd, _ := os.Getwd()
	file1 := path.Join(cwd, "templates", "layout.html")
	file2 := path.Join(cwd, "templates", "welcome.html")
	tmpl, err := template.ParseFiles(file1, file2)
	if err != nil {
		panic(err)
	}

	log.Println(tmpl.Execute(os.Stdout, map[string]any{
		"subject": "Verify Email",
		"name":    "Usman",
		"appUrl":  "http://localhost:8080",
		"year":    time.Now().Year(),
		"appName": "Erosync",
	}))
}
