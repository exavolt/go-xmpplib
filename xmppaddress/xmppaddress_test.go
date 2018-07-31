package xmppaddress

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalEmptyAddresses(t *testing.T) {
	var addresses Addresses
	xmlBuf, err := xml.Marshal(&addresses)

	assert.Nil(t, err)
	assert.Equal(t,
		`<addresses xmlns="http://jabber.org/protocol/address"></addresses>`,
		string(xmlBuf))
}

func TestMarshalSimpleTo(t *testing.T) {
	addresses := &Addresses{
		Address: []Address{
			{Type: AddressTypeTo},
		},
	}
	xmlBuf, err := xml.Marshal(&addresses)

	assert.Nil(t, err)
	assert.Equal(t,
		`<addresses xmlns="http://jabber.org/protocol/address"><address type="to"></address></addresses>`,
		string(xmlBuf))
}
