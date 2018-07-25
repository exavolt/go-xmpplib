// XEP-0092: Software Version

package xmppversion

import (
	"encoding/xml"
)

const (
	NS                    = "jabber:iq:version"
	IQQueryElementName    = NS + " query"
	IQQueryEncodedElement = `<query xmlns='jabber:iq:version'/>`
)

type IQQueryGet struct {
	XMLName xml.Name `xml:"jabber:iq:version query"`
}

type IQQueryResult struct {
	XMLName xml.Name `xml:"jabber:iq:version query"`
	Name    string   `xml:"name"`
	Version string   `xml:"version"`
	OS      string   `xml:"os,omitempty"`
}
