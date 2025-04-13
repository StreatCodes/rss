// Package opml implements decoding of OPML documents described at [opml.org].
// Only a subset of the format is implemented;
// enough to support the use of OPML as an interchange format by
// feed reader applications for web feed collections.
//
// [opml.org]: https://opml.org
package opml

import (
	"encoding/xml"
	"io"
)

type Document struct {
	XMLName struct{}  `xml:"opml"`
	Version string    `xml:"version,attr"`
	Body    []Outline `xml:"body>outline"`

	// head holds the document's metadata.
	// TODO(otl): should we even support this?
	// NetNewWire generates an empty tree.
	head struct{}
}

type Outline struct {
	XMLName     struct{}  `xml:"outline"`
	Text        string    `xml:"text,attr"`
	Title       string    `xml:"title,attr"`
	Description string    `xml:"description,attr,omitempty"`
	Type        string    `xml:"type,attr"`
	Version     string    `xml:"version,attr,omitempty"`
	HTML        string    `xml:"htmlUrl,attr"`
	XML         string    `xml:"xmlUrl,attr"`
	Children    []Outline `xml:"outline,omitempty"`
}

func Decode(r io.Reader) (*Document, error) {
	var doc Document
	if err := xml.NewDecoder(r).Decode(&doc); err != nil {
		return nil, err
	}
	return &doc, nil
}
