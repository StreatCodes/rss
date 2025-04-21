package service

import (
	"log"
	"net/http"
	"net/url"

	"github.com/streatCodes/rss/rss"
)

type ChannelResult struct {
	ShowSubscribeButton bool
	ChannelUrl          string
	Channel             rss.Channel
}

func (service *Service) findChannel(query string) ([]ChannelResult, error) {
	//If the search query is a URL then ingest the feed
	if parsedURL, err := url.ParseRequestURI(query); err == nil {
		channelUrl := parsedURL.String()

		//Check to see if we have the feed in the database
		if channel, err := service.db.GetChannel(channelUrl); channel != nil && err == nil {
			return []ChannelResult{{ChannelUrl: channelUrl, Channel: *channel}}, nil
		}

		//Fetch from the internet
		log.Printf("Indexing channel %s\n", channelUrl)
		res, err := http.Get(channelUrl)
		if err != nil {
			return nil, err
		}

		feed, err := rss.Decode(res.Body)
		if err != nil {
			return nil, err
		}

		err = service.db.SaveChannel(channelUrl, &feed.Channel)
		if err != nil {
			return nil, err
		}
		return []ChannelResult{{ChannelUrl: channelUrl, Channel: feed.Channel}}, nil
	}

	return []ChannelResult{}, nil
}
