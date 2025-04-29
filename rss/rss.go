package rss

import (
	"encoding/xml"
	"fmt"
	"io"
	"time"
)

// MediaType is RSS' MIME media type.
const MediaType = "application/rss+xml"

type RSS struct {
	Version string  `xml:"version,attr"`
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title         string    `xml:"title"`
	Description   string    `xml:"description"`
	Link          []string  `xml:"link"`
	Copyright     string    `xml:"copyright"`
	LastBuildDate time.Time `xml:"lastBuildDate"`
	PubDate       time.Time `xml:"pubDate"`
	TTL           int       `xml:"ttl"`
	Items         []Item    `xml:"item"`
	Language      string    `xml:"language"`

	//Additional  metadata
	Image      Image      `xml:"image"`
	Author     string     `xml:"author"`
	Categories []Category `xml:"category"`
	Owner      Owner      `xml:"owner"`
	Explicit   bool       `xml:"explicit"`
}

func (ch *Channel) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type alias Channel
	aux := &struct {
		LastBuildDate string `xml:"lastBuildDate"`
		PubDate       string `xml:"pubDate"`
		*alias
	}{
		alias: (*alias)(ch),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}

	if aux.PubDate != "" {
		t, err := parseTime(aux.PubDate)
		if err != nil {
			return fmt.Errorf("parse published date %q: %w", aux.PubDate, err)
		}
		ch.PubDate = t
	}
	if aux.LastBuildDate != "" {
		t, err := parseTime(aux.LastBuildDate)
		if err != nil {
			return fmt.Errorf("parse last build date %q: %w", aux.LastBuildDate, err)
		}
		ch.LastBuildDate = t
	}

	return nil
}

type AtomLink struct {
	Href string `xml:"href,attr"`
}

type Image struct {
	Href string `xml:"href,attr"`
}

type Category struct {
	Text string `xml:"text,attr"`
}

type Owner struct {
	Name  string `xml:"name"`
	Email string `xml:"email"`
}

type Item struct {
	Title       string    `xml:"title"`
	Description string    `xml:"description"`
	Link        string    `xml:"link"`
	GUUID       string    `xml:"guid"`
	PubDate     time.Time `xml:"pubDate"`
}

func (it *Item) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type alias Item
	aux := &struct {
		PubDate string `xml:"pubDate"`
		*alias
	}{
		alias: (*alias)(it),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	if aux.PubDate != "" {
		t, err := parseTime(aux.PubDate)
		if err != nil {
			return fmt.Errorf("parse published date %q: %w", aux.PubDate, err)
		}
		it.PubDate = t
	}
	return nil
}

func Marshal(rss *RSS) ([]byte, error) {
	return xml.MarshalIndent(rss, "", "\t")
}

func Decode(r io.Reader) (*RSS, error) {
	var rss RSS
	if err := xml.NewDecoder(r).Decode(&rss); err != nil {
		return nil, err
	}
	return &rss, nil
}

func parseTime(s string) (time.Time, error) {
	layouts := []string{
		time.RFC1123Z, time.RFC1123,
		time.RFC822Z, time.RFC822,
		"Mon, _2 Jan 2006 15:04:05 -0700",     // rfc1123z with no trailing zero
		"Mon, _2 January 2006 15:04:05 -0700", // long month name
	}
	for _, l := range layouts {
		t, err := time.Parse(l, s)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unsupported layout")
}
