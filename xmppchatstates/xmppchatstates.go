// XEP-0085: Chat State Notifications
package xmppchatstates

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
