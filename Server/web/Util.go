package web

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

// GetUUID return uuid
func GetUUID() uuid.UUID {
	id := uuid.NewV4()
	return id
}

// ParseID Parse uuid from string
func ParseID(id string) (*uuid.UUID, error) {
	uid, err := uuid.FromString(id)
	if err != nil {
		return nil, errors.New("Invalid User ID")
	}
	return &uid, nil
}

// ParseOrNil Return UUID or NilUUID.
func ParseOrNil(id string) *uuid.UUID {
	uid := uuid.FromStringOrNil(id)
	return &uid
}
