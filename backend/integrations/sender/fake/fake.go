package fake

import (
	"context"

	"github.com/nikola-enter21/wms/backend/integrations/sender"
	"github.com/nikola-enter21/wms/backend/logging"
)

var (
	log = logging.MustNewLogger()
)

type fakeSender struct {
	from string
}

func NewFakeSender(from string) *fakeSender {
	return &fakeSender{
		from: from,
	}
}

func (f *fakeSender) SendEmail(ctx context.Context, email *sender.Email) error {
	log.Infof("Sending email from %s to %+v", f.from, email)
	return nil
}
