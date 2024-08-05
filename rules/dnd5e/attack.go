package dnd5e

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"trpg/rules/dnd5e/damage"
	"trpg/utils"
)

func (c *Character) Attack(target *Character) *Damage {
	damage := NewDamage(damage.Slashing, 2, 6)

	damage.Source = c
	damage.Target = target

	return damage
}

type Damage struct {
	sync.Mutex
	Dice    [2]int
	Type    damage.Type
	Fixed   int
	Result  []int
	History [][]int
	Source  *Character
	Target  *Character
}

func NewDamage(damageType damage.Type, num, dice int) *Damage {
	return &Damage{Dice: [2]int{num, dice}, Type: damageType}
}

func (d *Damage) SetFixed(fixed int) *Damage {
	d.Fixed = fixed

	return d
}

func (d *Damage) SetSource(source *Character) *Damage {
	d.Source = source

	return d
}

func (d *Damage) SetTarget(target *Character) *Damage {
	d.Target = target

	return d
}

func (d *Damage) String() string {
	total := d.Fixed
	res := []string{}
	hasResistance := false
	hasVulnerability := false

	d.Lock()
	if len(d.Result) == 0 {
		d.Result = utils.Roll(d.Dice[0], d.Dice[1])
	}

	for i := range d.Result {
		res = append(res, strconv.Itoa(d.Result[i]))
		total += d.Result[i]
	}
	d.Unlock()

	if d.Target != nil {
		_, hasResistance = d.Target.Resistance[d.Type]
		_, hasVulnerability = d.Target.Vulnerability[d.Type]
	}

	var extra string

	if hasResistance && hasVulnerability {
		// Don't do anything
	} else if hasResistance {
		total = (total - (total % 2)) / 2
		extra = "「抗性减半」"
	} else if hasVulnerability {
		total = total * 2
		extra = "「易伤翻倍」"
	}

	if d.Fixed == 0 {
		return fmt.Sprintf("%d (%s) 的%s伤害"+extra, total, strings.Join(res, "+"), d.Type.Name())
	} else {
		return fmt.Sprintf("%d (%s+%d) 的%s伤害"+extra, total, strings.Join(res, "+"), d.Fixed, d.Type.Name())
	}
}
