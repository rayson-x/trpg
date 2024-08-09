package dnd5e

import (
	"fmt"
	"strconv"
	"strings"
	"trpg/rules/dnd5e/damage"
)

const (
	TypePlayer = 1
	TypeEnemy  = 2
	TypeOther  = 3
)

const (
	StatusNormal      = 1 // 正常
	StatusDeath       = 2 // 死亡
	StatusKnockDown   = 3 // 击倒
	StatusUnconscious = 4 // 昏迷
)

type Character struct {
	Name   string
	Player string

	MaxHP  int
	HP     int
	Status int

	Type int

	History []string

	Resistance    map[damage.Type]struct{}
	Vulnerability map[damage.Type]struct{}
	// Skills []
	// Conditions []Condition
}

func (c *Character) TakingDamage(list []*Damage) {
	for i := range list {
		res := []string{}
		hasResistance := false
		hasVulnerability := false
		totalDamage := list[i].Fixed

		for _, v := range list[i].GetResult() {
			totalDamage += v
			res = append(res, strconv.Itoa(v))
		}

		if list[i].Target != nil {
			_, hasResistance = list[i].Target.Resistance[list[i].Type]
			_, hasVulnerability = list[i].Target.Vulnerability[list[i].Type]
		}

		var extra string

		if hasResistance && hasVulnerability {
			// Don't do anything
		} else if hasResistance {
			totalDamage = (totalDamage - (totalDamage % 2)) / 2
			extra = "「抗性减半」"
		} else if hasVulnerability {
			totalDamage = totalDamage * 2
			extra = "「易伤翻倍」"
		}

		if list[i].Fixed == 0 {
			c.History = append(c.History, fmt.Sprintf("受到了 %d (%s) 的%s伤害"+extra, totalDamage, strings.Join(res, "+"), list[i].Type.Name()))
		} else {
			c.History = append(c.History, fmt.Sprintf("受到了 %d (%s)+%d 的%s伤害"+extra, totalDamage, strings.Join(res, "+"), list[i].Fixed, list[i].Type.Name()))
		}

		c.HP -= totalDamage
	}
}
