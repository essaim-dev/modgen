package main

import (
	"flag"
	"log"

	"go.essaim.dev/modgen"
)

var (
	configPath     = flag.String("config", "modgen.yaml", "path of the configuration file")
	targetPath     = flag.String("target", "gen/", "path where the site should be generated")
	indexTmplPath  = flag.String("index-tmpl", "", "path of an optional custom index template")
	moduleTmplPath = flag.String("module-tmpl", "", "path of an optional custom module template")
)

func main() {
	flag.Parse()

	config, err := modgen.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("could not parse config file: %s\n", err)
	}

	g := modgen.NewGenerator(config)

	if *indexTmplPath != "" {
		if err := g.WithIndexTemplate(*indexTmplPath); err != nil {
			log.Fatalf("could not parse custom index template file: %s\n", err)
		}
	}

	if *moduleTmplPath != "" {
		if err := g.WithIndexTemplate(*moduleTmplPath); err != nil {
			log.Fatalf("could not parse custom module template file: %s\n", err)
		}
	}

	if err := g.Generate(*targetPath); err != nil {
		log.Fatalf("could not generate the site: %s\n", err)
	}
}
