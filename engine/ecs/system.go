package ecs

// System represents a system that processes entities with specific components.
type System interface {
	Update(entities []Entity, dt float64)
}

// BaseSystem provides a base implementation of a system.
type BaseSystem struct {
	Name string
}

// Update is a default implementation of the Update method.
func (s *BaseSystem) Update(entities []Entity, dt float64) {
	// This can be overridden by specific systems.
}
