package service

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
)

var templateFuncs = template.FuncMap{
	"timeAgo":    timeAgo,
	"pathEscape": url.PathEscape,
}

func render(w http.ResponseWriter, name string, data any) error {
	tmpl := template.New("").Funcs(templateFuncs)
	tmpl = template.Must(tmpl.ParseGlob("internal/templates/*.tmpl"))
	return tmpl.ExecuteTemplate(w, name, data)
}

func (service *Service) homeHandler(w http.ResponseWriter, r *http.Request) {
	err := render(w, "homePage", nil)
	if err != nil {
		log.Printf("Error executing template - %s", err)
	}
}

func (service *Service) searchHandler(w http.ResponseWriter, r *http.Request) {
	isHtmx := r.Header.Get("HX-Request") == "true"
	searchQuery := r.URL.Query().Get("search")

	results, err := service.findChannel(searchQuery)
	if err != nil {
		msg := fmt.Sprintf("find %q: %v", searchQuery, err)
		log.Println(msg)
		w.WriteHeader(http.StatusInternalServerError)
		render(w, "error", msg)
		return
	}

	if isHtmx {
		render(w, "results", results)
		return
	}
	err = render(w, "homePage", results)
	if err != nil {
		log.Printf("Error executing template - %s", err)
	}
}

func (service *Service) channelHandler(w http.ResponseWriter, r *http.Request) {
	channelPage := ChannelResult{ShowSubscribeButton: true}
	channelUrl := strings.TrimPrefix(r.URL.Path, "/channel/")

	// Check to see if we have the feed in the database
	channel, err := service.db.GetChannel(channelUrl)
	if err != nil {
		msg := fmt.Sprintf("get %q: %v", channelUrl, err)
		log.Println(msg)
		w.WriteHeader(http.StatusInternalServerError)
		render(w, "error", msg)
		return
	}
	channelPage.Channel = *channel
	err = render(w, "channelPage", channelPage)
	if err != nil {
		log.Printf("Error executing template - %s", err)
	}
}
