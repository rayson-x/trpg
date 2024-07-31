package battle

import (
	"trpg/rules/dnd5e"
	"trpg/utils"
)

type Combat struct {
	Name       string
	Round      int
	Characters *utils.List[*CombatCharacter]
}

type CombatCharacter struct {
	dnd5e.Character

	Speed int
}

func NewCombat(name string) *Combat {
	return &Combat{Name: name, Round: 1, Characters: utils.NewList[*CombatCharacter]()}
}

func (c *Combat) AddCharacter(speed int, character dnd5e.Character) {
	elem := c.Characters.Front()
	new := &CombatCharacter{Speed: speed, Character: character}

	for {
		if elem == nil {
			c.Characters.PushBack(new)
			return
		}

		if speed > elem.Value.Speed {
			c.Characters.InsertBefore(new, elem)
			return
		} else if speed == elem.Value.Speed {
			// 玩家优先行动
			if character.Type == dnd5e.TypePlayer {
				c.Characters.InsertBefore(new, elem)
			} else {
				c.Characters.InsertAfter(new, elem)
			}
			return
		}

		elem = elem.Next()
	}
}
