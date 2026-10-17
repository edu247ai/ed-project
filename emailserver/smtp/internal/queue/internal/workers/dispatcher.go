package workers

import (
	"context"
	"email-server/smtp/internal/queue/pkg"
	"errors"
	"strings"
)

type Dispatcher struct {
	sender  *Sender
	storage pkg.QueueStorage
}

func NewDispatcher(sender *Sender, storage pkg.QueueStorage) *Dispatcher {
	return &Dispatcher{sender: sender, storage: storage}
}

func (d *Dispatcher) Deliver(ctx context.Context, item pkg.QueueItem) error {
	raw, mailFrom, err := d.storage.LoadMessage(ctx, item.MessageID)
	if err != nil {
		return err
	}

	at := strings.LastIndexByte(item.Recipient, '@')
	if at < 1 || at == len(item.Recipient)-1 {
		return errors.New("invalid queued recipient address")
	}
	domain := item.Recipient[at+1:]
	mxs, err := ResolveMX(domain)
	if err != nil {
		return err
	}
	if len(mxs) == 0 {
		return errors.New("recipient domain has no MX records")
	}

	var lastErr error
	for _, mx := range mxs {
		if err := d.sender.Send(mx.Host, item.Recipient, raw, mailFrom); err == nil {
			return nil
		} else {
			lastErr = err
		}
	}
	return lastErr
}
