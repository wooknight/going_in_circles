package handler

import (
	"net/http"

	"github.com/wooknight/going_in_circles/go/microservice/business"
	"github.com/wooknight/going_in_circles/go/microservice/config"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (rp *Repository) ListData(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	strMap["list"] = "List Data"
	remote_IP := rp.App.Session.GetString(r.Context(), "IP")
	strMap["remote_ip"] = remote_IP
	business.RenderTemplate(w, "list.page.tmpl", &config.TemplateData{
		StringMap: strMap,
	})
}

func (rp *Repository) AddData(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	rp.App.Session.Put(r.Context(), "IP", remoteIP)
	strMap := make(map[string]string)
	strMap["hello"] = "Hello world"
	business.RenderTemplate(w, "add.page.tmpl", &config.TemplateData{
		StringMap: strMap,
	})
}
