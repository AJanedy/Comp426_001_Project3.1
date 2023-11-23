package main

import "math"

func animateSmileyMan(game *Game) {
	game.enemies.smileyMan.frameDelay += 1
	if game.enemies.smileyMan.frameDelay%FRAMES_PER_SHEET == 0 {
		game.enemies.smileyMan.frame += 1
		if game.enemies.smileyMan.frame >= SMILEY_FRAMES_PER_SHEET {
			game.enemies.smileyMan.frame = 0
		}
	}
}
func moveSmileyMan(game *Game) {
	game.enemies.smileyMan.degreesDirection = math.Atan2(float64(game.player.yLoc-game.enemies.smileyMan.yLoc),
		float64(game.player.xLoc-(game.enemies.smileyMan.xLoc+160)))
	game.enemies.smileyMan.radiansDirection = game.enemies.smileyMan.degreesDirection * 180 / math.Pi

	if math.Cos(game.enemies.smileyMan.degreesDirection) > 0 {
		game.enemies.smileyMan.direction = 1
	}
	if math.Cos(game.enemies.smileyMan.degreesDirection) < 0 {
		game.enemies.smileyMan.direction = 0
	}
	game.enemies.smileyMan.xLoc += math.Cos(game.enemies.smileyMan.degreesDirection) * 1.5
	game.enemies.smileyMan.yLoc += math.Sin(game.enemies.smileyMan.degreesDirection) * 1.5
}
func animateLevel3(game *Game) {
	game.maps.frameDelay += 1
	if game.maps.frameDelay%FRAMES_PER_SHEET == 0 {
		game.maps.frame += 1
		if game.maps.frame >= 8 {
			game.maps.frame = 1
		}
	}
}
