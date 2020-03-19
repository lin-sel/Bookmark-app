package web

import (
	"errors"

	"github.com/google/uuid"
)

// GetUUID return uuid
func GetUUID() uuid.UUID {
	id, err := uuid.NewRandom()
	if err != nil {
		return GetUUID()
	}
	return id
}

// ParseID Parse uuid from string
func ParseID(id string) (*uuid.UUID, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("Invalid User ID")
	}
	return &uid, nil
}
