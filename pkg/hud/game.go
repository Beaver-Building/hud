package boardgamemaster

type game struct {
	playersOrder map[int]int
	stage        map[int]string
}

type gameHUD interface { // dealer HUD
	getGameStatus()
	setGameStatus()
}

func newGame() {
	/*
	** loadConfig()
	** loadDealerComp()
	** loadPlayers()
	 */
}

func loadStage() {
}

func settleOfGame() bool {
	return false
}
