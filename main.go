package main

import (
	"os"
	"vago/vago/extractor"
)

func main() {
	md, err := os.ReadFile("./source/first.md")
	if err != nil {
		panic(err)
	}
	extractor.ParseMarkdown(md)
}
