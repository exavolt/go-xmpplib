package xmppcore

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStreamErrorConditionBadFormatEncoding(t *testing.T) {
	def := StreamErrorConditionBadFormat
	xmlBuf, err := xml.Marshal(def)
	assert.Nil(t, err)
	assert.Equal(t,
		[]byte(`<bad-format xmlns="urn:ietf:params:xml:ns:xmpp-streams"></bad-format>`),
		xmlBuf)
}

func TestStreamErrorBasic(t *testing.T) {
	def := StreamError{
		Condition: StreamErrorConditionBadFormat,
	}
	xmlBuf, err := xml.Marshal(def)
	assert.Nil(t, err)
	assert.Equal(t,
		`<stream:error><bad-format xmlns="urn:ietf:params:xml:ns:xmpp-streams"></bad-format></stream:error>`,
		string(xmlBuf))
}

func TestBareStreamErrorUnmarshal(t *testing.T) {
	var streamError BareStreamError
	err := xml.Unmarshal([]byte(`<stream:error><bad-format xmlns="urn:ietf:params:xml:ns:xmpp-streams"></bad-format></stream:error>`), &streamError)
	assert.Nil(t, err)
	assert.Equal(t,
		BareStreamError{
			XMLName:   xml.Name{Space: "stream", Local: "error"},
			Condition: streamErrorCondition("bad-format"),
		}, streamError)
}
