package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"go.yaml.in/yaml/v4"
)

type FAQ struct {
	Items []struct {
		Q string `yaml:"q"`
		A string `yaml:"a"`
	} `yaml:"items"`
}

// builds the #-link for connecting questions with answers
func toURLFragment(s string) string {
	s = strings.ToLower(s)
	// Replace spaces with hyphens
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, "-")
	// Remove special characters
	s = regexp.MustCompile(`[^a-z0-9-]`).ReplaceAllString(s, "")
	return s
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: mdfaqgen <input.yaml> <output.md>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Read YAML file
	yamlData, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Parse YAML
	var faq FAQ
	err = yaml.Unmarshal(yamlData, &faq)
	if err != nil {
		fmt.Printf("Error parsing YAML: %v\n", err)
		os.Exit(1)
	}

	// Create markdown content
	var md strings.Builder

	// Questions list
	md.WriteString("# Frequently Asked Questions\n\n")
	for _, item := range faq.Items {
		urlFragment := toURLFragment(item.Q)
		md.WriteString(fmt.Sprintf("- [%s](#%s)\n", item.Q, urlFragment))
	}

	// Questions and answers
	for _, item := range faq.Items {
		urlFragment := toURLFragment(item.Q)
		md.WriteString(fmt.Sprintf("<a name=\"%s\"></a>\n", urlFragment))
		md.WriteString(fmt.Sprintf("## %s\n\n", item.Q))
		md.WriteString(item.A + "\n\n")
	}

	// Write markdown file
	err = ioutil.WriteFile(outputFile, []byte(md.String()), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully created %s\n", outputFile)
}

