package damage

var Types = []Type{
	Piercing,
	Bludgeoning,
	Slashing,
	Cold,
	Fire,
	Lightning,
	Thunder,
	Poison,
	Acid,
	Necrotic,
	Radiant,
	Force,
	Psychic,
}

var Piercing = defaultType{name: "穿刺", desc: "矛、怪物的啃咬等造成的伤害"}

var Bludgeoning = defaultType{name: "钝击", desc: "锤击，坠地，绞压等造成的伤害"}

var Slashing = defaultType{name: "挥砍", desc: "剑、斧、怪物的爪击等造成伤害"}

var Cold = defaultType{name: "寒冷", desc: "冰魔长矛所散发的炼狱刺寒、以及白龙吐息产生的急冻风暴，都会造成寒冰伤害"}

var Fire = defaultType{name: "火焰", desc: "红龙吐出的火焰，以及许多会产生烈焰的法术都会造成火焰伤害"}

var Lightning = defaultType{name: "闪电", desc: "闪电束法术和蓝龙的吐息都会造成闪电伤害"}

var Thunder = defaultType{name: "雷鸣", desc: "冲击性的音爆，像是雷鸣波法术的效果，会造成雷鸣伤害"}

var Poison = defaultType{name: "毒素", desc: "毒刺和绿龙吐息的巨毒气体都会造成毒素伤害"}

var Acid = defaultType{name: "强酸", desc: "黑龙吐息的腐蚀性喷雾，以及黑布丁分泌的溶解酵素，都会造成强酸伤害"}

var Necrotic = defaultType{name: "黯蚀", desc: "由某些不死生物和法术所造成的黯蚀伤害，会使物质凋零，甚至侵蚀灵魂"}

var Radiant = defaultType{name: "光耀", desc: "由牧师焰击术法术或天使神圣武器所造成的光耀伤害，会如同火焰般灼烧肉体，并使灵魂因能量而过载"}

var Force = defaultType{name: "立场", desc: "力场是被凝聚成致伤型态的纯粹魔法能量。大部分能够造成力场伤害的效果都是法术，包括魔法飞弹和灵体武器"}

var Psychic = defaultType{name: "心灵", desc: "心灵能力，例如夺心魔的心灵震爆，会造成心灵伤害"}

type Type interface {
	Name() string
	Desc() string
}

type defaultType struct {
	name string
	desc string
}

func (t defaultType) Name() string {
	return t.name
}

func (t defaultType) Desc() string {
	return t.desc
}
