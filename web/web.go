package web

import (
	"bytes"
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/Olprog59/golog"
)

//go:embed static
var StaticFiles embed.FS

//go:embed all:templates all:partials
var templateFS embed.FS

// Template cache
var templates map[string]*template.Template

// Cache pour les partials
var partials map[string]*template.Template

// Initialiser les partials en même temps que les templates
func InitTemplates() error {
	templates = make(map[string]*template.Template)
	partials = make(map[string]*template.Template)

	// Charger le layout et les partials (pour les templates complets)
	baseTemplate, err := template.New("base").ParseFS(templateFS,
		"templates/layouts/base.html",
		"partials/*.html",
		"partials/components/*.html",
		"partials/forms/*.html",
	)
	if err != nil {
		return err
	}

	// Charger tous les partials individuellement pour pouvoir les utiliser seuls
	partialFiles, err := fs.Glob(templateFS, "partials/**/*.html")
	if err != nil {
		return err
	}

	for _, partialFile := range partialFiles {
		// Créer un nouveau template pour chaque partial
		partialName := strings.TrimSuffix(filepath.Base(partialFile), ".html")
		tmpl := template.New(partialName)

		// Parser le fichier
		tmpl, err = tmpl.ParseFS(templateFS, partialFile)
		if err != nil {
			return err
		}

		// Stocker dans le cache de partials
		partials[partialName] = tmpl
	}

	// Charger toutes les pages (comme avant)
	pages, err := fs.Glob(templateFS, "templates/pages/*.html")
	if err != nil {
		return err
	}

	for _, page := range pages {
		pageTmpl, err := baseTemplate.Clone()
		if err != nil {
			return err
		}

		_, err = pageTmpl.ParseFS(templateFS, page)
		if err != nil {
			return err
		}

		pageName := strings.TrimSuffix(filepath.Base(page), ".html")
		templates[pageName] = pageTmpl
	}

	return nil
}

// RenderTemplate avec buffer
func RenderTemplate(w http.ResponseWriter, name string, data any) {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		golog.Err("Template not found: %s", name)
		return
	}

	// Utiliser un buffer pour capturer le résultat avant de l'écrire
	buf := new(bytes.Buffer)
	err := tmpl.ExecuteTemplate(buf, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		golog.Err("Error rendering template %s: %v", name, err)
		return
	}

	// Écrire l'en-tête et le contenu uniquement en cas de succès
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	buf.WriteTo(w)
}

// RenderPartial avec buffer
func RenderPartial(w http.ResponseWriter, name string, data any) {
	tmpl, ok := partials[name]
	if !ok {
		http.Error(w, "Partial not found: "+name, http.StatusInternalServerError)
		golog.Err("Partial not found: %s", name)
		return
	}

	// Utiliser un buffer pour capturer le résultat avant de l'écrire
	buf := new(bytes.Buffer)
	err := tmpl.Execute(buf, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		golog.Err("Error rendering partial %s: %v", name, err)
		return
	}

	// Écrire l'en-tête et le contenu uniquement en cas de succès
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	buf.WriteTo(w)
}
