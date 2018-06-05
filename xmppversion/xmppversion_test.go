package xmppversion

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalEmptyQueryResult(t *testing.T) {
	result := QueryResult{}
	xmlBuf, err := xml.Marshal(&result)

	assert.Nil(t, err)
	assert.Equal(t,
		`<query xmlns="jabber:iq:version"><name></name><version></version></query>`,
		string(xmlBuf))
}

func TestMarshalBasicQueryResult(t *testing.T) {
	result := QueryResult{Name: "XMPP Go", Version: "0.9.99"}
	xmlBuf, err := xml.Marshal(&result)

	assert.Nil(t, err)
	assert.Equal(t,
		`<query xmlns="jabber:iq:version"><name>XMPP Go</name><version>0.9.99</version></query>`,
		string(xmlBuf))
}
