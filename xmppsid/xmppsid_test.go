package xmppsid

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rez-go/xmpplib/xmppcore"
)

func TestMarshalEmptyStanzaID(t *testing.T) {
	var stanzaID StanzaID
	assert.Panics(t, func() { xml.Marshal(&stanzaID) })
}

func TestMarshalBasicStanzaID(t *testing.T) {
	stanzaID := StanzaID{By: &xmppcore.JID{}}
	xmlBuf, err := xml.Marshal(&stanzaID)

	assert.Nil(t, err)
	assert.Equal(t,
		`<stanza-id xmlns="urn:xmpp:sid:0" id="" by=""></stanza-id>`,
		string(xmlBuf))
}
