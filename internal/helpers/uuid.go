package helpers

import (
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
)

func UUIDFromString(id string) (*uuid.UUID, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID %q: %w", id, err)
	}

	return &uid, nil
}

func GraphQLStringIDFromUUID(uid uuid.UUID, resourceName string) string {
	id := fmt.Sprintf("%s:%s", resourceName, uid.String())

	return base64.StdEncoding.EncodeToString([]byte(id))
}
