package node

import (
	"time"

	"github.com/google/uuid"
)

type Node struct {
	ID         uuid.UUID
	Type       string
	Properties map[string]interface{}
	CreatedAt  time.Time
}
