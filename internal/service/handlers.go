package service

import (
	"html/template"
	"net/http"
	"net/url"
	"strings"

	"github.com/streatCodes/rss/rss"
)

var templateFuncs = template.FuncMap{
	"timeAgo":    timeAgo,
	"pathEscape": url.PathEscape,
}

func render(w http.ResponseWriter, name string, data any) {
	tmpl := template.New("").Funcs(templateFuncs)
	tmpl = template.Must(tmpl.ParseGlob("internal/templates/*.tmpl"))
	tmpl.ExecuteTemplate(w, name, data)
}

func (service *Service) homeHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "home", nil)
}

func (service *Service) searchHandler(w http.ResponseWriter, r *http.Request) {
	isHtmx := r.Header.Get("HX-Request") == "true"
	searchQuery := r.URL.Query().Get("search")

	results, err := service.findChannel(searchQuery)
	if err != nil {
		panic("TODO")
	}

	if isHtmx {
		render(w, "results", results)
		return
	}
	render(w, "home", results)
}

type ChannelPage struct {
	ShowSubscribeButton bool
	Channel             *rss.Channel
}

func (service *Service) channelHandler(w http.ResponseWriter, r *http.Request) {
	channelPage := ChannelPage{ShowSubscribeButton: true}
	channelUrl := strings.TrimPrefix(r.URL.Path, "/channel/")

	//Check to see if we have the feed in the database
	if channel, err := service.db.GetChannel(channelUrl); channel != nil && err == nil {
		channelPage.Channel = channel
	}

	render(w, "channelPage", channelPage)
}
