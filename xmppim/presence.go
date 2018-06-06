package xmppim

import (
	"encoding/xml"

	"github.com/exavolt/go-xmpplib/xmppcore"
)

const ClientPresenceElementName = xmppcore.JabberClientNS + " presence"

// RFC 6121   2.2.1 and 4.7.1
const (
	PresenceTypeUnavailable  = "unavailable"
	PresenceTypeError        = "error"
	PresenceTypeProbe        = "probe"
	PresenceTypeSubscribe    = "subscribe"
	PresenceTypeSubscribed   = "subscribed"
	PresenceTypeUnsubscribe  = "unsubscribe"
	PresenceTypeUnsubscribed = "unsubscribed"
)

const (
	ShowValueAway = "away"
	ShowValueChat = "chat"
	ShowValueDND  = "dnd"
	ShowValueXA   = "xa"
)

// RFC 6121  4.7.
type ClientPresence struct {
	XMLName  xml.Name              `xml:"jabber:client presence"`
	ID       string                `xml:"id,attr,omitempty"`
	Type     string                `xml:"type,attr,omitempty"`
	From     string                `xml:"from,attr,omitempty"`
	To       string                `xml:"to,attr,omitempty"`
	Error    *xmppcore.StanzaError `xml:",omitempty"`
	Show     *ClientShow           `xml:"show,omitempty"`
	Status   []ClientStatus        `xml:"status,omitempty"`
	Priority *ClientPriority       `xml:"priority,omitempty"`
	//TODO: 4.7.2.
	CapsC *CapsC `xml:",omitempty"`
	//TODO: X
}

type ClientShow struct {
	XMLName xml.Name `xml:"jabber:client show"`
	Value   string   `xml:",chardata"`
}

type ClientStatus struct {
	XMLName xml.Name `xml:"jabber:client status"`
	Value   string   `xml:",chardata"`
	Lang    string   `xml:"lang,attr,omitempty"`
}

type ClientPriority struct {
	XMLName xml.Name `xml:"jabber:client priority"`
	Value   int32    `xml:",chardata"`
}
