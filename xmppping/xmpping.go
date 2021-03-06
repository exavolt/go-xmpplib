// XEP-0199
package xmppping

import (
	"encoding/xml"
)

const (
	NS             = "urn:xmpp:ping"
	ElementName    = NS + " ping"
	EncodedElement = `<ping xmlns='urn:xmpp:ping'/>`
)

type IQGet struct {
	XMLName xml.Name `xml:"urn:xmpp:ping ping"`
}
