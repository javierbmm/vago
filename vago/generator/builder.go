package generator

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
	"vago/vago/extractor"
	"vago/vago/input"
	log2 "vago/vago/log"
)

func Build(config input.IOPath, noLog bool, noTime bool) {
	buildPages(config, noLog, noTime)
	buildStyles(config, noLog, noTime)
}

func buildPages(config input.IOPath, noLog bool, noTime bool) {
	files, err := os.ReadDir(config.InFolder)
	if err != nil {
		log.Fatal(err)
	}

	var logger log2.GeneratorLogger
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		logger.Init(file.Name(), noLog, noTime)
		logger.Info("Starting page")

		md, err := os.ReadFile(config.InFolder + file.Name())
		if err != nil {
			logger.Error(err)
		}

		logger.Info("Parsing markdown")
		out := extractor.ParseMarkdown(md)

		logger.Info("Building HTML output")
		buildPage(config.InTemplate, out, config.OutFolder+file.Name())
		logger.Info("Page built in output folder %s", config.OutFolder)
	}
}

func buildStyles(config input.IOPath, noLog bool, noTime bool) {
	var logger log2.GeneratorLogger
	theme := input.GetTheme(config.InTheme)
	styles, err := os.ReadDir(config.StylesFolder)
	if err != nil {
		panic(err)
	}

	for _, style := range styles {
		if style.IsDir() {
			continue
		}

		logger.Init(style.Name(), noLog, noTime)
		logger.Info("Parsing theme.")
		buildStyle(config.StylesFolder+style.Name(), theme, config.OutFolder+style.Name())
		logger.Info("Style built in output folder %s", config.OutFolder)
	}
}

func buildPage(templ string, content extractor.Out, filename string) {
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

func buildStyle(styles string, theme map[string]interface{}, filename string) {
	t := template.Must(template.ParseFiles(styles))
	filename = changeFileExtension(filename, "css")
	file := createFile(filename)
	err := t.Execute(file, theme)
	if err != nil {
		panic(err)
	}

	if err := file.Close(); err != nil {
		panic(err)
	}
}

func changeFileExtension(filename string, extension string) string {
	name := strings.TrimSuffix(filename, filepath.Ext(filename))

	return name + "." + extension
}
