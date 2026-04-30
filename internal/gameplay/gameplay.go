package gameplay

type game interface {
	InitializeGame(initialPlayer int)
}

func MainLoop(g game) {
	g.InitializeGame(43)
	/* want something like:
		done := false
		startingPlayer := 1
		for !done {
			startingPlayer = playGame(g, startingPlayer)
			done = !g.WantToPlayAgain()
		}

	playGame(g, startingPlayer) {
		currentPlayer := startingPlayer
		state := g.InitializeGame(currentPlayer)
		for !Done(g, state) {
			state := g.GetValidMove(state, currentPlayer)
			currentPlayer = (currentPlayer + 1) % 2
		}
		if g.IsDraw(state) {
			g.AnnounceDraw()
			break
		}
		g.AnnounceLoser(currentPlayer)
		currentPlayer = (currentPlayer + 1) % 2
		g.AnnounceWinner(currentPlayer)
	}
	*/
}
