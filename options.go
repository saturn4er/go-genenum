package genenum

import (
	"io"
	"os"
	"text/template"
)

type Option = func(generator *Generator)

func WithOutput(out io.Writer) func(*Generator) {
	return func(g *Generator) {
		g.out = out
	}
}

func WithOutputFile(filename string) func(*Generator) {
	return func(g *Generator) {
		out, err := os.Create(filename)
		if err != nil {
			panic(err)
		}

		g.out = out
	}
}

func WithFuncs(funcs template.FuncMap) func(*Generator) {
	return func(g *Generator) {
		for name, fn := range funcs {
			g.funcs[name] = fn
		}
	}
}

func WithPackageName(packageName string) func(*Generator) {
	return func(g *Generator) {
		g.generatorContext.packageName = packageName
	}
}
