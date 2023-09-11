package input

import (
	"gopkg.in/yaml.v3"
	"os"
)

func GetTheme(filename string) map[string]interface{} {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}

	err = yaml.Unmarshal(f, &data)
	if err != nil {
		panic(err)
	}

	return data
}
