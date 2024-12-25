package ecs

// Component is an interface that all components must implement.
type Component interface {
	GetID() int
}

// BaseComponent provides a base structure for components.
type BaseComponent struct {
	ID int
}

// GetID returns the unique ID of the component.
func (c *BaseComponent) GetID() int {
	return c.ID
}
