// XEP-0198: Stream Management

package xmppsm

import (
	"encoding/xml"
)

const NS = "urn:xmpp:sm:3"

const (
	SMElementName      = NS + " sm"
	AElementName       = NS + " a"
	EnableElementName  = NS + " enable"
	EnabledElementName = NS + " enabled"
	FailedElementName  = NS + " failed"
	RElementName       = NS + " r"
	ResumeElementName  = NS + " resume"
	ResumedElementName = NS + " resumed"
)

type SM struct {
	XMLName xml.Name `xml:"urn:xmpp:sm:3 sm"`
}

type A struct {
	XMLName xml.Name `xml:"urn:xmpp:sm:3 a"`
	H       int64    `xml:"h,attr"`
}

type Enable struct {
	XMLName xml.Name `xml:"urn:xmpp:sm:3 enable"`
	Max     int64    `xml:"max,attr,omitempty"`
	Resume  bool     `xml:"resume,attr,omitempty"`
}

type Enabled struct {
	XMLName  xml.Name `xml:"urn:xmpp:sm:3 enabled"`
	ID       string   `xml:"id,attr,omitempty"`
	Location string   `xml:"location,attr,omitempty"`
	Max      int64    `xml:"max,attr,omitempty"`
	Resume   bool     `xml:"resume,attr,omitempty"`
}

type Failed struct {
	XMLName xml.Name `xml:"urn:xmpp:sm:3 failed"`
	H       int64    `xml:"h,attr,omitempty"`
	//TODO: Stream errors
}

type R struct {
	XMLName xml.Name `xml:"urn:xmpp:sm:3 r"`
}

type Resumption struct {
	H      uint32 `xml:"h,attr"`
	PrevID string `xml:"previd,attr"`
}

type Resume struct {
	XMLName xml.Name `xml:"urn:xmpp:sm:3 resume"`
	Resumption
}

type Resumed struct {
	XMLName xml.Name `xml:"urn:xmpp:sm:3 resumed"`
	Resumption
}
