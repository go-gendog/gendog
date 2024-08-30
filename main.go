package main

import (
	"embed"
	"log"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

//go:embed templates/*
var templates embed.FS

func main() {
	// Read the template files from the templates
	// directory in the embedded filesystem.
	dirEntries, err := templates.ReadDir("templates")
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	// Read all files into a map
	tmpl := make(map[string]string, len(dirEntries))
	for _, entry := range dirEntries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".tmpl") {
			content, err := templates.ReadFile("templates/" + entry.Name())
			if err != nil {
				log.Fatalf("Failed to read file %s: %v", entry.Name(), err)
			}
			name := strings.TrimSuffix(entry.Name(), ".tmpl")
			tmpl[name] = string(content)
		}
	}

	loader := openapi3.NewLoader()
	doc, err := loader.LoadFromFile("my-openapi-spec.json")

}
