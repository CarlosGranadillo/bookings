package handlers

import (
	"net/http"

	"github.com/CarlosGranadillo/bookings/pkg/config"
	"github.com/CarlosGranadillo/bookings/pkg/models"
	"github.com/CarlosGranadillo/bookings/pkg/render"
)

// RepositoryType is the type (struct) to implement the repository pattern
type RepositoryType struct {
	AppConfig *config.ApplicationConfig
}

// NewRepo created a new repository (instance of RepositoryType) with the app config
func NewRepo(repoAppConfig *config.ApplicationConfig) *RepositoryType {
	return &RepositoryType{
		AppConfig: repoAppConfig,
	}
}

// Repo is the repository used by the handlers
var Repo *RepositoryType

// InitHandlersRepo sets the repository for the handlers
func InitHandlersRepo(r *RepositoryType) {
	Repo = r
}

// Home is the home page handler
func (receiver *RepositoryType) Home(w http.ResponseWriter, r *http.Request) {

	// playing with session -> storing the remote IP address in the session
	remoteIP := r.RemoteAddr
	receiver.AppConfig.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (receiver *RepositoryType) About(w http.ResponseWriter, r *http.Request) {
	// some data to pass to the template
	stringsForTemplate := make(map[string]string)
	stringsForTemplate["example"] = "This is a string that comes from the handler!!!!"

	// getting the remote IP address from the session
	// remoteIP will an empty string if there is no value in the session for the key "remote_ip"
	remoteIP := receiver.AppConfig.Session.GetString(r.Context(), "remote_ip")
	stringsForTemplate["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringsForTemplate,
	})

}
