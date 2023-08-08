package generator

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
	"vago/vago/extractor"
	"vago/vago/input"
	"vago/vago/log/generator"
)

func Build(config input.IOPath) {
	files, err := os.ReadDir(config.InFolder)
	if err != nil {
		log.Fatal(err)
	}

	var logger generator.GeneratorLogger
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		logger.Init(file.Name(), false, false)
		logger.Info("Starting page")

		md, err := os.ReadFile(config.InFolder + file.Name())
		if err != nil {
			logger.Error(err)
		}

		logger.Info("Parsing markdown")
		out := extractor.ParseMarkdown(md)

		logger.Info("Building HTML output")
		BuildPage(config.InTemplate, out, config.OutFolder+file.Name())
		logger.Info("Page built in output folder %s", config.OutFolder)
	}
}

func BuildPage(templ string, content extractor.Out, filename string) {
	t := template.Must(template.ParseFiles(templ))
	filename = changeFileExtension(filename, "html")
	file := createFile(filename)
	err := t.Execute(file, content)
	if err != nil {
		panic(err)
	}

	if err := file.Close(); err != nil {
		panic(err)
	}
}

func createFile(filename string) *os.File {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	return f
}

func changeFileExtension(filename string, extension string) string {
	name := strings.TrimSuffix(filename, filepath.Ext(filename))

	return name + "." + extension
}
