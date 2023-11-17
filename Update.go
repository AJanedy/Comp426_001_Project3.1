package main

func (game *Game) Update() error {
	updatePlayerAnimations(game)
	//updateRainbowMenAnimations(game)
	//checkProximityAndCollisions(game)
	updateBeeAnimations(game)
	return nil
}

func updateBeeAnimations(game *Game) {
	animateBee(game)
	//moveBee(game)
}

func updateRainbowMenAnimations(game *Game) {
	moveRainbowMen(game)
	animateRainbowMen(game)
	attackingRainbowMen(game)
	moveRainbowMenFireballs(game)
	animateRainbowMenFireballs(game)
	removeDeadRainbowMen(game)
}

func updatePlayerAnimations(game *Game) {

	moveFireballs(game)
	animatePlayerFireballs(game)
	moveBuzzBall(game)
	animateBuzzBall(game)
	limitCursorDistanceFromPlayer(game)
	animatePlayerNuke(game)
	getPlayerInput(game)
	if game.player.canMove {
		movePlayer(game)
	}
	animatePlayer(game)

}
func (game Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
