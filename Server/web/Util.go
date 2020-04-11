package web

import (
	"errors"
	"strconv"

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

// ParseInt64 Parse string to integer64
func ParseInt64(s string) *int64 {
	v, err := strconv.ParseInt(s, 8, 64)
	if err != nil {
		v = 0
		return &v
	}
	return &v
}
