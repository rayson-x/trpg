package feats

import (
	"trpg/rules/dnd5e"
)

type BattleFeat interface {
	InDamage(*dnd5e.Damage)
}

// type BattleFeat interface {
// 	BeforeAttack()
// 	AfterAttack()

// 	BeforeSpellcasting()
// 	AfterSpellcasting()
// }

// type CustomFeat struct {
// }
