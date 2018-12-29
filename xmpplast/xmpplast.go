// Package xmpplast contains support for XEP-0012
package xmpplast

import (
	"encoding/xml"
)

const (
	NS               = "jabber:iq:last"
	QueryElementName = NS + " query"
)

type IQQueryGet struct {
	XMLName xml.Name `xml:"jabber:iq:last query"`
}

type IQQueryResult struct {
	XMLName    xml.Name `xml:"jabber:iq:last query"`
	Seconds    int64    `xml:"seconds,attr"` // The spec says unsigned long (64bit) but Smack is expecting signed
	StatusText string   `xml:",chardata"`
}
