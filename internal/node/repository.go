package node

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, node *Node) error
	GetByID(ctx context.Context, id string) (*Node, error)
	AddContent(ctx context.Context, nodeID uuid.UUID, item ContentItem) error
}
