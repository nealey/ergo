package history

import (
	"log"
	"time"
)

// LogHistory provides a HistoryInterface that doesn't do anything other than log calls.
//
// This was created to test the interface mechanism.
type LogHistory struct {
}

// Close closes the database
func (hi *LogHistory) Close() error {
	return nil
}

// AddChannelItem adds a non-message event to history
func (hi *LogHistory) AddChannelItem(target string, item Item, account string) error {
	log.Println("AddChannelItem", target, item, account)
	return nil
}

// AddDirectMessage adds a message event to history
func (hi *LogHistory) AddDirectMessage(sender, senderAccount, recipient, recipientAccount string, item Item) error {
	log.Println("AddDirectMessage", sender, senderAccount, recipient, recipientAccount, item)
	return nil
}

// DeleteMessageId removes a message, specified by msgid and account
func (hi *LogHistory) DeleteMessage(account, messageId string) error {
	log.Println("DeleteMessage", messageId, account)
	return nil
}

// DeleteAllMessages removes all messages sent by account
func (hi *LogHistory) DeleteAllMessages(account string) error {
	log.Println("DeleteAllMessages", account)
	return nil
}

// ListChannels returns a list of all targets (users) in the given list of case-folded channels
func (hi *LogHistory) ListChannels(cfchannels []string) ([]TargetListing, error) {
	log.Println("ListChannels", cfchannels)
	return nil, nil
}

// MakeSequence returns a Sequence of messages between a and b, from cutoff to now
func (hi *LogHistory) MakeSequence(a, b string, cutoff time.Time) Sequence {
	return MemorySequence{}
}

type MemorySequence struct{}

func (s MemorySequence) Between(start, end Selector, limit int) ([]Item, error) {
	return nil, nil
}

func (s MemorySequence) Around(start Selector, limit int) ([]Item, error) {
	return nil, nil
}

func (s MemorySequence) ListCorrespondents(start, end Selector, limit int) ([]TargetListing, error) {
	return nil, nil
}

func (s MemorySequence) Cutoff() time.Time {
	return time.Time{}
}

func (s MemorySequence) Ephemeral() bool {
	return true
}
