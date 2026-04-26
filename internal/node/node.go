// Package node provides the core Node type and related content structures
// for the mdgraph multi-dimensional node graph engine.
package node

import (
	"mdgraph/internal/registry"
	"time"

	"github.com/google/uuid"
)

// NodeType represents the type of a node in the graph.
// Applications register their own node types via the engine registry.
type NodeType string

const (
	// NodeTypeGeneric is the default node type available to all applications.
	NodeTypeGeneric NodeType = "generic"
)

// Node is the fundamental unit of the graph engine.
// Every node has a globally unique ID that never changes.
// Nodes participate in multiple dimensions through edges.
type Node struct {
	ID         uuid.UUID
	Type       NodeType
	Content    []ContentItem
	Properties map[string]interface{}
	CreatedAt  time.Time
	UpdatedAt  time.Time
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

// RegisterBaseTypes registers the built-in node types into the provided registry.
// Call this once when the engine starts before registering application types.
func RegisterBaseTypes(reg *registry.Registry) {
	reg.Register(string(NodeTypeGeneric))
}
