package persistence

import (
	"context"
	"database/sql"
	"melodia-events/internal/domain/entities"

	"github.com/google/uuid"
)

type PostgresEventRepository struct {
	DB *sql.DB
}

func (r *PostgresEventRepository) Register(ctx context.Context, event *entities.Event) error {

	query := `INSERT INTO events (title, topic, description, service, payload, error_level, token, is_published) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
			  RETURNING id`

	var id uuid.UUID
	err := r.DB.QueryRowContext(ctx, query,
		event.Title, event.Topic, event.Description, event.Service, event.Payload, event.ErrorLevel, event.Token, event.IsPublished).Scan(&id)
	if err != nil {
		return err
	}

	event.ID = id

	return nil
}
