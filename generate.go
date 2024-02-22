package genenum

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"text/template"

	"github.com/gertd/go-pluralize"
	"github.com/stoewer/go-strcase"
	"golang.org/x/tools/imports"

	enumTemplate "github.com/saturn4er/go-genenum/template"
)

type Generator struct {
	config           []*ConfigEnum
	out              io.Writer
	templateContent  string
	funcs            template.FuncMap
	generatorContext *goGeneratorContext
}

func NewGenerator(config []*ConfigEnum, opts ...Option) (*Generator, error) {
	templateContent, err := enumTemplate.Load()
	if err != nil {
		return nil, err
	}

	generatorContext := &goGeneratorContext{
		packageName: "enums",
	}

	result := &Generator{
		config:           config,
		out:              os.Stdout,
		templateContent:  templateContent,
		generatorContext: generatorContext,
		funcs: map[string]any{
			"import":     generatorContext.packageImport,
			"uCamelCase": strcase.UpperCamelCase,
			"lCamelCase": strcase.LowerCamelCase,
			"snakeCase":  strcase.SnakeCase,
			"plural":     pluralize.NewClient().Plural,
			"addInts": func(values ...int) int {
				result := 0
				for _, value := range values {
					result += value
				}

				return result
			},
		},
	}

	for _, opt := range opts {
		opt(result)
	}

	return result, nil
}

func (g *Generator) Generate() error {
	templateInstance, err := template.New("enum").
		Funcs(g.funcs).
		Parse(g.templateContent)
	if err != nil {
		return fmt.Errorf("parse template: %w", err)
	}

	bodyBuf := bytes.NewBuffer(nil)
	if err := templateInstance.Execute(bodyBuf, map[string]any{
		"enums": g.config,
	}); err != nil {
		return fmt.Errorf("execute template: %w", err)
	}

	fileBuf := bytes.NewBuffer(nil)

	if _, err := fileBuf.Write([]byte(g.generatorContext.head())); err != nil {
		return fmt.Errorf("write head: %w", err)
	}

	if _, err := fileBuf.Write(bodyBuf.Bytes()); err != nil {
		return fmt.Errorf("write content: %w", err)
	}

	imports.LocalPrefix = ""
	// execute goimports
	formattedResult, err := imports.Process("name", fileBuf.Bytes(), &imports.Options{
		TabIndent:  true,
		TabWidth:   8,
		Comments:   true,
		Fragment:   true,
		FormatOnly: false,
	})
	if err != nil {
		if _, err := g.out.Write(fileBuf.Bytes()); err != nil {
			return fmt.Errorf("write result: %w", err)
		}

		return fmt.Errorf("goimports: %w", err)
	}

	_, err = g.out.Write(formattedResult)
	if err != nil {
		return fmt.Errorf("write result: %w", err)
	}

	return nil
}
