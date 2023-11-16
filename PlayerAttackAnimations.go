package main

import "math"

func moveFireballs(game *Game) {
	for i, _ := range game.player.fireballs {
		playerFireballsDeltaXY(game, i)
	}
}
func playerFireballsDeltaXY(game *Game, i int) {
	game.player.fireballs[i].xLoc += math.Cos(game.player.fireballs[i].trajectory) * 3
	game.player.fireballs[i].yLoc += math.Sin(game.player.fireballs[i].trajectory) * 3
}
func animatePlayerFireballs(game *Game) {
	for i, _ := range game.player.fireballs {
		incrementPlayerFireballFrameDelay(game)
		if i < len(game.player.fireballs) && game.player.fireballs[i].frameDelay%FRAMES_PER_SHEET == 0 {
			game.player.fireballs[i].row += 1
			if game.player.fireballs[i].row >= FRAMES_PER_SHEET { // frames per row
				rewindPlayerFireballAnimation(game, i)
				if game.player.fireballs[i].column >= 2 {
					removeFireball(game, i)
				}
			}
		}
	}
}
func incrementPlayerFireballFrameDelay(game *Game) {
	for i, _ := range game.player.fireballs {
		game.player.fireballs[i].frameDelay += 1
	}
}
func rewindPlayerFireballAnimation(game *Game, i int) {
	game.player.fireballs[i].column += 1
	game.player.fireballs[i].row = 0
}
func moveIceWall(game *Game) {
	for i, _ := range game.player.iceWalls {
		playerIceWallsDeltaXY(game, i)
	}
}
func playerIceWallsDeltaXY(game *Game, i int) {
	game.player.iceWalls[i].xLoc += math.Cos(game.player.iceWalls[i].trajectory) * 5
	game.player.iceWalls[i].yLoc += math.Sin(game.player.iceWalls[i].trajectory) * 5
}
func animateIceWall(game *Game) {
	for i, _ := range game.player.iceWalls {
		incrementPlayerIceWallFrameDelay(game)
		if i < len(game.player.iceWalls) && game.player.iceWalls[i].frameDelay%2 == 0 {
			game.player.iceWalls[i].row += 1
			if game.player.iceWalls[i].row >= ICE_FRAMES_PER_SHEET { // frames per row
				game.player.iceWalls[i].column += 1
				game.player.iceWalls[i].row = 0
				if game.player.iceWalls[i].column > 5 {
					removeIceWall(game, i)
				}
			}
		}
	}
}

func removeIceWall(game *Game, i int) {
	game.player.iceWalls[i] = game.player.iceWalls[len(game.player.iceWalls)-1]
	game.player.iceWalls = game.player.iceWalls[:len(game.player.iceWalls)-1]
}

func incrementPlayerIceWallFrameDelay(game *Game) {
	for i, _ := range game.player.iceWalls {
		game.player.iceWalls[i].frameDelay += 1
	}
}
func removeFireball(game *Game, i int) {
	game.player.fireballs[i] = game.player.fireballs[len(game.player.fireballs)-1]
	game.player.fireballs = game.player.fireballs[:len(game.player.fireballs)-1]
}
func animatePlayerNuke(game *Game) {
	for i, _ := range game.player.nukes {
		incrementPlayerNukeFrameDelay(game)
		if game.player.nukes[i].frameDelay%EXPLOSION_FRAMES_PER_SHEET == 0 {
			incrementPlayerNukeRow(game, i)
			if game.player.nukes[i].row >= EXPLOSION_FRAMES_PER_SHEET {
				incrementPlayerNukeColumn(game, i)
				if game.player.nukes[i].column > 4 {
					removeNuke(game, i)
				}
			}
		}
	}
}
func incrementPlayerNukeFrameDelay(game *Game) {
	for i, _ := range game.player.nukes {
		game.player.nukes[i].frameDelay += 1
	}
}
func incrementPlayerNukeRow(game *Game, i int) {
	game.player.nukes[i].row += 1
}
func incrementPlayerNukeColumn(game *Game, i int) {
	game.player.nukes[i].column += 1
	game.player.nukes[i].row = 0
}
func removeNuke(game *Game, i int) {
	game.player.nukes[i] = game.player.nukes[len(game.player.nukes)-1]
	game.player.nukes = game.player.nukes[:len(game.player.nukes)-1]
}
