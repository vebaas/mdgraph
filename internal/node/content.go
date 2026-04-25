package node

// ContentKind represents the type of content stored in a ContentItem.
type ContentKind string

const (
	// ContentKindText represents plain text content.
	ContentKindText ContentKind = "text"
	// ContentKindImage represents an image referenced by URL.
	ContentKindImage ContentKind = "image"
)

// ContentItem represents a single piece of content on a node.
// MinLOD controls the minimum zoom level at which this item becomes visible.
// A MinLOD of 0 means the item is always visible.
type ContentItem struct {
	Kind   ContentKind
	Value  string
	MinLOD int
}
