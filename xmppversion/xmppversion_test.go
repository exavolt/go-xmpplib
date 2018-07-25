package xmppversion

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalEmptyQueryResult(t *testing.T) {
	result := IQQueryResult{}
	xmlBuf, err := xml.Marshal(&result)

	assert.Nil(t, err)
	assert.Equal(t,
		`<query xmlns="jabber:iq:version"><name></name><version></version></query>`,
		string(xmlBuf))
}

func TestMarshalBasicQueryResult(t *testing.T) {
	result := IQQueryResult{Name: "XMPP Go", Version: "1.23.456"}
	xmlBuf, err := xml.Marshal(&result)

	assert.Nil(t, err)
	assert.Equal(t,
		`<query xmlns="jabber:iq:version"><name>XMPP Go</name><version>1.23.456</version></query>`,
		string(xmlBuf))
}
