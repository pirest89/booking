package requests

import "github.com/google/uuid"

type SearchFlightRequest struct {
	Level int64
}

type FindJobRequest struct {
	Id uuid.UUID
}
