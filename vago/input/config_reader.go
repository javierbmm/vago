package input

import (
	"gopkg.in/yaml.v3"
	"os"
)

type ConfigYAML struct {
	Input    string
	Output   string
	Template string
	Styles   string
	Theme    string
}

func ReadYAML(filename string) ConfigYAML {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var cyaml ConfigYAML
	err = yaml.Unmarshal(f, &cyaml)
	if err != nil {
		panic(err)
	}

	return cyaml
}

func (c ConfigYAML) AsIOPath() IOPath {
	iopath := IOPath{
		InFolder:     c.Input,
		InTemplate:   c.Template,
		OutFolder:    c.Output,
		StylesFolder: c.Styles,
		InTheme:      c.Theme,
	}

	return iopath
}
