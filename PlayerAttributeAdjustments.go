package main

func lowerPlayerHealth(game *Game, attackStrength int) {
	game.player.health -= attackStrength
}
