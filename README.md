# mdfaqgen

A Go program that converts structured YAML input into a clean, formatted FAQ.md file in Markdown. This tool automatically generates #-links that connect questions with their corresponding answers, ensuring that the questions and answers remain in sync and preventing issues such as copy and paste errors or misaligned content. Additionally, it allows for further formatting of the answers in Markdown, providing flexibility in how the FAQ is presented.

For a practical demonstration, an example [FAQ.md](/FAQ.md) is included in this repository, generated from the [faq.yaml](/faq.yaml) file.

## Getting Started

```
go run main.go input.yaml output.md
```
