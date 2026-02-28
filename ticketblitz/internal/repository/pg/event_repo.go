package pg

import (
	"database/sql"
	"ticketblitz/internal/domain"
)

type PostgresEventRepo struct {
	db *sql.DB
}



func NewPotgresEventRepo(db *sql.DB) domain.EventRepository {
	return &PostgresEventRepo{
		db: db,
	}
}

//GetEventByID implements [domain.EventRepository].
func (h *PostgresEventRepo) GetEventByID(id string) (*domain.Event, error) {

	query := `SELECT id, name, total_tickets, availed_tickets FROM events WHERE id=$1`

	var event domain.Event

	err := h.db.QueryRow(query, id).Scan(
		&event.ID,
		&event.Name,
		&event.TotalTickets,
		&event.TotalTickets,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &event, nil
}

// CreateEvent implements [domain.EventRepository].
func (h *PostgresEventRepo) CreateEvent(event *domain.Event) error {
	
	query:= `INSERT INTO events (name, total_tickets, available_tickets) VALUES ($1, $2, $3) RETURNING id`

	err := h.db.QueryRow(query, &event.Name, &event.TotalTickets, &event.AvailedTickets).Scan(&event.ID)

	return err
}