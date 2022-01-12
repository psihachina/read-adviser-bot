package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(e Event)
}

type Type int

const (
	Unknow Type = iota
	Message
)

type Event struct {
	Type Type
	Text string
}
