package xmppim

// XEP-0359: Unique and Stable Stanza IDs

const MessageSIDNamespace = "urn:xmpp:sid:0"

const (
	MessageSIDStanzaIDElementName = MessageSIDNamespace + " stanza-id"
	MessageSIDOriginIDElementName = MessageSIDNamespace + " origin-id"
)

type MessageSIDStanzaID struct {
	XMLName xml.Name      `xml:"urn:xmpp:sid:0 stanza-id"`
	ID      string        `xml:"id,attr"`
	By      *xmppcore.JID `xml:"by,attr"`
}

type MessageSIDOriginID struct {
	XMLName xml.Name `xml:"urn:xmpp:sid:0 origin-id"`
	ID      string   `xml:"id,attr"`
}
