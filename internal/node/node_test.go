package node

import (
	"testing"

	"github.com/google/uuid"
)

func TestNonEmptyUUID(t *testing.T) {
	node := NewNode(NodeTypeGeneric)

	if node.ID == uuid.Nil {
		t.Error("expected node ID to be non-empty but it was empty")
	}
}

func TestType(t *testing.T) {
	node := NewNode(NodeTypeGeneric)

	if node.Type != NodeTypeGeneric {
		t.Error("expected node type to be generic but it was not")
	}
}

func TestContent(t *testing.T) {
	node := NewNode(NodeTypeGeneric)

	if node.Content == nil {
		t.Error("expected node content to be initialized but it was nil")
	}
}

func TestProperties(t *testing.T) {
	node := NewNode(NodeTypeGeneric)

	if node.Properties == nil {
		t.Error("expected node properties to be initialized but it was nil")
	}
}
