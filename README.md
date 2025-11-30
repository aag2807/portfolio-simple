# Portfolio Static Site Generator

A lightweight static site generator written in Go that builds a professional developer portfolio from JSON data.

## Features

- **Fast Generation**: Builds static HTML in milliseconds
- **Data-Driven**: All content loaded from `portfolio.json`
- **Dark/Light Mode**: Manual toggle with localStorage persistence
- **Responsive**: Mobile-first design with Tailwind CSS
- **Print-Friendly**: Optimized for PDF export
- **Zero Dependencies**: Pure Go standard library

## Quick Start

### Prerequisites

- Go 1.21 or later

### Build & Generate

```bash
# Build the generator
go build -o generator.exe ./cmd/generator

# Generate the static site
./generator.exe

# Or run directly
go run ./cmd/generator
```

### Output

The generated site will be in the `dist/` directory:

```
dist/
└── index.html
```

## Configuration

### Command Line Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-data` | `portfolio.json` | Path to portfolio JSON data |
| `-template` | `templates/index.html` | Path to HTML template |
| `-output` | `dist` | Output directory |

### Example

```bash
./generator.exe -data=my-data.json -output=public
```

## Editing Content

Edit `portfolio.json` to update your portfolio content. The structure includes:

- **personal**: Name, title, contact info
- **about**: Summary and highlights
- **polyglot**: Languages/tools grid
- **skills**: Categorized skills
- **experience**: Work history
- **projects**: Portfolio projects
- **certifications**: Professional certifications
- **education**: Educational background
- **languages**: Spoken languages
- **interests**: Personal interests

## Deployment to GitHub Pages

### Option 1: Manual Deployment

1. Generate the site:
   ```bash
   go run ./cmd/generator
   ```

2. Push the `dist/` folder contents to your `gh-pages` branch or repository root.

### Option 2: GitHub Actions

Create `.github/workflows/deploy.yml`:

```yaml
name: Deploy Portfolio

on:
  push:
    branches: [main]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      
      - name: Build
        run: go run ./cmd/generator
      
      - name: Deploy to GitHub Pages
        uses: peaceiris/actions-gh-pages@v3
        with:
          github_token: ${{ secrets.GITHUB_TOKEN }}
          publish_dir: ./dist
```

## Project Structure

```
portfolio/
├── cmd/
│   └── generator/
│       └── main.go           # Entry point
├── internal/
│   └── portfolio/
│       ├── data.go           # Data structures
│       └── generator.go      # Site generation
├── templates/
│   └── index.html            # HTML template
├── dist/                     # Generated output
├── portfolio.json            # Your data
├── go.mod
└── README.md
```

## Customization

### Styling

The template uses Tailwind CSS via CDN. Customize colors and fonts in the `<script>` block:

```javascript
tailwind.config = {
    theme: {
        extend: {
            colors: {
                ink: { /* your colors */ }
            }
        }
    }
}
```

### Template

Modify `templates/index.html` to change the layout. The template uses Go's `html/template` syntax.

## License

MIT

