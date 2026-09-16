package rl

import (
	"github.com/bennicholls/tyumi/log"
	"github.com/bennicholls/tyumi/rl/ecs"
)

func init() {
	ecs.Register[EntityContainerComponent]()
}

type EntityContainerComponent struct {
	ecs.Component

	Entity
}

func (ecc EntityContainerComponent) Empty() bool {
	return ecc.Entity == Entity(ecs.INVALID_ID)
}

func (ecc *EntityContainerComponent) Add(entity Entity) {
	if ecc.Entity != Entity(ecs.INVALID_ID) {
		log.Debug("Overwriting entity!!")
	}
	ecc.Entity = entity
}

func (ecc *EntityContainerComponent) Remove() {
	ecc.Entity = Entity(ecs.INVALID_ID)
}
