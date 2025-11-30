package portfolio

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Generator struct {
	dataPath     string
	templatePath string
	staticPath   string
	outputPath   string
	portfolio    *Portfolio
}

func NewGenerator(dataPath string, templatePath string, staticPath string, outputPath string) *Generator {
	return &Generator{
		dataPath:     dataPath,
		templatePath: templatePath,
		staticPath:   staticPath,
		outputPath:   outputPath,
	}
}

func (g *Generator) LoadData() error {
	data, err := os.ReadFile(g.dataPath)
	if err != nil {
		return fmt.Errorf("failed to read portfolio data: %w", err)
	}

	g.portfolio = &Portfolio{}
	if err := json.Unmarshal(data, g.portfolio); err != nil {
		return fmt.Errorf("failed to parse portfolio data: %w", err)
	}

	return nil
}

func (g *Generator) Generate() error {
	if g.portfolio == nil {
		return fmt.Errorf("portfolio data not loaded")
	}

	if err := os.MkdirAll(g.outputPath, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	funcMap := template.FuncMap{
		"join": strings.Join,
		"currentYear": func() int {
			return time.Now().Year()
		},
		"categoryColor": func(category string) string {
			colors := map[string]string{
				"backend":  "bg-emerald-500/20 text-emerald-400 border-emerald-500/30",
				"frontend": "bg-blue-500/20 text-blue-400 border-blue-500/30",
				"database": "bg-amber-500/20 text-amber-400 border-amber-500/30",
				"mobile":   "bg-purple-500/20 text-purple-400 border-purple-500/30",
				"systems":  "bg-red-500/20 text-red-400 border-red-500/30",
				"styling":  "bg-pink-500/20 text-pink-400 border-pink-500/30",
			}
			if color, ok := colors[category]; ok {
				return color
			}
			return "bg-slate-500/20 text-slate-400 border-slate-500/30"
		},
		"categoryColorLight": func(category string) string {
			colors := map[string]string{
				"backend":  "bg-emerald-100 text-emerald-700 border-emerald-300",
				"frontend": "bg-blue-100 text-blue-700 border-blue-300",
				"database": "bg-amber-100 text-amber-700 border-amber-300",
				"mobile":   "bg-purple-100 text-purple-700 border-purple-300",
				"systems":  "bg-red-100 text-red-700 border-red-300",
				"styling":  "bg-pink-100 text-pink-700 border-pink-300",
			}
			if color, ok := colors[category]; ok {
				return color
			}
			return "bg-slate-100 text-slate-700 border-slate-300"
		},
		"githubUsername": func(url string) string {
			parts := strings.Split(url, "/")
			if len(parts) > 0 {
				return parts[len(parts)-1]
			}
			return url
		},
	}

	tmpl, err := template.New("index.html").Funcs(funcMap).ParseFiles(g.templatePath)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	outputFile := filepath.Join(g.outputPath, "index.html")
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, g.portfolio); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	if err := g.copyStaticFiles(); err != nil {
		return fmt.Errorf("failed to copy static files: %w", err)
	}

	return nil
}

func (g *Generator) copyStaticFiles() error {
	return filepath.Walk(g.staticPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(g.staticPath, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(g.outputPath, relPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, 0755)
		}

		return copyFile(path, destPath)
	})
}

func copyFile(src string, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

func (g *Generator) GetPortfolio() *Portfolio {
	return g.portfolio
}

