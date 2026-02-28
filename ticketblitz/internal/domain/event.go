package domain

type Event struct{
	ID string `json:"id"`
	Name string `json:"name"`
	TotalTickets int `json:"total_tickets"`
	AvailableTickets int `json:"available_tickets"`
}

type EventRepository interface{
	GetEventByID(id string) (*Event, error)
	CreateEvent(event *Event) error
}