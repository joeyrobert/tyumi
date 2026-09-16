package rl

import (
	"github.com/bennicholls/tyumi/event"
	"github.com/bennicholls/tyumi/rl/ecs"
	"github.com/bennicholls/tyumi/vec"
)

func init() {
	ecs.Register[PositionComponent]()
}

var EV_ENTITYMOVED = event.Register("Entity moved.")

type EntityMovedEvent struct {
	event.EventPrototype

	Entity   Entity
	From, To vec.Coord
}

func fireEntityMovedEvent(entity Entity, from, to vec.Coord) {
	event.Fire(EV_ENTITYMOVED, &EntityMovedEvent{Entity: entity, From: from, To: to})
}

type PositionComponent struct {
	ecs.Component
	vec.Coord

	Static bool
}
