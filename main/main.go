package main

import (
	"embed"
	"flag"
	"io/fs"
	"os"
	"text/template"
)

//go:embed project_template/*
var projectFiles embed.FS

var newFolderStructure = []string{
	"%v/handlers", "%v/models", "%v/routes", "%v/main",
}

func createNewProjectStructure(s string) {
	tmplFiles, _ := fs.ReadDir(projectFiles, s)

	for _, tmpl := range tmplFiles {
		if tmpl.IsDir() {
			continue
		}

		pt, _ := template.ParseFS(projectFiles, s+"/"+tmpl.Name())

		file, _ := os.Create(s + "/" + tmpl.Name())
		defer file.Close()

		// apply the template to the vars map and write the result to file.
		_ = pt.Execute(file, nil)
	}

}

func CreateApp(s string) error {
	createNewProjectStructure(s)
	return nil
}

func main() {
	flag.Func("create-app", "DjanGopher Manage", CreateApp)
	flag.Parse()
}
