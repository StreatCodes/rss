package rss

import (
	"encoding/xml"
	"io"
	"time"
)

type RFC822Time struct {
	time.Time
}

const rfc822Layout = "Mon, 2 Jan 2006 15:04:05 -0700"

func (ct *RFC822Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := d.DecodeElement(&content, &start); err != nil {
		return err
	}
	t, err := time.Parse(rfc822Layout, content)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

type RSS struct {
	Version string  `xml:"version,attr"`
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title         string     `xml:"title"`
	Description   string     `xml:"description"`
	Link          []string   `xml:"link"`
	Copyright     string     `xml:"copyright"`
	LastBuildDate RFC822Time `xml:"lastBuildDate"`
	PubDate       RFC822Time `xml:"pubDate"`
	TTL           int        `xml:"ttl"`
	Items         []Item     `xml:"item"`

	ITunesImage      string           `xml:"itunes:image"`
	ITunesAuthor     string           `xml:"itunes:author"`
	ITunesCategories []ItunesCategory `xml:"itunes:category"`
	ITunesOwner      []ItunesOwner    `xml:"itunes:owner"`
	ITunesExplicit   bool             `xml:"itunes:explicit"`
}

type AtomLink struct {
	Href string `xml:"href,attr"`
}

type ItunesCategory struct {
	Text string `xml:"text,attr"`
}

type ItunesOwner struct {
	Name  string `xml:"itunes:name"`
	Email string `xml:"itunes:email"`
}

type Item struct {
	Title       string     `xml:"title"`
	Description string     `xml:"description"`
	Link        string     `xml:"link"`
	GUUID       string     `xml:"guid"`
	PubDate     RFC822Time `xml:"pubDate"`
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
