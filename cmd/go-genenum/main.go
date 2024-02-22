package main

import (
	"flag"
	"log"
	"os"

	"github.com/saturn4er/go-genenum"
)

type Args struct {
	OutputFile  string
	PackageName string
}

func main() {
	if len(os.Args) < 2 {
		log.Println("Usage: genenum <yaml file>")
		os.Exit(1)
	}

	var args Args
	flag.StringVar(&args.OutputFile, "o", "", "output file")
	flag.StringVar(&args.PackageName, "p", "", "go package name")
	flag.Parse()

	config, err := genenum.LoadConfig(flag.Arg(0))
	if err != nil {
		log.Printf("load args: %v\n", err)
		os.Exit(1)
	}

	var options []genenum.Option
	if args.OutputFile != "" {
		options = append(options, genenum.WithOutputFile(args.OutputFile))
	}
	if args.PackageName != "" {
		options = append(options, genenum.WithPackageName(args.PackageName))
	}

	generator, err := genenum.NewGenerator(config, options...)
	if err != nil {
		log.Printf("init generator: %v\n", err)
		os.Exit(1)
	}

	if err := generator.Generate(); err != nil {
		log.Printf("generate: %v\n", err)
		os.Exit(1)
	}
}
