package repositories

import (
	"context"
	"time"

	"github.com/hauzanrafiattallah/ticket-booking-backend/models"
)

type EventRepository struct {
	db any
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}

	events = append(events, &models.Event{
		ID:        "1",
		Name:      "Event 1",
		Location:  "Location 1",
		Date:      time.Now(),
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	})

	return events, nil
}

func (r *EventRepository) GetOne(ctx context.Context, eventId string) (*models.Event, error) {
	return nil, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event models.Event) (*models.Event, error) {
	return nil, nil
}

func NewEventRepository(db any) models.EventRepository {
	return &EventRepository{
		db: db,
	}
}
