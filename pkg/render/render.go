package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/CarlosGranadillo/bookings/pkg/config"
	"github.com/CarlosGranadillo/bookings/pkg/models"
)

var appConfigRender *config.ApplicationConfig

// InitAppConfig receives and stores the config for the render package
func InitAppConfig(appConfigPointer *config.ApplicationConfig) {
	appConfigRender = appConfigPointer
}

func addDefaultData(dataToInject *models.TemplateData) *models.TemplateData {
	// add default data here

	return dataToInject
}

func RenderTemplate(w http.ResponseWriter, templateName string, dataToInject *models.TemplateData) {

	var templateCache map[string]*template.Template
	var err error

	if appConfigRender.UseCache {
		// get cache from app config
		templateCache = appConfigRender.TemplateCache
	} else {
		// create the template cache
		templateCache, err = CreateTemplateCache()
		if err != nil {
			log.Fatal("error trying to create the template cache")
		}
	}

	// get requested template (from cache)
	parsedTemplate, isInCache := templateCache[templateName]
	if !isInCache {
		log.Fatal("template was not found in the cache")
	}

	dataToInject = addDefaultData(dataToInject)

	// render (execute) the template
	// err = parsedTemplate.Execute(w, nil)
	// use a buffer -> this extra step allows to see if the error comes from the value stores in the map
	buf := new(bytes.Buffer)
	err = parsedTemplate.Execute(buf, dataToInject)
	if err != nil {
		log.Println(err)
	}
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// createTemplateCache creates the entire cache at once -> Populate the map with everything that is available
func CreateTemplateCache() (map[string]*template.Template, error) {

	var templateCache = map[string]*template.Template{}

	// get list of paths (strings) of all the files that end with ".page.tmpl"
	pagePaths, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return templateCache, err
	}

	for _, pagePath := range pagePaths {

		// get filename from path
		pageFilename := filepath.Base(pagePath)

		// parsing file (current page)
		parsedTemplateSet, err := template.New(pageFilename).ParseFiles(pagePath)
		if err != nil {
			return templateCache, err
		}

		// check if there are layouts
		layoutPaths, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return templateCache, err
		}
		// if they are, add them to the current parsedTemplateSet being built in this iteration
		if len(layoutPaths) > 0 {
			parsedTemplateSet, err = parsedTemplateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return templateCache, err
			}
		}

		templateCache[pageFilename] = parsedTemplateSet
	}

	return templateCache, nil

}
