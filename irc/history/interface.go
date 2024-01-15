package history

import (
	"io"
	"time"
)

// HistoryInterfaceZCube is the history interface defined by ZCube in
// https://github.com/ZCube/ergo/commit/0a6925cc049a5a92981752cec2284c23f5de8792
type HistoryInterfaceZCube interface {
	// Open opens the history interface?
	Open() (err error)

	// AddChannelItem adds a non-message event to history
	AddChannelItem(target string, item Item, account string) (err error)

	// AddDirectMessage adds a message event to history
	AddDirectMessage(sender, senderAccount, recipient, recipientAccount string, item Item) (err error)

	// Close closes the database
	Close()

	// DeleteMsgid removes a message, specified by msgid and accountName
	DeleteMsgid(msgid, accountName string) (err error)

	// Export packages up all messages sent by account, and writes them as a JSON object to writer
	Export(account string, writer io.Writer)

	// Forget removes all messages sent by account
	Forget(account string)

	// ListChannels returns a list of all targets in the given list of case-folded channels
	ListChannels(cfchannels []string) (results []TargetListing, err error)

	// MakeSequence returns a Sequence of messages between correspondent and target
	// from cutoff to now
	MakeSequence(target, correspondent string, cutoff time.Time) Sequence
}
