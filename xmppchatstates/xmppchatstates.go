// XEP-0085: Chat State Notifications
package xmppchatstates

import (
	"encoding/xml"
)

const NS = "http://jabber.org/protocol/chatstates"

type ChatState struct {
	XMLName xml.Name // Deliberately un-tagged
}

var (
	ChatStateActive    = chatState("active")
	ChatStateInactive  = chatState("inactive")
	ChatStateGone      = chatState("gone")
	ChatStateComposing = chatState("composing")
	ChatStatePaused    = chatState("paused")
)

func chatState(local string) ChatState {
	return ChatState{xml.Name{Space: NS, Local: local}}
}

// Active is to indicate that the user is actively participating in the
// chat session.
type Active struct {
	XMLName xml.Name `xml:"http://jabber.org/protocol/chatstates active"`
}

// Inactive is to indicate that the user has not been actively
// participating in the chat session.
type Inactive struct {
	XMLName xml.Name `xml:"http://jabber.org/protocol/chatstates inactive"`
}

// Gone is to indicate that the user has effectively ended their
// participation in the chat session.
type Gone struct {
	XMLName xml.Name `xml:"http://jabber.org/protocol/chatstates gone"`
}

// Composing is to indicate that the user is composing a message.
type Composing struct {
	XMLName xml.Name `xml:"http://jabber.org/protocol/chatstates composing"`
}

// Paused is to indicate that the user had been composing but now has
// stopped.
type Paused struct {
	XMLName xml.Name `xml:"http://jabber.org/protocol/chatstates paused"`
}
