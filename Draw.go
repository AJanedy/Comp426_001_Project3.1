package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	"image/color"
	"strconv"
	"time"
)

func (game Game) Draw(screen *ebiten.Image) {
	drawPlayerAnimations(game, screen)
	drawHud(game, screen)
	drawEnemy1Animations(game, screen)
}
func drawPlayerAnimations(game Game, screen *ebiten.Image) {
	drawPlayer(game, screen)
	drawFireBalls(game, screen)
	drawNukes(game, screen)
	drawIceWalls(game, screen)
}
func drawHud(game Game, screen *ebiten.Image) {
	drawHealthBar(game, screen)
	drawPlayerAttributes(game, screen)
	drawCoolDowns(game, screen)
	drawCursor(game, screen)
	drawWASD(game, screen)
}
func drawEnemy1Animations(game Game, screen *ebiten.Image) {
	drawLevelOneEnemies(game, screen)
	drawLevelOneEnemiesAttack(game, screen)
}
func drawPlayer(game Game, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	drawOptions.GeoM.Reset()
	drawOptions.GeoM.Translate(float64(game.player.xLoc), float64(game.player.yLoc))
	screen.DrawImage(game.player.playerSprite.SubImage(image.Rect(
		game.player.frame*PLAYER_FRAME_WIDTH,
		game.player.direction*PLAYER_FRAME_HEIGHT,
		game.player.frame*PLAYER_FRAME_WIDTH+PLAYER_FRAME_WIDTH,
		game.player.direction*PLAYER_FRAME_HEIGHT+PLAYER_FRAME_HEIGHT)).(*ebiten.Image),
		drawOptions)
}
func drawCursor(game Game, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	drawOptions.GeoM.Reset()
	drawOptions.GeoM.Translate(game.cursor.xLoc, game.cursor.yLoc)
	screen.DrawImage(game.cursor.cursor, drawOptions)
}
func drawFireBalls(game Game, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	for i, _ := range game.player.fireballs {
		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(game.player.fireballs[i].xLoc, game.player.fireballs[i].yLoc)
		screen.DrawImage(game.animations.fireball.animation.SubImage(image.Rect(
			game.player.fireballs[i].row*PLAYER_FIREBALL_FRAME_WIDTH,
			game.player.fireballs[i].column*PLAYER_FIREBALL_FRAME_HEIGHT,
			game.player.fireballs[i].row*PLAYER_FIREBALL_FRAME_WIDTH+PLAYER_FIREBALL_FRAME_WIDTH,
			game.player.fireballs[i].column*PLAYER_FIREBALL_FRAME_HEIGHT+PLAYER_FIREBALL_FRAME_HEIGHT)).(*ebiten.Image),
			drawOptions)
	}
}
func drawIceWalls(game Game, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	for i, _ := range game.player.iceWalls {
		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(game.player.iceWalls[i].xLoc, game.player.iceWalls[i].yLoc)
		screen.DrawImage(game.animations.iceWall.animation.SubImage(image.Rect(
			game.player.iceWalls[i].row*ENEMY_ATTACK_FRAME_WIDTH,
			game.player.iceWalls[i].column*ENEMY_ATTACK_FRAME_HEIGHT,
			game.player.iceWalls[i].row*ENEMY_ATTACK_FRAME_WIDTH+ENEMY_ATTACK_FRAME_WIDTH,
			game.player.iceWalls[i].column*ENEMY_ATTACK_FRAME_HEIGHT+ENEMY_ATTACK_FRAME_HEIGHT)).(*ebiten.Image),
			drawOptions)
	}
}
func drawHealthBar(game Game, screen *ebiten.Image) {
	barWidth := 30.0
	barHeight := 2
	barX := game.player.xLoc
	barY := game.player.yLoc + 55
	currentBarWidth := barWidth * float64(game.player.health) / float64(game.player.maxHealth)

	ebitenutil.DrawRect(screen, float64(barX), float64(barY),
		float64(barWidth), float64(barHeight), color.Gray{192})
	ebitenutil.DrawRect(screen, float64(barX), float64(barY),
		float64(currentBarWidth), float64(barHeight), color.RGBA{255, 0, 0, 255})
}
func drawPlayerAttributes(game Game, screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "Health: "+strconv.Itoa(game.player.health), 100, 800)
	ebitenutil.DebugPrintAt(screen, "Armor: "+strconv.Itoa(game.player.armor), 100, 820)
	ebitenutil.DebugPrintAt(screen, "Magic Power: "+strconv.Itoa(game.player.magicPower), 100, 840)
}
func drawCoolDowns(game Game, screen *ebiten.Image) {
	fireBallCoolDown := time.Since(game.gameTimers.fireballTimer).Seconds() / 3
	electricityCoolDown := time.Since(game.gameTimers.electricityTimer).Seconds() / 5
	nukeCoolDown := time.Since(game.gameTimers.nukeTimer).Seconds() / 15
	teleportCoolDown := time.Since(game.gameTimers.teleportTimer).Seconds() / 10

	barWidth := 74
	barHeight := 4

	if fireBallCoolDown > 1.0 {
		fireBallCoolDown = 1
	}
	if electricityCoolDown > 1.0 {
		electricityCoolDown = 1
	}
	if nukeCoolDown > 1.0 {
		nukeCoolDown = 1
	}
	if teleportCoolDown > 1.0 {
		teleportCoolDown = 1
	}
	ebitenutil.DebugPrintAt(screen, "Fireball", 100, 860)
	ebitenutil.DrawRect(screen, 175, 865, float64(barWidth),
		float64(barHeight), color.Gray{192})
	ebitenutil.DrawRect(screen, 175, 865, float64(barWidth)*fireBallCoolDown,
		float64(barHeight), color.RGBA{255, 0, 0, 255})

	ebitenutil.DebugPrintAt(screen, "Electricity", 100, 880)
	ebitenutil.DrawRect(screen, 175, 885, float64(barWidth),
		float64(barHeight), color.Gray{192})
	ebitenutil.DrawRect(screen, 175, 885, float64(barWidth)*electricityCoolDown,
		float64(barHeight), color.RGBA{255, 0, 0, 255})

	ebitenutil.DebugPrintAt(screen, "Nuke", 100, 900)
	ebitenutil.DrawRect(screen, 175, 905, float64(barWidth),
		float64(barHeight), color.Gray{192})
	ebitenutil.DrawRect(screen, 175, 905, float64(barWidth)*nukeCoolDown,
		float64(barHeight), color.RGBA{255, 0, 0, 255})

	ebitenutil.DebugPrintAt(screen, "Teleport", 100, 920)
	ebitenutil.DrawRect(screen, 175, 925, float64(barWidth),
		float64(barHeight), color.Gray{192})
	ebitenutil.DrawRect(screen, 175, 925, float64(barWidth)*teleportCoolDown,
		float64(barHeight), color.RGBA{255, 0, 0, 255})
}
func drawWASD(game Game, screen *ebiten.Image) {
	drawOption := &ebiten.DrawImageOptions{}
	drawOption.GeoM.Reset()
	drawOption.GeoM.Translate(75, 50)
	screen.DrawImage(game.gameHUD.WASD, drawOption)
	ebitenutil.DebugPrintAt(screen, "Movement", 105, 150)
}
func drawLevelOneEnemies(game Game, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	drawOptions.GeoM.Reset()
	for i, _ := range game.enemies.enemy1 {
		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(game.enemies.enemy1[i].xLoc, game.enemies.enemy1[i].yLoc)
		screen.DrawImage(game.enemies.enemy1[i].enemySprite.SubImage(image.Rect(
			game.enemies.enemy1[i].frame*PLAYER_FRAME_WIDTH,
			game.enemies.enemy1[i].direction*PLAYER_FRAME_HEIGHT,
			game.enemies.enemy1[i].frame*PLAYER_FRAME_WIDTH+PLAYER_FRAME_WIDTH,
			game.enemies.enemy1[i].direction*PLAYER_FRAME_HEIGHT+PLAYER_FRAME_HEIGHT)).(*ebiten.Image),
			drawOptions)
	}
}
func drawLevelOneEnemiesAttack(game Game, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}

	for i, _ := range game.enemies.enemy1 {
		for j, _ := range game.enemies.enemy1[i].attacks {
			drawOptions.GeoM.Reset()
			drawOptions.GeoM.Translate(game.enemies.enemy1[i].attacks[j].xLoc,
				game.enemies.enemy1[i].attacks[j].yLoc)
			screen.DrawImage(game.animations.enemy1Attack.animation.SubImage(image.Rect(
				game.enemies.enemy1[i].attacks[j].row*ENEMY_ATTACK_FRAME_WIDTH,
				game.enemies.enemy1[i].attacks[j].column*ENEMY_ATTACK_FRAME_HEIGHT,
				game.enemies.enemy1[i].attacks[j].row*ENEMY_ATTACK_FRAME_WIDTH+ENEMY_ATTACK_FRAME_WIDTH,
				game.enemies.enemy1[i].attacks[j].column*ENEMY_ATTACK_FRAME_HEIGHT+ENEMY_ATTACK_FRAME_HEIGHT)).(*ebiten.Image),
				drawOptions)
		}
	}
}
func drawNukes(game Game, screen *ebiten.Image) {
	drawOption := &ebiten.DrawImageOptions{}
	for i, _ := range game.player.nukes {
		drawOption.GeoM.Reset()
		drawOption.GeoM.Translate(float64(game.player.nukes[i].xLoc-35),
			float64(game.player.nukes[i].yLoc-50))
		screen.DrawImage(game.animations.explosion.animation.SubImage(image.Rect(
			game.player.nukes[i].row*EXPLOSION_FRAME_WIDTH,
			game.player.nukes[i].column*EXPLOSION_FRAME_HEIGHT,
			game.player.nukes[i].row*EXPLOSION_FRAME_WIDTH+EXPLOSION_FRAME_WIDTH,
			game.player.nukes[i].column*EXPLOSION_FRAME_HEIGHT+EXPLOSION_FRAME_HEIGHT)).(*ebiten.Image),
			drawOption)
	}
}
