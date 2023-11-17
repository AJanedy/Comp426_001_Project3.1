package main

func movePlayer(game *Game) error {

	if game.direction.moveNorth && game.direction.moveEast {
		moveNorthEast(game)
	} else if game.direction.moveNorth && game.direction.moveWest {
		moveNorthWest(game)
	} else if game.direction.moveSouth && game.direction.moveEast {
		moveSouthEast(game)
	} else if game.direction.moveSouth && game.direction.moveWest {
		moveSouthWest(game)
	} else if game.direction.moveNorth {
		moveNorth(game)
	} else if game.direction.moveEast {
		moveEast(game)
	} else if game.direction.moveSouth {
		moveSouth(game)
	} else if game.direction.moveWest {
		moveWest(game)
	} else {
		standStill(game)
	}

	return nil
}

func animatePlayer(game *Game) {
	game.player.frameDelay += 1

	if game.player.frameDelay%FRAMES_PER_SHEET == 0 {
		game.player.frame += 1
		if game.player.frame >= FRAMES_PER_SHEET {
			game.player.frame = 0
		}
	}
}

func moveNorth(game *Game) {
	game.player.yLoc -= 1
	game.player.direction = NORTH
}
func moveNorthEast(game *Game) {
	game.player.yLoc -= 1
	game.player.xLoc += 1
	game.player.direction = NORTH_EAST
}
func moveEast(game *Game) {
	game.player.xLoc += 1
	game.player.direction = EAST
}
func moveSouthEast(game *Game) {
	game.player.yLoc += 1
	game.player.xLoc += 1
	game.player.direction = SOUTH_EAST
}
func moveSouth(game *Game) {
	game.player.yLoc += 1
	game.player.direction = SOUTH
}
func moveSouthWest(game *Game) {
	game.player.yLoc += 1
	game.player.xLoc -= 1
	game.player.direction = SOUTH_WEST
}
func moveWest(game *Game) {
	game.player.xLoc -= 1
	game.player.direction = WEST
}
func moveNorthWest(game *Game) {
	game.player.yLoc -= 1
	game.player.xLoc -= 1
	game.player.direction = NORTH_WEST
}
func standStill(game *Game) {
	game.player.direction = STAND_STILL
}
