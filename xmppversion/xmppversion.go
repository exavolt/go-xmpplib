// XEP-0092: Software Version

package xmppversion

import (
	"encoding/xml"
)

const NS = "jabber:iq:version"

const QueryElementName = NS + " query"

type QueryResult struct {
	XMLName xml.Name `xml:"jabber:iq:version query"`
	Name    string   `xml:"name"`
	Version string   `xml:"version"`
	OS      string   `xml:"os,omitempty"`
}
