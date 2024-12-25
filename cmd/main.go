package main

import (
	"fmt"

	"github.com/Galygious/GalygoEngine/engine/ecs"
)

func main() {
	fmt.Println("Welcome to Galygo Engine!")

	// Example: Create an entity and attach a component to it
	entity := ecs.Entity{ID: 1}
	component := ecs.BaseComponent{ID: 101}

	fmt.Printf("Entity ID: %d\n", entity.ID)
	fmt.Printf("Component ID: %d\n", component.GetID())

	// Example: Create a system and simulate an update
	system := ecs.BaseSystem{Name: "RenderSystem"}
	fmt.Printf("System Name: %s\n", system.Name)
	system.Update([]ecs.Entity{entity}, 0.016) // Simulate 16ms frame update
}
