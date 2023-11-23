package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"math"
	"time"
)

func getPlayerInput(game *Game) error {

	if inpututil.IsKeyJustPressed(ebiten.KeyW) { // NORTH
		game.direction.moveNorth = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyW) {
		game.direction.moveNorth = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) { // SOUTH
		game.direction.moveSouth = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyS) {
		game.direction.moveSouth = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyD) { // EAST
		game.direction.moveEast = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyD) {
		game.direction.moveEast = false
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyA) { // WEST
		game.direction.moveWest = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyA) {
		game.direction.moveWest = false
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if time.Since(game.gameTimers.fireballTimer) >= FIREBALL_SHOT_CLOCK {
			mouseX, mouseY := ebiten.CursorPosition()
			game.player.fireballs = append(game.player.fireballs,
				setupFireballAttack(mouseX, mouseY, *game))
			playAvidiKidiviSound(game)
			game.gameTimers.fireballTimer = time.Now()
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		if time.Since(game.gameTimers.nukeTimer) >= NUKE_SHOT_CLOCK {
			game.player.nukes = append(game.player.nukes,
				setupNukeAttack(*game))
			playAvadaKedavraSound(game)
			game.gameTimers.nukeTimer = time.Now()
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		if time.Since(game.gameTimers.teleportTimer) >= TELEPORT_SHOT_CLOCK {
			mouseX, mouseY := game.cursor.xLoc, game.cursor.yLoc
			if 70 < mouseX && mouseX < 890 && 70 < mouseY && mouseY < 890 {
				playTeleportSound(game)
				disappearPlayer(game)
				go reappearPlayer(game, mouseX, mouseY)
				game.gameTimers.teleportTimer = time.Now()
			}
		}
	}
	return nil
}

func reappearPlayer(game *Game, mouseX float64, mouseY float64) {

	time.Sleep(time.Second)
	game.player.frame = 0
	game.player.direction = REAPPEARING
	game.player.xLoc = mouseX
	game.player.yLoc = mouseY
	time.Sleep(time.Second)
	game.player.canMove = true
}

func disappearPlayer(game *Game) {
	game.player.canMove = false
	game.player.frame = 0
	game.player.direction = DISAPPEARING
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
