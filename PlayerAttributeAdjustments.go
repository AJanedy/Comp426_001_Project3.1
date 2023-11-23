package main

func lowerPlayerHealth(game *Game, attackStrength int) {
	game.player.health -= attackStrength
	if game.player.health <= 0 {
		playGameOverSound(game)
		game.player.canMove = false
		game.level = 0
	}
}
func updatePlayerAttribute(game *Game, powerUpType int) {
	if powerUpType == HEALTH_POT {
		game.player.health += 15
		if game.player.health > 100 {
			game.player.health = 100
		}
	}
	if powerUpType == SPEED_POT {
		game.player.moveSpeed += .25
	}
	if powerUpType == MANA_POT {
		game.player.magicPower += .25
	}
	if powerUpType == KEY {
		game.player.keysCollected += 1
		game.player.xLoc = WINDOW_WIDTH / 6
		game.player.yLoc = WINDOW_HEIGHT / 2
	}
}
