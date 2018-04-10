package xmppim

import (
	"encoding/xml"

	"github.com/exavolt/go-xmpplib/xmppcore"
)

const (
	ClientMessageElementName        = xmppcore.JabberClientNS + " message"
	ClientMessageBodyElementName    = xmppcore.JabberClientNS + " body"
	ClientMessageSubjectElementName = xmppcore.JabberClientNS + " subject"
	ClientMessageThreadElementName  = xmppcore.JabberClientNS + " thread"
)

// RFC 6121 section 5.2.2
const (
	MessageTypeChat      = "chat"
	MessageTypeError     = "error"
	MessageTypeGroupChat = "groupchat"
	MessageTypeHeadline  = "headline"
	MessageTypeNormal    = "normal"
)

// ClientMessageAttributes is the base structure for a message stanza. It defines
// all standard attributes.
type ClientMessageAttributes struct {
	XMLName xml.Name      `xml:"jabber:client message"`
	ID      string        `xml:"id,attr,omitempty"`
	From    *xmppcore.JID `xml:"from,attr,omitempty"`
	To      *xmppcore.JID `xml:"to,attr,omitempty"`
	Type    string        `xml:"type,attr"` // Any of MessageType*
}

// ClientMessage is a data structure for a message stanza. It contains
// all standard attributes and sub-elements.
type ClientMessage struct {
	ClientMessageAttributes
	Error   *xmppcore.StanzaError `xml:",omitempty"`
	Body    []ClientMessageBody    `xml:"body"`
	Subject []ClientMessageSubject `xml:"subject"`
	Thread  *ClientMessageThread   `xml:"thread"`
}

type ClientMessageBody struct {
	XMLName xml.Name `xml:"jabber:client body"`
	Lang    string   `xml:"lang,attr"`
	Text    string   `xml:",chardata"`
}

type ClientMessageSubject struct {
	XMLName xml.Name `xml:"jabber:client subject"`
	Lang    string   `xml:"lang,attr"`
	Text    string   `xml:",chardata"`
}

type ClientMessageThread struct {
	XMLName xml.Name `xml:"jabber:client thread"`
	Parent  string   `xml:"parent,attr"`
	Text    string   `xml:",chardata"`
}
