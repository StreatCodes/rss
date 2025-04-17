package rss

import (
	"encoding/xml"
	"io"
	"time"
)

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
	t, err := time.Parse(time.RFC1123Z, aux.PubDate)
	if err != nil {
		return err
	}
	ch.PubDate = t

	t, err = time.Parse(time.RFC1123Z, aux.LastBuildDate)
	if err != nil {
		return err
	}
	ch.LastBuildDate = t
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
	t, err := time.Parse(time.RFC1123Z, aux.PubDate)
	if err != nil {
		return err
	}
	it.PubDate = t
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
