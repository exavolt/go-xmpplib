package xmppchatstates

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalActive(t *testing.T) {
	xmlBuf, err := xml.Marshal(&ChatStateActive)

	assert.Nil(t, err)
	assert.Equal(t,
		`<active xmlns="http://jabber.org/protocol/chatstates"></active>`,
		string(xmlBuf))
}

func TestMarshalComposing(t *testing.T) {
	xmlBuf, err := xml.Marshal(&ChatStateComposing)

	assert.Nil(t, err)
	assert.Equal(t,
		`<composing xmlns="http://jabber.org/protocol/chatstates"></composing>`,
		string(xmlBuf))
}

func TestMarshalPaused(t *testing.T) {
	xmlBuf, err := xml.Marshal(&ChatStatePaused)

	assert.Nil(t, err)
	assert.Equal(t,
		`<paused xmlns="http://jabber.org/protocol/chatstates"></paused>`,
		string(xmlBuf))
}

func TestMarshalInactive(t *testing.T) {
	xmlBuf, err := xml.Marshal(&ChatStateInactive)

	assert.Nil(t, err)
	assert.Equal(t,
		`<inactive xmlns="http://jabber.org/protocol/chatstates"></inactive>`,
		string(xmlBuf))
}

func TestMarshalGone(t *testing.T) {
	xmlBuf, err := xml.Marshal(&ChatStateGone)

	assert.Nil(t, err)
	assert.Equal(t,
		`<gone xmlns="http://jabber.org/protocol/chatstates"></gone>`,
		string(xmlBuf))
}
