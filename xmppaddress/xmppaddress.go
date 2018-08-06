// Package xmppaddress contains constants and structs
// to support XEP-0033 - Extended Stanza Addressing
package xmppaddress

import (
	"encoding/xml"

	"github.com/exavolt/go-xmpplib/xmppcore"
)

const (
	NS                   = "http://jabber.org/protocol/address"
	AddressesElementName = NS + " addresses"
)

type Addresses struct {
	XMLName xml.Name  `xml:"http://jabber.org/protocol/address addresses"`
	Address []Address `xml:"address,omitempty"`
}

type Address struct {
	XMLName   xml.Name      `xml:"address"` // no NS; assumed under Addresses
	JID       *xmppcore.JID `xml:"jid,attr,omitempty"`
	URI       string        `xml:"uri,attr,omitempty"`
	Node      string        `xml:"node,attr,omitempty"`
	Desc      string        `xml:"desc,attr,omitempty"`
	Delivered bool          `xml:"delivered,attr,omitempty"` // fixed to 'true'
	Type      string        `xml:"type,attr"`                // Any of AddressType*
}

const (
	AddressTypeTo        = "to"
	AddressTypeCC        = "cc"
	AddressTypeBCC       = "bcc"
	AddressTypeReplyTo   = "replyto"
	AddressTypeReplyRoom = "replyroom"
	AddressTypeNoReply   = "noreply"
	AddressTypeOFrom     = "ofrom"
)
