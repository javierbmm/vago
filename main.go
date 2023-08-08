package main

import (
	"vago/vago/generator"
	"vago/vago/input"
	"vago/vago/server"
)

var config input.IOPath

func main() {
	config = input.ReadYAML("./config.yaml").AsIOPath()
	generator.Build(config)
	server.Serve(8080, config.OutFolder)
}
