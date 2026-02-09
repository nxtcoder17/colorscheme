package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

// PaletteFile represents the YAML file structure
type PaletteFile struct {
	Sets map[string]Palette `yaml:",inline"`
}

// Palette represents a color palette (dark or light)
type Palette struct {
	BG     map[string]string `yaml:"bg"`
	FG     map[string]string `yaml:"fg"`
	Colors map[string]string `yaml:"colors"`
}

// Flavor represents a color scheme flavor (light/dark)
type Flavor struct {
	Name       string
	Identifier string
	IsDark     bool
	Palette    Palette
}

// TemplateData holds all data passed to templates
type TemplateData struct {
	Flavor Flavor
	Mode   string
	Env    map[string]string
	// Flattened color access: "bg.default", "fg.muted", "colors.red"
	Colors map[string]string
}

func main() {
	var (
		flavorName   = flag.String("flavor", "", "Flavor to use: dark or light")
		templatePath = flag.String("template", "", "Path to template file (.gotmpl)")
		outputPath   = flag.String("output", "", "Output file path (optional, prints to stdout if not set)")
		colorsFile   = flag.String("colors", "palette.yml", "Path to palette YAML file")
		mode         = flag.String("mode", "", "Mode: light or dark (defaults based on flavor)")
	)
	flag.Parse()

	if *flavorName == "" || *templatePath == "" {
		fmt.Fprintln(os.Stderr, "Usage: colorscheme -flavor <dark|light> -template <path> [-output <path>] [-colors <path>]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Load colors configuration
	paletteFile, err := loadPalette(*colorsFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading palette: %v\n", err)
		os.Exit(1)
	}

	// Select flavor
	flavor, err := selectFlavor(paletteFile, *flavorName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Determine mode
	modeVal := *mode
	if modeVal == "" {
		modeVal = flavor.Identifier
	}

	// Build template data
	data := TemplateData{
		Flavor: flavor,
		Mode:   modeVal,
		Env:    loadEnv(),
		Colors: flattenPalette(flavor.Palette),
	}

	// Process template
	output, err := processTemplate(*templatePath, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing template: %v\n", err)
		os.Exit(1)
	}

	// Write output
	if *outputPath != "" {
		// Ensure output directory exists
		if err := os.MkdirAll(filepath.Dir(*outputPath), 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
			os.Exit(1)
		}
		if err := os.WriteFile(*outputPath, []byte(output), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing output: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "Generated: %s\n", *outputPath)
	} else {
		fmt.Print(output)
	}
}

func loadPalette(path string) (*PaletteFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading palette file: %w", err)
	}

	var paletteFile PaletteFile
	if err := yaml.Unmarshal(data, &paletteFile); err != nil {
		return nil, fmt.Errorf("parsing palette YAML: %w", err)
	}

	return &paletteFile, nil
}

func selectFlavor(pf *PaletteFile, name string) (Flavor, error) {
	name = strings.ToLower(name)

	palette, ok := pf.Sets[name]
	if !ok {
		available := make([]string, 0, len(pf.Sets))
		for k := range pf.Sets {
			available = append(available, k)
		}
		return Flavor{}, fmt.Errorf("unknown flavor: %s (available: %s)", name, strings.Join(available, ", "))
	}

	return Flavor{
		Name:       strings.Title(name),
		Identifier: name,
		IsDark:     name == "dark",
		Palette:    palette,
	}, nil
}

func loadEnv() map[string]string {
	env := make(map[string]string)
	for _, e := range os.Environ() {
		if i := strings.Index(e, "="); i >= 0 {
			env[e[:i]] = e[i+1:]
		}
	}
	return env
}

// flattenPalette converts nested palette to flat map with dotted keys
// e.g., palette.BG["default"] -> "bg.default"
func flattenPalette(p Palette) map[string]string {
	colors := make(map[string]string)

	for name, hex := range p.BG {
		colors["bg."+name] = strings.TrimPrefix(hex, "#")
	}
	for name, hex := range p.FG {
		colors["fg."+name] = strings.TrimPrefix(hex, "#")
	}
	for name, hex := range p.Colors {
		colors["colors."+name] = strings.TrimPrefix(hex, "#")
	}

	return colors
}

func processTemplate(templatePath string, data TemplateData) (string, error) {
	funcMap := template.FuncMap{
		// String manipulation
		"lower":      strings.ToLower,
		"upper":      strings.ToUpper,
		"capitalize": strings.Title,
		"trimPrefix": strings.TrimPrefix,
		"trimSuffix": strings.TrimSuffix,
		"replace":    strings.ReplaceAll,

		// Color helpers - direct access by category
		"bg": func(name string) string {
			if v, ok := data.Flavor.Palette.BG[name]; ok {
				return strings.TrimPrefix(v, "#")
			}
			return ""
		},
		"fg": func(name string) string {
			if v, ok := data.Flavor.Palette.FG[name]; ok {
				return strings.TrimPrefix(v, "#")
			}
			return ""
		},
		"colors": func(name string) string {
			if v, ok := data.Flavor.Palette.Colors[name]; ok {
				return strings.TrimPrefix(v, "#")
			}
			return ""
		},

		// Environment access
		"env": func(name string) string {
			return os.Getenv(name)
		},

		// Conditional helpers
		"isDark": func() bool {
			return data.Flavor.IsDark
		},
		"isLight": func() bool {
			return !data.Flavor.IsDark
		},
		"ternary": func(cond bool, ifTrue, ifFalse string) string {
			if cond {
				return ifTrue
			}
			return ifFalse
		},
	}

	tmpl, err := template.New(filepath.Base(templatePath)).Funcs(funcMap).ParseFiles(templatePath)
	if err != nil {
		return "", fmt.Errorf("parsing template: %w", err)
	}

	var buf strings.Builder
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("executing template: %w", err)
	}

	return buf.String(), nil
}
