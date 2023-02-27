package main

import (
	sassplugin "github.com/hsuehic/esbuild-sass"
)

func main() {
	var opt sassplugin.SassPluginOptions
	plugin := sassplugin.GetSassPlugin(opt)
	print(plugin.Name)
}
