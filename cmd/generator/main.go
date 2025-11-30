package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aag2807/portfolio/internal/portfolio"
)

func main() {
	dataPath := flag.String("data", "portfolio.json", "Path to portfolio JSON data file")
	templatePath := flag.String("template", "templates/index.html", "Path to HTML template")
	staticPath := flag.String("static", "static", "Path to static files directory")
	outputPath := flag.String("output", "dist", "Output directory for generated site")
	flag.Parse()

	generator := portfolio.NewGenerator(*dataPath, *templatePath, *staticPath, *outputPath)

	fmt.Println("Loading portfolio data...")
	if err := generator.LoadData(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Generating static site...")
	if err := generator.Generate(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Site generated successfully in '%s' directory\n", *outputPath)
}

