package rss

import (
	"encoding/xml"
	"io"
	"time"
)

type RFC1123Time struct {
	time.Time
}

func (ct *RFC1123Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var content string
	if err := d.DecodeElement(&content, &start); err != nil {
		return err
	}
	t, err := time.Parse(time.RFC1123Z, content)
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
	Title         string      `xml:"title"`
	Description   string      `xml:"description"`
	Link          []string    `xml:"link"`
	Copyright     string      `xml:"copyright"`
	LastBuildDate RFC1123Time `xml:"lastBuildDate"`
	PubDate       RFC1123Time `xml:"pubDate"`
	TTL           int         `xml:"ttl"`
	Items         []Item      `xml:"item"`
	Language      string      `xml:"language"`

	//Additional  metadata
	Image      Image      `xml:"image"`
	Author     string     `xml:"author"`
	Categories []Category `xml:"category"`
	Owner      Owner      `xml:"owner"`
	Explicit   bool       `xml:"explicit"`
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
	Title       string      `xml:"title"`
	Description string      `xml:"description"`
	Link        string      `xml:"link"`
	GUUID       string      `xml:"guid"`
	PubDate     RFC1123Time `xml:"pubDate"`
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
