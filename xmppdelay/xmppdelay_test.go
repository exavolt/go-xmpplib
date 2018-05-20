package xmppdelay

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMarshalEmptyDelay(t *testing.T) {
	var delay Delay
	xmlBuf, err := xml.Marshal(&delay)

	assert.Nil(t, err)
	assert.Equal(t,
		`<delay xmlns="urn:xmpp:delay" stamp="0001-01-01T00:00:00Z"></delay>`,
		string(xmlBuf))
}

func TestMarshalBasicDelay(t *testing.T) {
	delay := Delay{Stamp: time.Unix(0, 0)}
	xmlBuf, err := xml.Marshal(&delay)

	assert.Nil(t, err)
	assert.Equal(t,
		`<delay xmlns="urn:xmpp:delay" stamp="1970-01-01T00:00:00Z"></delay>`,
		string(xmlBuf))
}
