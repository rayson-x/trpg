package dnd5e

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
	// Skills []
	// Conditions []Condition
}
