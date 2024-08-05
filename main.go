package main

import (
	"fmt"
	"trpg/rules/dnd5e"
	"trpg/rules/dnd5e/damage"
)

func main() {
	fmt.Println(
		dnd5e.NewDamage(damage.Force, 6, 10).
			SetFixed(40).
			SetTarget(&dnd5e.Character{
				Resistance: map[damage.Type]struct{}{
					damage.Force: {},
				},
				Vulnerability: map[damage.Type]struct{}{
					// damage.Force: {},
				},
			}).
			String(),
	)
}
