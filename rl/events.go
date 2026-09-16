package rl

import (
	"github.com/bennicholls/tyumi/event"
	"github.com/bennicholls/tyumi/vec"
)

var (
	EV_ENTITYBEINGDESTROYED  = event.Register("Entity being destroyed/removed from the ECS.")
	EV_TILECHANGEDVISIBILITY = event.Register("A Tile Changed visibility state (opaque or transparent)")
	
)

type EntityEvent struct {
	event.EventPrototype

	Entity Entity
}

type TileChangedVisibilityEvent struct {
	event.EventPrototype

	Pos    vec.Coord
	Opaque bool
}

