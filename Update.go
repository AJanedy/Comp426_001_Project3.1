package main

func (game *Game) Update() error {
	if game.player.keysCollected == 1 {
		game.level = 2
	}
	if game.player.keysCollected == 2 {
		game.level = 3
	}
	if game.player.health <= 0 {
		game.level = 0
	}
	for i, _ := range game.barrierTiles {
		if checkBarrierCollision(game, i) {
			game.wallDetected = true
		}
	}
	if game.enemies.smileyMan.health <= 0 {
		game.level = 4
	}
	updatePlayerAnimations(game)
	animateHearts(game)
	animateSpeedBoost(game)
	animateManaPot(game)
	animateKey(game)
	checkPowerUpCollision(game)

	if !game.gameSounds.gameMusic.IsPlaying() {
		playGameMusic(game)
	}
	if game.level == 1 {
		animateLevel3(game)
		updateRainbowMenAnimations(game)
		checkRainbowMenHitByFireballs(game)
		checkRainbowMenHitByNukes(game)
		checkRainbowMenShotsFired(game)
		checkPlayerRainbowMenProximity(game)
	}
	if game.level == 2 {
		animateLevel3(game)
		updateBeeAnimations(game)
		checkBeesHitByFireballs(game)
		checkBeesHitByNukes(game)
		checkBeeShotsFired(game)
		checkPlayerBeeProximity(game)
	}
	if game.level == 3 {
		animateLevel3(game)
		animateSmileyMan(game)
		moveSmileyMan(game)
		for i, _ := range game.player.fireballs {
			isSmileyManHitByFireball(game, i)
		}
		for i, _ := range game.player.nukes {
			isSmileyManHitByNuke(game, i)
		}
		didSmileyManCatchPlayer(game)
	}
	return nil
}

func updateBeeAnimations(game *Game) {
	animateBee(game)
	removeDeadBees(game)
	attackingBees(game)
	moveBeeAttack(game)
	animateBeeAttack(game)
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
