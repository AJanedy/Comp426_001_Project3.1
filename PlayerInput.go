package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
		if time.Since(game.gameTimers.attack1Timer) >= LEFT_CLICK_SHOT_CLOCK {
			mouseX, mouseY := ebiten.CursorPosition()
			game.player.fireballs = append(game.player.fireballs,
				setupShootingAttack(mouseX, mouseY, *game))
			playAvidiKidiviSound(game)
			game.gameTimers.attack1Timer = time.Now()
		}
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		if time.Since(game.gameTimers.attack2Timer) >= RIGHT_CLICK_SHOT_CLOCK {
			game.player.nukes = append(game.player.nukes,
				setupStationaryAttack(*game))
			game.gameTimers.attack2Timer = time.Now()
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		mouseX, mouseY := ebiten.CursorPosition()
		if time.Since(game.gameTimers.qAttackTimer) >= Q_ATTACK_SHOT_CLOCK {
			game.player.iceWalls = append(game.player.iceWalls,
				setupShootingAttack(mouseX, mouseY, *game))
			game.gameTimers.qAttackTimer = time.Now()
		}
	}
	return nil
}
