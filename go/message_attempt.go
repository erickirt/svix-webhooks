package svix

import (
	"context"
	"time"

	"github.com/svix/svix-webhooks/go/internal/openapi"
)

type MessageAttempt struct {
	api *openapi.APIClient
}

type (
	MessageStatus                         = openapi.MessageStatus
	StatusCodeClass                       = openapi.StatusCodeClass
	ListResponseMessageAttemptOut         = openapi.ListResponseMessageAttemptOut
	MessageAttemptOut                     = openapi.MessageAttemptOut
	ListResponseEndpointMessageOut        = openapi.ListResponseEndpointMessageOut
	EndpointMessageOut                    = openapi.EndpointMessageOut
	ListResponseMessageEndpointOut        = openapi.ListResponseMessageEndpointOut
	MessageEndpointOut                    = openapi.MessageEndpointOut
	ListResponseMessageAttemptEndpointOut = openapi.ListResponseMessageAttemptEndpointOut
	MessageAttemptEndpointOut             = openapi.MessageAttemptEndpointOut
)

type MessageAttemptListOptions struct {
	Iterator        *string
	Limit           *int32
	Status          *MessageStatus
	EventTypes      *[]string
	Before          *time.Time
	After           *time.Time
	StatusCodeClass *StatusCodeClass
	Channel         *string
	EndpointId      *string
	WithContent     *bool
	WithMsg         *bool
	Tag             *string
}

// Deprecated: use `ListByMsg` or `ListByEndpoint` instead
func (m *MessageAttempt) List(ctx context.Context, appId string, msgId string, options *MessageAttemptListOptions) (*ListResponseMessageAttemptOut, error) {
	return m.ListByMsg(ctx, appId, msgId, options)
}

func (m *MessageAttempt) ListByMsg(ctx context.Context, appId string, msgId string, options *MessageAttemptListOptions) (*ListResponseMessageAttemptOut, error) {
	req := m.api.MessageAttemptAPI.V1MessageAttemptListByMsg(ctx, appId, msgId)
	if options != nil {
		if options.Iterator != nil {
			req = req.Iterator(*options.Iterator)
		}
		if options.Limit != nil {
			req = req.Limit(*options.Limit)
		}
		if options.Status != nil {
			req = req.Status(*options.Status)
		}
		if options.EventTypes != nil {
			req = req.EventTypes(*options.EventTypes)
		}
		if options.Before != nil {
			req = req.Before(*options.Before)
		}
		if options.After != nil {
			req = req.After(*options.After)
		}
		if options.StatusCodeClass != nil {
			req.StatusCodeClass(*options.StatusCodeClass)
		}
		if options.Channel != nil {
			req = req.Channel(*options.Channel)
		}
		if options.EndpointId != nil {
			req = req.EndpointId(*options.EndpointId)
		}
		if options.WithContent != nil {
			req = req.WithContent(*options.WithContent)
		}
		if options.Tag != nil {
			req = req.Tag(*options.Tag)
		}
	}
	ret, res, err := req.Execute()
	if err != nil {
		return nil, wrapError(err, res)
	}
	return ret, nil
}

func (m *MessageAttempt) ListByEndpoint(ctx context.Context, appId string, endpointId string, options *MessageAttemptListOptions) (*ListResponseMessageAttemptOut, error) {
	req := m.api.MessageAttemptAPI.V1MessageAttemptListByEndpoint(ctx, appId, endpointId)
	if options != nil {
		if options.Iterator != nil {
			req = req.Iterator(*options.Iterator)
		}
		if options.Limit != nil {
			req = req.Limit(*options.Limit)
		}
		if options.Status != nil {
			req = req.Status(*options.Status)
		}
		if options.EventTypes != nil {
			req = req.EventTypes(*options.EventTypes)
		}
		if options.Before != nil {
			req = req.Before(*options.Before)
		}
		if options.After != nil {
			req = req.After(*options.After)
		}
		if options.StatusCodeClass != nil {
			req.StatusCodeClass(*options.StatusCodeClass)
		}
		if options.Channel != nil {
			req = req.Channel(*options.Channel)
		}
		if options.WithContent != nil {
			req = req.WithContent(*options.WithContent)
		}
		if options.WithMsg != nil {
			req = req.WithMsg(*options.WithMsg)
		}
		if options.Tag != nil {
			req = req.Tag(*options.Tag)
		}
	}
	ret, res, err := req.Execute()
	if err != nil {
		return nil, wrapError(err, res)
	}
	return ret, nil
}

