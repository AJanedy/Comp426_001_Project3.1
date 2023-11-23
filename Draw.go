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
	if game.level == 1 {
		drawLevelThree(game, screen)
		if game.player.getMessage == true {
			drawQuestMessage(game, screen)
		}
		drawHud(game, screen)
		drawWeirdYellowDude(game, screen)
		drawPlayerAnimations(game, screen)
		drawRainbowManAnimations(game, screen)
	}
	if game.level == 2 {
		drawLevelThree(game, screen)
		drawHud(game, screen)
		drawPlayerAnimations(game, screen)
		drawBeeAnimations(game, screen)
	}
	if game.level == 3 {
		drawLevelThree(game, screen)
		drawHud(game, screen)
		drawSmileyMan(game, screen)
		drawPlayerAnimations(game, screen)
	}
	if game.level == 0 {
		drawGameOver(game, screen)
	}
	if game.level == 4 {
		drawVictory(game, screen)
	}
	drawHearts(game, screen)
	drawSpeedBoost(game, screen)
	drawManaPot(game, screen)
	drawKey(game, screen)
}
func drawGameOver(game Game, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	drawOptions.GeoM.Reset()
	drawOptions.GeoM.Translate(180, 200)
	screen.DrawImage(game.animations.gameOver.animation, drawOptions)
}
func drawVictory(game Game, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	drawOptions.GeoM.Reset()
	drawOptions.GeoM.Translate(0, 300)
	screen.DrawImage(game.animations.victory.animation, drawOptions)
}
func drawLevelThree(game Game, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	for tileY := 0; tileY < game.maps.level3[game.maps.frame].level.Height; tileY += 1 {
		for tileX := 0; tileX < game.maps.level3[game.maps.frame].level.Width; tileX += 1 {
			drawOptions.GeoM.Reset()
			TileXpos := float64(game.maps.level3[game.maps.frame].level.TileWidth * tileX)
			TileYpos := float64(game.maps.level3[game.maps.frame].level.TileHeight * tileY)
			drawOptions.GeoM.Translate(TileXpos, TileYpos)
			tileToDraw :=
				game.maps.level3[game.maps.frame].level.Layers[0].Tiles[tileY*game.maps.level3[game.maps.frame].level.Width+tileX]
			ebitenTileToDraw := game.maps.level3[game.maps.frame].tileHash[tileToDraw.ID]
			screen.DrawImage(ebitenTileToDraw,
				drawOptions)
		}
	}
	for tileY := 0; tileY < game.maps.level3[game.maps.frame].level.Height; tileY += 1 {
		for tileX := 0; tileX < game.maps.level3[game.maps.frame].level.Width; tileX += 1 {
			drawOptions.GeoM.Reset()
			TileXpos := float64(game.maps.level3[game.maps.frame].level.TileWidth * tileX)
			TileYpos := float64(game.maps.level3[game.maps.frame].level.TileHeight * tileY)
			drawOptions.GeoM.Translate(TileXpos, TileYpos)
			tileToDraw :=
				game.maps.level3[game.maps.frame].level.Layers[1].Tiles[tileY*game.maps.level3[game.maps.frame].level.Width+tileX]
			ebitenTileToDraw := game.maps.level3[game.maps.frame].tileHash[tileToDraw.ID]
			screen.DrawImage(ebitenTileToDraw,
				drawOptions)
		}
	}
}
func drawPlayerAnimations(game Game, screen *ebiten.Image) {
	drawPlayer(game, screen)
	drawFireBalls(game, screen)
	drawNukes(game, screen)
}
func drawHud(game Game, screen *ebiten.Image) {
	drawHealthBar(game, screen)
	drawPlayerAttributes(game, screen)
	drawCoolDowns(game, screen)
	drawCursor(game, screen)
}
func drawHearts(game Game, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	for i, _ := range game.powerUps {
		if game.powerUps[i].powerUpType == HEALTH_POT {
			drawOptions.GeoM.Reset()
			drawOptions.GeoM.Translate(float64(game.powerUps[i].xLoc), float64(game.powerUps[i].yLoc))
			screen.DrawImage(game.animations.heart.animation.SubImage(image.Rect(
				game.powerUps[i].row*HEART_FRAME_WIDTH,
				game.powerUps[i].column*HEART_FRAME_HEIGHT,
				game.powerUps[i].row*HEART_FRAME_WIDTH+HEART_FRAME_WIDTH,
				game.powerUps[i].column*HEART_FRAME_HEIGHT+HEART_FRAME_HEIGHT)).(*ebiten.Image),
				drawOptions)
		}
	}
}
func drawSpeedBoost(game Game, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	for i, _ := range game.powerUps {
		if game.powerUps[i].powerUpType == SPEED_POT {
			drawOptions.GeoM.Reset()
			drawOptions.GeoM.Translate(float64(game.powerUps[i].xLoc), float64(game.powerUps[i].yLoc))
			screen.DrawImage(game.animations.speedBoost.animation.SubImage(image.Rect(
				game.powerUps[i].row*SPEED_BOOST_FRAME_WIDTH,
				game.powerUps[i].column*SPEED_BOOST_FRAME_HEIGHT,
				game.powerUps[i].row*SPEED_BOOST_FRAME_WIDTH+SPEED_BOOST_FRAME_WIDTH,
				game.powerUps[i].column*SPEED_BOOST_FRAME_HEIGHT+SPEED_BOOST_FRAME_HEIGHT)).(*ebiten.Image),
				drawOptions)
		}
	}
}
func drawManaPot(game Game, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	for i, _ := range game.powerUps {
		if game.powerUps[i].powerUpType == MANA_POT {
			drawOptions.GeoM.Reset()
			drawOptions.GeoM.Translate(float64(game.powerUps[i].xLoc), float64(game.powerUps[i].yLoc))
			screen.DrawImage(game.animations.manaPot.animation.SubImage(image.Rect(
				game.powerUps[i].row*MANA_POT_FRAME_WIDTH,
				game.powerUps[i].column*MANA_POT_FRAME_HEIGHT,
				game.powerUps[i].row*MANA_POT_FRAME_WIDTH+MANA_POT_FRAME_WIDTH,
				game.powerUps[i].column*MANA_POT_FRAME_HEIGHT+MANA_POT_FRAME_HEIGHT)).(*ebiten.Image),
				drawOptions)
		}
	}
}
func drawKey(game Game, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	for i, _ := range game.powerUps {
		if game.powerUps[i].powerUpType == KEY {
			drawOptions.GeoM.Reset()
			drawOptions.GeoM.Translate(float64(game.powerUps[i].xLoc), float64(game.powerUps[i].yLoc))
			screen.DrawImage(game.animations.key.animation.SubImage(image.Rect(
				game.powerUps[i].row*KEY_FRAME_WIDTH,
				game.powerUps[i].column*KEY_FRAME_HEIGHT,
				game.powerUps[i].row*KEY_FRAME_WIDTH+KEY_FRAME_WIDTH,
				game.powerUps[i].column*KEY_FRAME_HEIGHT+KEY_FRAME_HEIGHT)).(*ebiten.Image),
				drawOptions)
		}
	}
}
func drawRainbowManAnimations(game Game, screen *ebiten.Image) {
	drawRainbowMan(game, screen)
	drawRainbowManAttacks(game, screen)
}
func drawBeeAnimations(game Game, screen *ebiten.Image) {
	drawBees(game, screen)
	drawBeeAttacks(game, screen)
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
func drawWeirdYellowDude(game Game, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	drawOptions.GeoM.Reset()
	drawOptions.GeoM.Translate(150, 150)
	screen.DrawImage(game.animations.questGiver.animation, drawOptions)
}
func drawQuestMessage(game Game, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	drawOptions.GeoM.Reset()
	drawOptions.GeoM.Translate(195, 175)
	screen.DrawImage(game.animations.questInstructions.animation, drawOptions)
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
	ebitenutil.DebugPrintAt(screen, "Health: "+strconv.Itoa(game.player.health), 70, 770)
	ebitenutil.DebugPrintAt(screen, "Move Speed: "+strconv.FormatFloat(
		game.player.moveSpeed, 'f', -1, 64), 70, 790)
	ebitenutil.DebugPrintAt(screen, "Magic Power: "+strconv.FormatFloat(
		game.player.magicPower, 'f', -1, 64), 70, 810)
}
func drawCoolDowns(game Game, screen *ebiten.Image) {
	fireBallCoolDown := time.Since(game.gameTimers.fireballTimer).Seconds() / 3
	nukeCoolDown := time.Since(game.gameTimers.nukeTimer).Seconds() / 15
	teleportCoolDown := time.Since(game.gameTimers.teleportTimer).Seconds() / 10

	barWidth := 74
	barHeight := 4

	if fireBallCoolDown > 1.0 {
		fireBallCoolDown = 1
	}
	if nukeCoolDown > 1.0 {
		nukeCoolDown = 1
	}
	if teleportCoolDown > 1.0 {
		teleportCoolDown = 1
	}
	ebitenutil.DebugPrintAt(screen, "Fireball                  LClick", 70, 830)
	ebitenutil.DrawRect(screen, 145, 835, float64(barWidth),
		float64(barHeight), color.Gray{192})
	ebitenutil.DrawRect(screen, 145, 835, float64(barWidth)*fireBallCoolDown,
		float64(barHeight), color.RGBA{255, 0, 0, 255})

	ebitenutil.DebugPrintAt(screen, "Nuke                      Q", 70, 850)
	ebitenutil.DrawRect(screen, 145, 855, float64(barWidth),
		float64(barHeight), color.Gray{192})
	ebitenutil.DrawRect(screen, 145, 855, float64(barWidth)*nukeCoolDown,
		float64(barHeight), color.RGBA{255, 0, 0, 255})

	ebitenutil.DebugPrintAt(screen, "Teleport                  E", 70, 870)
	ebitenutil.DrawRect(screen, 145, 875, float64(barWidth),
		float64(barHeight), color.Gray{192})
	ebitenutil.DrawRect(screen, 145, 875, float64(barWidth)*teleportCoolDown,
		float64(barHeight), color.RGBA{255, 0, 0, 255})
}
func drawRainbowMan(game Game, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	drawOptions.GeoM.Reset()
	for i, _ := range game.enemies.rainbowMan {
		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(game.enemies.rainbowMan[i].xLoc, game.enemies.rainbowMan[i].yLoc)
		screen.DrawImage(game.enemies.rainbowMan[i].enemySprite.SubImage(image.Rect(
			game.enemies.rainbowMan[i].frame*PLAYER_FRAME_WIDTH,
			game.enemies.rainbowMan[i].direction*PLAYER_FRAME_HEIGHT,
			game.enemies.rainbowMan[i].frame*PLAYER_FRAME_WIDTH+PLAYER_FRAME_WIDTH,
			game.enemies.rainbowMan[i].direction*PLAYER_FRAME_HEIGHT+PLAYER_FRAME_HEIGHT)).(*ebiten.Image),
			drawOptions)
	}
}
func drawSmileyMan(game Game, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	drawOptions.GeoM.Reset()
	drawOptions.GeoM.Translate(game.enemies.smileyMan.xLoc, game.enemies.smileyMan.yLoc)
	screen.DrawImage(game.enemies.smileyMan.enemySprite.SubImage(image.Rect(
		game.enemies.smileyMan.frame*SMILEY_FRAME_WIDTH,
		game.enemies.smileyMan.direction*SMILEY_FRAME_HEIGHT,
		game.enemies.smileyMan.frame*SMILEY_FRAME_WIDTH+SMILEY_FRAME_WIDTH,
		game.enemies.smileyMan.direction*SMILEY_FRAME_HEIGHT+SMILEY_FRAME_HEIGHT)).(*ebiten.Image),
		drawOptions)
}
func drawBees(game Game, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	drawOptions.GeoM.Reset()
	for i, _ := range game.enemies.bees {
		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(game.enemies.bees[i].xLoc, game.enemies.bees[i].yLoc)
		screen.DrawImage(game.enemies.bees[i].enemySprite.SubImage(image.Rect(
			game.enemies.bees[i].frame*BEE_FRAME_WIDTH,
			game.enemies.bees[i].direction*BEE_FRAME_HEIGHT,
			game.enemies.bees[i].frame*BEE_FRAME_WIDTH+BEE_FRAME_WIDTH,
			game.enemies.bees[i].direction*BEE_FRAME_HEIGHT+BEE_FRAME_HEIGHT)).(*ebiten.Image),
			drawOptions)
	}
}
func drawRainbowManAttacks(game Game, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	for i, _ := range game.enemies.rainbowMan {
		for j, _ := range game.enemies.rainbowMan[i].attacks {
			drawOptions.GeoM.Reset()
			drawOptions.GeoM.Translate(game.enemies.rainbowMan[i].attacks[j].xLoc,
				game.enemies.rainbowMan[i].attacks[j].yLoc)
			screen.DrawImage(game.animations.rainbowManAttack.animation.SubImage(image.Rect(
				game.enemies.rainbowMan[i].attacks[j].row*ENEMY_ATTACK_FRAME_WIDTH,
				game.enemies.rainbowMan[i].attacks[j].column*ENEMY_ATTACK_FRAME_HEIGHT,
				game.enemies.rainbowMan[i].attacks[j].row*ENEMY_ATTACK_FRAME_WIDTH+ENEMY_ATTACK_FRAME_WIDTH,
				game.enemies.rainbowMan[i].attacks[j].column*ENEMY_ATTACK_FRAME_HEIGHT+ENEMY_ATTACK_FRAME_HEIGHT)).(*ebiten.Image),
				drawOptions)
		}
	}
}
func drawBeeAttacks(game Game, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	for i, _ := range game.enemies.bees {
		for j, _ := range game.enemies.bees[i].attacks {
			drawOptions.GeoM.Reset()
			drawOptions.GeoM.Translate(game.enemies.bees[i].attacks[j].xLoc,
				game.enemies.bees[i].attacks[j].yLoc)
			screen.DrawImage(game.animations.beeAttack.animation.SubImage(image.Rect(
				game.enemies.bees[i].attacks[j].row*ENEMY_ATTACK_FRAME_WIDTH,
				game.enemies.bees[i].attacks[j].column*ENEMY_ATTACK_FRAME_HEIGHT,
				game.enemies.bees[i].attacks[j].row*ENEMY_ATTACK_FRAME_WIDTH+ENEMY_ATTACK_FRAME_WIDTH,
				game.enemies.bees[i].attacks[j].column*ENEMY_ATTACK_FRAME_HEIGHT+ENEMY_ATTACK_FRAME_HEIGHT)).(*ebiten.Image),
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
