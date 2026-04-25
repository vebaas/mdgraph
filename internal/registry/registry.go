// Package registry provides a generic key registry for validating
// registered types across the engine.
package registry

// Registry holds a set of registered keys and validates against them.
type Registry struct {
	keys map[string]bool
}

// NewRegistry creates a new initialized Registry.
func NewRegistry() *Registry {
	return &Registry{
		keys: make(map[string]bool),
	}
}

// Register adds a new valid key to the registry.
func (r *Registry) Register(name string) {
	r.keys[name] = true
}

// IsValid returns true if the given key has been registered.
func (r *Registry) IsValid(name string) bool {
	return r.keys[name]
}
