package postgres

import (
	"database/sql"
	"ticketblitz/internal/domain"
)


type PostgresEventRepo struct{
	db *sql.DB
}

func NewPotgresEventRepo(db *sql.DB) domain.EventRepository{
	return &PostgresEventRepo{
		db: db,
	}
}

func(h *PostgresEventRepo) GetEventByID(id string)(*domain.Event, error){

	query := `SELECT id, name, total_tickets, availed_tickets FROM events WHERE id=$1`

	var event domain.Event

	err := h.db.QueryRow(query, id).Scan(
		&event.ID,
		&event.Name,
		&event.TotalTickets,
		&event.TotalTickets,
	)

	if err!=nil{
		if err == sql.ErrNoRows{
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &event, nil
}