package pg

import (
	"database/sql"
	"fmt"
	"ticketblitz/internal/domain"
)

type PostgresEventRepo struct {
	db *sql.DB
}



func NewPostgresEventRepo(db *sql.DB) domain.EventRepository {
	return &PostgresEventRepo{
		db: db,
	}
}

//GetEventByID implements [domain.EventRepository].
func (h *PostgresEventRepo) GetEventByID(id string) (*domain.Event, error) {

	query := `SELECT id, name, total_tickets, available_tickets FROM events WHERE id=$1`

	var event domain.Event

	err := h.db.QueryRow(query, id).Scan(
		&event.ID,
		&event.Name,
		&event.TotalTickets,
		&event.AvailableTickets,
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

	err := h.db.QueryRow(query, &event.ID, &event.TotalTickets, &event.AvailableTickets).Scan(&event.ID)

	return err
}

// PurchaseTicket implements [domain.EventRepository]
func(h *PostgresEventRepo)PurchaseTicket(userID string, eventID string) error{
	tx, err := h.db.Begin()
	if err != nil{
		return err
	}
	defer tx.Rollback()	

	

	query := `SELECT available_tickets FROM events WHERE id=$1 FOR UPDATE`
	var availableTickets int

	err = tx.QueryRow(query, eventID).Scan(
		&availableTickets,
	)
	if err != nil{
		if err == sql.ErrNoRows{
			return fmt.Errorf("Event Not Found")
		}
		return fmt.Errorf("failed to lock event: %w", err)
	}

	if availableTickets <= 0 {
		return fmt.Errorf("sold out")
	}

	updateQuery := `UPDATE events SET available_tickets = available_tickets - 1 WHERE id=$1`

	_, err = tx.Exec(updateQuery, eventID)
	if err != nil {
		return fmt.Errorf("failed to update ticket count: %w", err)
	}

	insertQuery := `INSERT INTO orders (user_id, event_id, status) VALUES ($1, $2, $3)`

	_, err = tx.Exec(insertQuery, userID, eventID, "paid")
	if err != nil{
		return fmt.Errorf("failed to insert order: %w", err)
	}

	return tx.Commit()
}