func (m *MessageAttempt) Get(ctx context.Context, appId string, msgId string, attemptID string) (*MessageAttemptOut, error) {
	req := m.api.MessageAttemptAPI.V1MessageAttemptGet(ctx, appId, msgId, attemptID)
	ret, res, err := req.Execute()
	if err != nil {
		return nil, wrapError(err, res)
	}
	return ret, nil
}

func (m *MessageAttempt) Resend(ctx context.Context, appId string, msgId string, endpointId string) error {
	return m.ResendWithOptions(ctx, appId, msgId, endpointId, nil)
}

func (m *MessageAttempt) ResendWithOptions(ctx context.Context, appId string, msgId string, endpointId string, options *PostOptions) error {
	req := m.api.MessageAttemptAPI.V1MessageAttemptResend(ctx, appId, msgId, endpointId)
	if options != nil {
		if options.IdempotencyKey != nil {
			req = req.IdempotencyKey(*options.IdempotencyKey)
		}
	}
	res, err := req.Execute()
	return wrapError(err, res)
}

func (m *MessageAttempt) ListAttemptedMessages(ctx context.Context, appId string, endpointId string, options *MessageAttemptListOptions) (*ListResponseEndpointMessageOut, error) {
	req := m.api.MessageAttemptAPI.V1MessageAttemptListAttemptedMessages(ctx, appId, endpointId)
	if options != nil {
		if options.Iterator != nil {
			req = req.Iterator(*options.Iterator)
		}
		if options.Limit != nil {
			req = req.Limit(*options.Limit)
		}
		if options.Status != nil {
			req = req.Status(openapi.MessageStatus(*options.Status))
		}
		if options.Before != nil {
			req = req.Before(*options.Before)
		}
		if options.After != nil {
			req = req.After(*options.After)
		}
		if options.Channel != nil {
			req = req.Channel(*options.Channel)
		}
		if options.WithContent != nil {
			req = req.WithContent(*options.WithContent)
		}
		if options.EventTypes != nil {
			req = req.EventTypes(*options.EventTypes)
		}
		if options.Tag != nil {
			req = req.Tag(*options.Tag)
		}
	}
	ret, res, err := req.Execute()
	if err != nil {
		return nil, wrapError(err, res)
	}
	return ret, nil
}

func (m *MessageAttempt) ListAttemptedDestinations(ctx context.Context, appId string, msgId string, options *MessageAttemptListOptions) (*ListResponseMessageEndpointOut, error) {
	req := m.api.MessageAttemptAPI.V1MessageAttemptListAttemptedDestinations(ctx, appId, msgId)
	if options != nil {
		if options.Iterator != nil {
			req = req.Iterator(*options.Iterator)
		}
		if options.Limit != nil {
			req = req.Limit(*options.Limit)
		}
	}
	ret, res, err := req.Execute()
	if err != nil {
		return nil, wrapError(err, res)
	}
	return ret, nil
}

// Deprecated: use `ListByMsg` instead, passing the endpoint ID through options
func (m *MessageAttempt) ListAttemptsForEndpoint(ctx context.Context, appId string, msgId string, endpointId string, options *MessageAttemptListOptions) (*ListResponseMessageAttemptEndpointOut, error) {
	req := m.api.MessageAttemptAPI.V1MessageAttemptListByEndpointDeprecated(ctx, appId, msgId, endpointId)
	if options != nil {
		if options.Iterator != nil {
			req = req.Iterator(*options.Iterator)
		}
		if options.Limit != nil {
			req = req.Limit(*options.Limit)
		}
		if options.Status != nil {
			req = req.Status(openapi.MessageStatus(*options.Status))
		}
		if options.EventTypes != nil {
			req = req.EventTypes(*options.EventTypes)
		}
		if options.Before != nil {
			req = req.Before(*options.Before)
		}
		if options.After != nil {
			req = req.After(*options.After)
		}
		if options.Channel != nil {
			req = req.Channel(*options.Channel)
		}
		if options.Tag != nil {
			req = req.Tag(*options.Tag)
		}
	}
	ret, res, err := req.Execute()
	if err != nil {
		return nil, wrapError(err, res)
	}
	return ret, nil
}

func (m *MessageAttempt) ExpungeContent(ctx context.Context, appId string, msgId string, attemptId string) error {
	req := m.api.MessageAttemptAPI.V1MessageAttemptExpungeContent(ctx, appId, msgId, attemptId)
	res, err := req.Execute()
	return wrapError(err, res)
}
