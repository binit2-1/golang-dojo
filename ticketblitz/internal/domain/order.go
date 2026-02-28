package domain

import "time"

type Orders struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	EventID   string    `json:"event_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
