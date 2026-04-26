package node

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresRepository struct {
	connectionPool *pgxpool.Pool
}

var _ Repository = (*postgresRepository)(nil)

func NewPostgresRepository(connectionPool *pgxpool.Pool) *postgresRepository {
	return &postgresRepository{
		connectionPool: connectionPool,
	}
}

func (r *postgresRepository) Create(ctx context.Context, node *Node) error {
	_, err := r.connectionPool.Exec(ctx,
		"INSERT INTO nodes (id, type, properties) VALUES ($1, $2, $3)",
		node.ID,
		node.Type,
		node.Properties,
	)
	return err
}

func (r *postgresRepository) GetByID(ctx context.Context, id string) (*Node, error) {
	var node Node

	err := r.connectionPool.QueryRow(
		ctx,
		"SELECT id, type, properties, created_at, updated_at FROM nodes WHERE id = $1",
		id,
	).Scan(&node.ID, &node.Type, &node.Properties, &node.CreatedAt, &node.UpdatedAt)
	if err != nil {
		return nil, err
	}

	rows, err := r.connectionPool.Query(
		ctx,
		"SELECT kind, value, min_lod FROM node_content WHERE node_id = $1",
		node.ID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item ContentItem
		err := rows.Scan(&item.Kind, &item.Value, &item.MinLOD)
		if err != nil {
			return nil, err
		}
		node.Content = append(node.Content, item)
	}

	return &node, nil
}

func (r *postgresRepository) AddContent(ctx context.Context, nodeID uuid.UUID, item ContentItem) error {
	_, err := r.connectionPool.Exec(ctx,
		"INSERT INTO node_content (id, node_id, kind, value, min_lod) VALUES ($1, $2, $3, $4, $5)",
		uuid.New(),
		nodeID,
		item.Kind,
		item.Value,
		item.MinLOD,
	)
	return err
}
