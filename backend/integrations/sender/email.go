package sender

import "context"

type Email struct {
	RecipientAddress string
	HTMLBody         string
	TextBody         string
}

type EmailSender interface {
	SendEmail(ctx context.Context, email *Email) error
}
