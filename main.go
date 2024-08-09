package main

import (
	"trpg/rules/dnd5e"
	"trpg/rules/dnd5e/damage"
)

func main() {
	// 穿刺
	dnd5e.NewDamage(damage.Piercing, 3, 8).SetFixed(0)

	// 钝击
	dnd5e.NewDamage(damage.Bludgeoning, 3, 8).SetFixed(0)

	// 挥砍
	dnd5e.NewDamage(damage.Slashing, 3, 8).SetFixed(0)

	// 寒冷
	dnd5e.NewDamage(damage.Cold, 3, 8).SetFixed(0)

	// 火焰
	dnd5e.NewDamage(damage.Fire, 3, 8).SetFixed(0)

	// 闪电
	dnd5e.NewDamage(damage.Lightning, 3, 8).SetFixed(0)

	// 雷鸣
	dnd5e.NewDamage(damage.Thunder, 3, 8).SetFixed(0)

	// 毒素
	dnd5e.NewDamage(damage.Poison, 3, 8).SetFixed(0)

	// 强酸
	dnd5e.NewDamage(damage.Acid, 3, 8).SetFixed(0)

	// 黯蚀
	dnd5e.NewDamage(damage.Necrotic, 3, 8).SetFixed(0)

	// 光耀
	dnd5e.NewDamage(damage.Radiant, 3, 8).SetFixed(0)

	// 力场
	dnd5e.NewDamage(damage.Force, 3, 8).SetFixed(0)

	// 心灵
	dnd5e.NewDamage(damage.Psychic, 3, 8).SetFixed(0)

	// &dnd5e.Character{
	// 	Resistance: map[damage.Type]struct{}{
	// 		damage.Force: {},
	// 	},
	// 	Vulnerability: map[damage.Type]struct{}{
	// 		// damage.Force: {},
	// 	},
	// }
	// fmt.Println(
	// 	utils.Roll(22, 8),
	// 	dnd5e.NewDamage(damage.Force, 6, 10).SetFixed(40).
	// 		SetTarget(&dnd5e.Character{
	// 			Resistance: map[damage.Type]struct{}{
	// 				damage.Force: {},
	// 			},
	// 			Vulnerability: map[damage.Type]struct{}{
	// 				// damage.Force: {},
	// 			},
	// 		}).
	// 		String(),
	// )
}
