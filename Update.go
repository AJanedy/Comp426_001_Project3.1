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
	movePlayer(game)
	moveFireballs(game)
	animatePlayerFireballs(game)
	moveIceWall(game)
	animateIceWall(game)
	limitCursorDistanceFromPlayer(game)
	animatePlayerNuke(game)
	getPlayerInput(game)
}
func (game Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
