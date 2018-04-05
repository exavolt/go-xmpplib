package xmppcore

import (
	"encoding/xml"
)

const StanzasNS = "urn:ietf:params:xml:ns:xmpp-stanzas"

// RFC 6120 § 8.3.2
const (
	StanzaErrorTypeAuth     = "auth"
	StanzaErrorTypeCancel   = "cancel"
	StanzaErrorTypeContinue = "continue"
	StanzaErrorTypeModify   = "modify"
	StanzaErrorTypeWait     = "wait"
)

// RFC 6120 § 8.3.2
//NOTE: the spec says that this element might contain
// application-specific error condition.
type StanzaError struct {
	XMLName   xml.Name `xml:"jabber:client error"`
	By        string   `xml:"by,attr,omitempty"`
	Type      string   `xml:"type,attr"`
	Condition StanzaErrorCondition
	Text      string `xml:"text,omitempty"`
}

type StanzaErrorCondition struct {
	XMLName xml.Name // Deliberately un-tagged
}

// RFC 6120 § 8.3.3
var (
	StanzaErrorConditionBadRequest            = stanzaErrorCondition("bad-request")
	StanzaErrorConditionConflict              = stanzaErrorCondition("conflict")
	StanzaErrorConditionFeatureNotImplemented = stanzaErrorCondition("feature-not-implemented")
	StanzaErrorConditionForbidden             = stanzaErrorCondition("forbidden")
	StanzaErrorConditionGone                  = stanzaErrorCondition("gone")
	StanzaErrorConditionInternalServerError   = stanzaErrorCondition("internal-server-error")
	StanzaErrorConditionItemNotFound          = stanzaErrorCondition("item-not-found")
	StanzaErrorConditionJIDMalformed          = stanzaErrorCondition("jid-malformed")
	StanzaErrorConditionNotAcceptable         = stanzaErrorCondition("not-acceptable")
	StanzaErrorConditionNotAllowed            = stanzaErrorCondition("not-allowed")
	StanzaErrorConditionNotAuthorized         = stanzaErrorCondition("not-authorized")
	StanzaErrorConditionPolicyViolation       = stanzaErrorCondition("policy-violation")
	StanzaErrorConditionRecipientUnavailable  = stanzaErrorCondition("recipient-unavailable")
	StanzaErrorConditionRedirect              = stanzaErrorCondition("redirect")
	StanzaErrorConditionRegistrationRequired  = stanzaErrorCondition("registration-required")
	StanzaErrorConditionRemoteServerNotFound  = stanzaErrorCondition("remote-server-not-found")
	StanzaErrorConditionRemoteServerTimeout   = stanzaErrorCondition("remote-server-timeout")
	StanzaErrorConditionResourceConstraint    = stanzaErrorCondition("resource-constraint")
	StanzaErrorConditionServiceUnavailable    = stanzaErrorCondition("service-unavailable")
	StanzaErrorConditionSubscriptionRequired  = stanzaErrorCondition("subscription-required")
	StanzaErrorConditionUndefinedCondition    = stanzaErrorCondition("undefined-condition")
	StanzaErrorConditionUnexpectedRequest     = stanzaErrorCondition("unexpected-request")
)

func stanzaErrorCondition(local string) StanzaErrorCondition {
	return StanzaErrorCondition{xml.Name{Space: StanzasNS, Local: local}}
}
