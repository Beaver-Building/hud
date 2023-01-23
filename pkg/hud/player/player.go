package player

type Ability struct {
	desc      string
	typeRefer int // 0 1 2: Active Passive Auto-cast
	effect    int // ack-heal-gold(to self)--ack-heal-gold(to target)
}

func (a Ability) active(selfId, targetId int) int {
	return 0
}

type player struct { // player HUD
	ack  int
	hp   int
	gold int

	abilities []Ability
}

type playerHUD interface {
	getPlayerInfo()
	setPlayerInfo()
}

func (p player) movement() (next turnEvent) {
	// loadTrigger()
	// move()
	// settleOfTurn()
	return turnEvent{tid: 1, desc: "next turn"}
}
