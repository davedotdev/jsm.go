package advisory

import (
	"github.com/nats-io/jsm.go/api/event"
)

// ActionAdvisoryTypeV1 indicates which action against a stream, consumer or template triggered an advisory
type ActionAdvisoryTypeV1 string

const (
	CreateEvent ActionAdvisoryTypeV1 = "create"
	DeleteEvent ActionAdvisoryTypeV1 = "delete"
	ModifyEvent ActionAdvisoryTypeV1 = "modify"
)

// JetStreamAPIAuditV1 is a advisory published on create, modify or delete of a Stream
//
// NATS Schema Type io.nats.jetstream.advisory.v1.stream_action
type JSStreamActionAdvisoryV1 struct {
	event.NATSEvent

	Stream   string               `json:"stream"`
	Action   ActionAdvisoryTypeV1 `json:"action"`
	Template string               `json:"template,omitempty"`
}