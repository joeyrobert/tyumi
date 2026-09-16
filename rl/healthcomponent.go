package rl

import (
	"github.com/bennicholls/tyumi/event"
	"github.com/bennicholls/tyumi/rl/ecs"
)

func init() {
	ecs.Register[HealthComponent]()
}

type HealthComponent struct {
	ecs.Component

	HP Stat[int]
}

func (hc *HealthComponent) SetHealth(health int) {
	oldHealth := hc.HP.Get()

	if oldHealth == health {
		return
	}

	hc.HP.Set(health)

	if hc.HP.Get() == oldHealth {
		return
	}

	e := Entity(hc.GetEntity())
	event.Fire(EV_ENTITYHEALTHCHANGED, &EntityHealthChangedEvent{
		Entity: e,
		OldHP:  oldHealth,
		NewHP:  hc.HP.Get(),
	})

	if hc.HP.Get() == 0 {
		event.Fire(EV_ENTITYDIED, &EntityEvent{Entity: e})
	}
}

func (hc *HealthComponent) ChangeHealth(delta int) {
	if delta == 0 {
		return
	}

	hc.SetHealth(hc.HP.Get() + delta)
}

func (hc HealthComponent) Dead() bool {
	return hc.HP.Get() == 0
}
