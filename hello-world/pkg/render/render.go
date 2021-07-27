package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var tc map[string]*template.Template
	var err error

	tc, err = CreateTemplateCache()
	if err != nil {
		log.Fatalln("error render createtemplatecache:", err)
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatalln(err)
	}

	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("error byte buffer:", err)
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("../../templates/*.page.tmpl.html")
	if err != nil {
		log.Println("error render pages:", err)
		return myCache, nil
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			log.Println("error render parse pages:", err)
			return myCache, nil
		}
		matches, err := filepath.Glob("../../templates/*.layout.tmpl.html")
		if err != nil {
			log.Println("error render matches:", err)
			return myCache, nil
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../templates/*.layout.tmpl.html")
			if err != nil {
				log.Println("error render parse layout:", err)
				return myCache, nil
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
