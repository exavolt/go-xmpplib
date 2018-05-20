// XEP-0203: Delayed Delivery
package xmppdelay

import (
	"encoding/xml"
	"time"

	"github.com/exavolt/go-xmpplib/xmppcore"
)

const (
	NS               = "urn:xmpp:delay"
	DelayElementName = NS + " delay"
)

type Delay struct {
	XMLName xml.Name      `xml:"urn:xmpp:delay delay"`
	From    *xmppcore.JID `xml:"from,attr,omitempty"`
	Stamp   time.Time     `xml:"stamp,attr"`
	Text    string        `xml:",chardata"`
}
