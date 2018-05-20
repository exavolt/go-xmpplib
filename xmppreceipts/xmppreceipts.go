// XEP-0184: Message Delivery Receipts
package xmppreceipts

import (
	"encoding/xml"
)

const (
	NS                  = "urn:xmpp:receipts"
	RequestElementName  = NS + " request"
	ReceivedElementName = NS + " received"
)

type Request struct {
	XMLName xml.Name `xml:"urn:xmpp:receipts request"`
}

type Received struct {
	XMLName xml.Name `xml:"urn:xmpp:receipts received"`
	ID      string   `xml:"id,attr,omitempty"`
}
