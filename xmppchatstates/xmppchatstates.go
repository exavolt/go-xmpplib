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
	ChatStateComposing = chatState("composing")
	ChatStatePaused    = chatState("paused")
	ChatStateInactive  = chatState("inactive")
	ChatStateGone      = chatState("gone")
)

func chatState(local string) ChatState {
	return ChatState{xml.Name{Space: NS, Local: local}}
}

type Active struct {
	XMLName xml.Name `xml:"http://jabber.org/protocol/chatstates active"`
}

type Composing struct {
	XMLName xml.Name `xml:"http://jabber.org/protocol/chatstates composing"`
}

type Paused struct {
	XMLName xml.Name `xml:"http://jabber.org/protocol/chatstates paused"`
}

type Inactive struct {
	XMLName xml.Name `xml:"http://jabber.org/protocol/chatstates inactive"`
}

type Gone struct {
	XMLName xml.Name `xml:"http://jabber.org/protocol/chatstates gone"`
}
