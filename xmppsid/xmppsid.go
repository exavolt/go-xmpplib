// XEP-0359: Unique and Stable Stanza IDs
package xmppsid

import (
	"encoding/xml"

	"github.com/rez-go/xmpplib/xmppcore"
)

const (
	NS                  = "urn:xmpp:sid:0"
	StanzaIDElementName = NS + " stanza-id"
	OriginIDElementName = NS + " origin-id"
)

type StanzaID struct {
	XMLName xml.Name      `xml:"urn:xmpp:sid:0 stanza-id"`
	ID      string        `xml:"id,attr"`
	By      *xmppcore.JID `xml:"by,attr"`
}

type OriginID struct {
	XMLName xml.Name `xml:"urn:xmpp:sid:0 origin-id"`
	ID      string   `xml:"id,attr"`
}
