package events

// NewBlockEvent :
type NewBlockEvent struct {
	BaseEvent
}

// CreateNewBlockEvent :
func CreateNewBlockEvent(blockNumber uint64) NewBlockEvent {
	return NewBlockEvent{
		BaseEvent{
			Name:        NewBlockEventName,
			BlockNumber: blockNumber,
		},
	}
}