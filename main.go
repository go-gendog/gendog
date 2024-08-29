package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func main() {
	// Open the templates folder
	tmpl := "templates/"
	dirEntries, err := os.ReadDir(tmpl)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	// Read all files into a []string
	var templates []string
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			file, err := os.Open(filepath.Join(tmpl, entry.Name()))
			if err != nil {
				log.Fatalf("Failed to open file %s: %v", entry.Name(), err)
			}
			defer func(file *os.File) {
				_ = file.Close()
			}(file)

			content, err := io.ReadAll(file)
			if err != nil {
				log.Fatalf("Failed to read file %s: %v", entry.Name(), err)
			}

			templates = append(templates, string(content))
		}
	}

	// Create a new text/template with each string and execute with an example map
	data := map[string]string{"Name": "bob"}
	for _, tmplStr := range templates {
		tmpl, err := template.New("example").Parse(tmplStr)
		if err != nil {
			log.Fatalf("Failed to parse template: %v", err)
		}
		err = tmpl.Execute(os.Stdout, data)
		if err != nil {
			log.Fatalf("Failed to execute template: %v", err)
		}
		fmt.Println()
	}
}
