package service

import (
	"net/http"
	"net/url"

	"github.com/streatCodes/rss/rss"
)

func (service *Service) findChannel(query string) ([]rss.Channel, error) {
	//If the search query is a URL then ingest the feed
	if parsedURL, err := url.ParseRequestURI(query); err == nil {
		//Check to see if we have the feed in the database
		if channel, err := service.db.GetChannel(parsedURL.String()); channel != nil && err == nil {
			return []rss.Channel{*channel}, nil
		}

		//Fetch from the internet
		res, err := http.Get(query)
		if err != nil {
			return nil, err
		}

		feed, err := rss.Decode(res.Body)
		if err != nil {
			return nil, err
		}

		err = service.db.SaveChannel(parsedURL.String(), &feed.Channel)
		if err != nil {
			return nil, err
		}
		return []rss.Channel{feed.Channel}, nil
	}

	return []rss.Channel{}, nil
}
