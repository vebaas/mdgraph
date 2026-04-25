// Package node provides the core Node type and related content structures
// for the mdgraph multi-dimensional node graph engine.
package node

import (
	"time"

	"github.com/google/uuid"
)

// NodeType represents the type of a node in the graph.
// Applications built on the engine define their own node types
// by registering them with the engine registry.
type NodeType string

const (
	// NodeTypeGeneric is the default node type available to all applications.
	NodeTypeGeneric NodeType = "generic"
)

// Node is the fundamental unit of the graph engine.
// Every node has a globally unique ID that never changes.
// Nodes participate in multiple dimensions through edges.
type Node struct {
	// ID is the permanent global identifier for this node.
	ID uuid.UUID
	// Type describes what kind of node this is.
	Type NodeType
	// Content holds the node's content items at various LOD levels.
	Content []ContentItem
	// Properties holds flexible structured data specific to the node type.
	Properties map[string]interface{}
	// CreatedAt is the time this node was created.
	CreatedAt time.Time
	// UpdatedAt is the time this node was last modified.
	UpdatedAt time.Time
}

// NewNode creates a new initialized Node with a generated UUID.
// The caller must provide a NodeType — use NodeTypeGeneric for a basic node
// or a registered application type for domain-specific nodes.
func NewNode(nodeType NodeType) *Node {
	return &Node{
		ID:         uuid.New(),
		Type:       nodeType,
		Content:    []ContentItem{},
		Properties: make(map[string]interface{}),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
