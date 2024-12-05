package enum

type EventType string

const (
	Posted  EventType = "POSTED"
	Queued  EventType = "QUEUED"
	Failed  EventType = "FAILED"
	Success EventType = "SUCCESS"
)
