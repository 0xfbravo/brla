package interfaces

import (
	"github.com/0xfbravo/brla/model"
)

type WebhookHandler interface {
	// HandleQueued handles the `queued` webhook response.
	HandleQueued(req *model.WebhookResponse) error

	// HandlePosted handles the posted webhook response
	HandlePosted(req *model.WebhookResponse) error

	// HandleSuccess handles the success webhook response
	HandleSuccess(req *model.WebhookResponse) error

	// HandleFailed handles the failed webhook response
	HandleFailed(req *model.WebhookResponse) error
}
