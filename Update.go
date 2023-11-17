package main

func (game *Game) Update() error {
	updatePlayerAnimations(game)
	updateEnemy1Animations(game)
	checkProximityAndCollisions(game)
	return nil
}

func updateEnemy1Animations(game *Game) {
	moveLevelOneEnemy(game)
	animateEnemy1(game)
	enemy1Attackers(game)
	moveEnemy1Fireballs(game)
	animateEnemy1Fireball(game)
	removeDeadLevelOneEnemies(game)
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
