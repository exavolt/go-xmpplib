package xmppreceipts

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalEmptyRequest(t *testing.T) {
	var request Request
	xmlBuf, err := xml.Marshal(&request)

	assert.Nil(t, err)
	assert.Equal(t,
		`<request xmlns="urn:xmpp:receipts"></request>`,
		string(xmlBuf))
}

func TestMarshalEmptyReceived(t *testing.T) {
	var received Received
	xmlBuf, err := xml.Marshal(&received)

	assert.Nil(t, err)
	assert.Equal(t,
		`<received xmlns="urn:xmpp:receipts"></received>`,
		string(xmlBuf))
}

func TestMarshalBasicReceived(t *testing.T) {
	received := Received{ID: "12345"}
	xmlBuf, err := xml.Marshal(&received)

	assert.Nil(t, err)
	assert.Equal(t,
		`<received xmlns="urn:xmpp:receipts" id="12345"></received>`,
		string(xmlBuf))
}
