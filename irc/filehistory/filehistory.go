package filehistory

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ergochat/ergo/irc/history"
)

type FileHistory struct {
	Path string
}

// NewFileHistory checks the validity of path, and returns a new FileHistory
func NewFileHistory(path string) (*FileHistory, error) {
	if info, err := os.Stat(path); err != nil {
		return nil, err
	} else if !info.IsDir() {
		return nil, fmt.Errorf("is not a directory: %s", err.Error())
	}
	hi := FileHistory{
		Path: path,
	}
	return &hi, nil
}

// Close does nothing, but is required to satisfy history.HistoryInterface
func (d *FileHistory) Close() error {
	return nil
}

// AddChannelItem adds a non-message event to history
func (hi *FileHistory) AddChannelItem(senderAccount, target string, item history.Item) error {
	log.Println("AddChannelItem", senderAccount, target, item)
	return nil
}

// AddDirectMessage adds a message event to history
func (hi *FileHistory) AddDirectMessage(sender, senderAccount, recipient, recipientAccount string, item history.Item) error {
	log.Println("AddDirectMessage", sender, senderAccount, recipient, recipientAccount, item)
	return nil
}

// DeleteMessageId removes a message, specified by msgid and account
func (hi *FileHistory) DeleteMessage(account, messageId string) error {
	log.Println("DeleteMessage", messageId, account)
	return nil
}

// DeleteAllMessages removes all messages sent by account
func (hi *FileHistory) DeleteAllMessages(account string) error {
	log.Println("DeleteAllMessages", account)
	return nil
}

// ListChannels returns a list of all targets (users) in the given list of case-folded channels
func (hi *FileHistory) ListChannels(cfchannels []string) ([]history.TargetListing, error) {
	log.Println("ListChannels", cfchannels)
	return nil, nil
}

// MakeSequence returns a Sequence of messages between a and b, from cutoff to now
func (hi *FileHistory) MakeSequence(a, b string, cutoff time.Time) history.Sequence {
	return FileHistorySequence{}
}

type FileHistorySequence struct{}

func (s FileHistorySequence) Between(start, end history.Selector, limit int) ([]history.Item, error) {
	return nil, nil
}

func (s FileHistorySequence) Around(start history.Selector, limit int) ([]history.Item, error) {
	return nil, nil
}

func (s FileHistorySequence) ListCorrespondents(start, end history.Selector, limit int) ([]history.TargetListing, error) {
	return nil, nil
}

func (s FileHistorySequence) Cutoff() time.Time {
	return time.Time{}
}

func (s FileHistorySequence) Ephemeral() bool {
	return true
}
