package history

import (
	"time"
)

// HistoryInterface is the history interface defined by ZCube in
// https://github.com/ZCube/ergo/commit/0a6925cc049a5a92981752cec2284c23f5de8792
type HistoryInterface interface {
	// Close closes the history provider
	Close() error

	// AddChannelItem adds a non-message event to history
	AddChannelItem(senderAccount, target string, item Item) error

	// AddDirectMessage adds a message event to history
	AddDirectMessage(sender, senderAccount, recipient, recipientAccount string, item Item) error

	// DeleteMessage removes a message, specified by message Id and account
	DeleteMessage(account, messageId string) error

	// DeleteAllMessages removes all messages sent by account
	DeleteAllMessages(account string) error

	// ListChannels returns a list of all targets in the given list of case-folded channels
	ListChannels(cfchannels []string) ([]TargetListing, error)

	// MakeSequence returns a Sequence of messages between correspondent and target
	// from cutoff to now
	MakeSequence(target, correspondent string, cutoff time.Time) Sequence
}

// Encoder defines an interface for encoding objects.
//
// This is implemented by several go builtins, including encoding/json.Encoder and encoding/gob.Encoder
type Encoder interface {
	Encode(e any) error
}
