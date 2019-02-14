package xmppcore

import (
	"encoding/xml"
)

//TODO: xml:lang

const ClientIQElementName = JabberClientNS + " iq"

// Standard IQ types
//
// RFC 6120  8.2.3
const (
	IQTypeGet    = "get"
	IQTypeSet    = "set"
	IQTypeResult = "result"
	IQTypeError  = "error"
)

// ClientIQ represents basic client IQ.
type ClientIQ struct {
	StanzaBaseAttributes
	XMLName xml.Name     `xml:"jabber:client iq"`
	Type    string       `xml:"type,attr"` // Any of IQType*
	Payload []byte       `xml:",innerxml"` //TODO:FIXME: this would contain all child elements
	Error   *StanzaError `xml:",omitempty"`
}

// HasFrom returns true if 'from' is provided and not empty.
func (clientIQ *ClientIQ) HasFrom() bool {
	return clientIQ != nil && clientIQ.From != nil && !clientIQ.From.IsEmpty()
}

// HasTo returns true if 'to' is provided and not empty.
func (clientIQ *ClientIQ) HasTo() bool {
	return clientIQ != nil && clientIQ.To != nil && !clientIQ.To.IsEmpty()
}
