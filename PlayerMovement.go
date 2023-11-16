package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

func movePlayer(game *Game) error {

	if game.direction.moveEast || game.direction.moveWest || game.direction.moveSouth || game.direction.moveNorth {

		game.player.frameDelay += 1

		if game.player.frameDelay%FRAMES_PER_SHEET == 0 {
			game.player.frame += 1
			if game.player.frame >= FRAMES_PER_SHEET {
				game.player.frame = 0
			}
		}
	}
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
	}

	return nil
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
func limitCursorDistanceFromPlayer(game *Game) {
	maxDistance := 380
	mouseX, mouseY := ebiten.CursorPosition()
	dX := mouseX - int(game.player.xLoc)
	dY := mouseY - int(game.player.yLoc)
	distance := math.Sqrt(float64(dX*dX + dY*dY))
	game.cursor.xLoc = float64(mouseX)
	game.cursor.yLoc = float64(mouseY)
	if distance > float64(maxDistance) {
		angle := math.Atan2(float64(dY), float64(dX))
		game.cursor.xLoc = game.player.xLoc + float64(maxDistance)*math.Cos(angle)
		game.cursor.yLoc = game.player.yLoc + float64(maxDistance)*math.Sin(angle)
	}
}
