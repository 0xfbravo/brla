package interfaces

import (
	"context"
	"github.com/0xfbravo/brla/model"
)

type WebhookHandler interface {
	// HandleQueued handles the `queued` webhook response.
	HandleQueued(c context.Context, req *model.WebhookResponse) error

	// HandlePosted handles the posted webhook response
	HandlePosted(c context.Context, req *model.WebhookResponse) error

	// HandleSuccess handles the success webhook response
	HandleSuccess(c context.Context, req *model.WebhookResponse) error

	// HandleFailed handles the failed webhook response
	HandleFailed(c context.Context, req *model.WebhookResponse) error
}
