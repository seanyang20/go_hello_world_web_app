package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/tsawler/go-course/pkg/config"
	"github.com/tsawler/go-course/pkg/models"
)

// ---------
// Creating a simple template cache
// ----------

// func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
// 	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("error parsing template:", err)
// 		return
// 	}
// }

// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	// check to see if we already have the template in our cache
// 	_, inMap := tc[t]
// 	if !inMap {
// 		// need to create the template
// 		log.Println("creating template and adding to cache")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		// we have the template in the cache
// 		log.Println("using cached template")
// 	}

// 	tmpl = tc[t]

// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.html",
// 	}

// 	// parse the template
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	// add template to cache
// 	tc[t] = tmpl

// 	return nil
// }

// ---------
// Creating a more complex template cache
// ----------

// func RenderTemplate(w http.ResponseWriter, tmpl string) {
// 	// create a template cache
// 	tc, err := CreateTemplateCache()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// get requested template from cache
// 	t, ok := tc[tmpl]
// 	if !ok {
// 		log.Fatal(err)
// 	}

// 	buf := new(bytes.Buffer) // this buffer will hold bytes

// 	err = t.Execute(buf, nil) // trying to execute the value got from map with buffer instead of directly then write it out
// 	if err != nil {           // using buffer for finer grain error checking
// 		log.Println(err)
// 	}

// 	// render the template
// 	_, err = buf.WriteTo(w)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func CreateTemplateCache() (map[string]*template.Template, error) {
// 	// myCache := make(map[string]*template.Template)
// 	myCache := map[string]*template.Template{} // same as immediate above

// 	// get all of the files named *.page.tmpl from ./templates
// 	pages, err := filepath.Glob("./templates/*.page.html")
// 	if err != nil {
// 		return myCache, err
// 	}

// 	// range through all files ending with *.page.html
// 	for _, page := range pages {
// 		name := filepath.Base(page) // getting the name of the file minus the full path (eg. just 'home.page.html'")
// 		// I populate the variable ts (template) while checking for an error
// 		// when I parse the file named page and store in a template called name
// 		ts, err := template.New(name).ParseFiles(page)
// 		if err != nil {
// 			return myCache, err
// 		}

// 		matches, err := filepath.Glob("./templates/*.layout.html") // looking my layouts
// 		if err != nil {
// 			return myCache, err
// 		}

// 		if len(matches) > 0 {
// 			// if a layout template exists I populate my ts (template)
// 			ts, err = ts.ParseGlob("./templates/*.layout.html")
// 			if err != nil {
// 				return myCache, err
// 			}
// 		}

// 		myCache[name] = ts
// 	}

// 	return myCache, nil
// }

// ---------
// Optimizing our template cache
// ----------

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// // create a template cache
	// tc, err := CreateTemplateCache()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer) // this buffer will hold bytes

	td = AddDefaultData(td)

	_ = t.Execute(buf, td) // trying to execute the value got from map with buffer instead of directly then write it out
	// using buffer for finer grain error checking

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{} // same as immediate above

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.html
	for _, page := range pages {
		name := filepath.Base(page) // getting the name of the file minus the full path (eg. just 'home.page.html'")
		// I populate the variable ts (template) while checking for an error
		// when I parse the file named page and store in a template called name
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html") // looking my layouts
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			// if a layout template exists I populate my ts (template)
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
