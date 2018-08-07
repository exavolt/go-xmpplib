package xmppcore

import (
	"encoding/xml"
)

const (
	StreamStreamElementName   = JabberStreamsNS + " stream"
	StreamErrorElementName    = JabberStreamsNS + " error"
	StreamFeaturesElementName = JabberStreamsNS + " features"
)

// RFC 6120 § 4.3.2  Streams Features Format

type NegotiationStreamFeatures struct {
	XMLName    xml.Name        `xml:"stream:features"`
	Mechanisms *SASLMechanisms `xml:"mechanisms,omitempty"`
	//TODO: TLS
	//TODO: allow mods to provide more features
}

// AuthenticatedStreamFeatures is used in the second stream
type AuthenticatedStreamFeatures struct {
	XMLName xml.Name `xml:"stream:features"`
	Bind    BindBind `xml:"bind"`
	//TODO: get more features from the mods
}

// RFC 6120 § 4.9  Stream Errors

// StreamError represents a stream error (RFC 6120 § 4.9.2) element.
//
// To add application-specific condition element, embed this
// struct into application-specific stream error struct.
//
//TODO:FIXME: this will properly generate a valid XML, but
// this struct can't be used to unmarshal the generated
// XML. See: https://github.com/golang/go/issues/9519
type StreamError struct {
	XMLName   xml.Name `xml:"stream:error"`
	Condition StreamErrorCondition
	Text      string `xml:"text,omitempty"`
}

func (streamError StreamError) String() string {
	condition := streamError.Condition.XMLName.Local
	if condition == "" {
		condition = "undefined-condition"
	}
	if streamError.Text != "" {
		return condition + ": " + streamError.Text
	}
	return condition
}

// InboundStreamError is a struct used as the unmarshal
// receiver for inbound stream errors. We need this
// to be separated from StreamError because Go's
// stdlib XML codec doesn't fully support namespace
// well.
type InboundStreamError struct {
	XMLName   xml.Name `xml:"http://etherx.jabber.org/streams error"`
	Condition StreamErrorCondition
	Text      string `xml:"text,omitempty"`
}

func (streamError InboundStreamError) String() string {
	condition := streamError.Condition.XMLName.Local
	if condition == "" {
		condition = "undefined-condition"
	}
	if streamError.Text != "" {
		return condition + ": " + streamError.Text
	}
	return condition
}

// StreamErrorCondition holds condition information about a stream error.
//
// Per latest revision of RFC 6120, stream error conditions are empty elements.
type StreamErrorCondition struct {
	XMLName xml.Name // Deliberately un-tagged
}

// RFC 6120 § 4.9.3
var (
	StreamErrorConditionBadFormat           = streamErrorCondition("bad-format")
	StreamErrorConditionConnectionTimeout   = streamErrorCondition("connection-timeout")
	StreamErrorConditionHostUnknown         = streamErrorCondition("host-unknown")
	StreamErrorConditionInternalServerError = streamErrorCondition("internal-server-error")
	StreamErrorConditionInvalidFrom         = streamErrorCondition("invalid-from")
	StreamErrorConditionNotAuthorized       = streamErrorCondition("not-authorized")
	StreamErrorConditionSystemShutdown      = streamErrorCondition("system-shutdown")
	StreamErrorConditionUndefinedCondition  = streamErrorCondition("undefined-condition")
)

func streamErrorCondition(local string) StreamErrorCondition {
	// Without NS?
	return StreamErrorCondition{xml.Name{Space: StreamsNS, Local: local}}
}
