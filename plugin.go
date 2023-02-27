package main

import (
	"bytes"
	"io"
	"os"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/wellington/go-libsass"
)

type SassPluginOptions interface{}

func GetSassPlugin(opt SassPluginOptions) api.Plugin {
	var sassplugin = api.Plugin{
		Name: "sass",
		Setup: func(build api.PluginBuild) {
			build.OnResolve(api.OnResolveOptions{Filter: `\.sa|css$`}, func(args api.OnResolveArgs) (api.OnResolveResult, error) {
				return api.OnResolveResult{
					Path:      args.Path,
					Namespace: args.Namespace,
				}, nil
			})

			build.OnLoad(api.OnLoadOptions{Filter: `\.sa|css$`}, func(args api.OnLoadArgs) (api.OnLoadResult, error) {
				r, _ := os.Open(args.Path)
				var b bytes.Buffer
				w := io.Writer(&b)
				c, _ := libsass.New(w, r)
				c.Run()
				s := b.String()
				return api.OnLoadResult{
					Contents: &s,
					Loader:   api.LoaderCSS,
				}, nil
			})
		},
	}
	return sassplugin
}
