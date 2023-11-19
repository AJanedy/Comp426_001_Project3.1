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
	//drawLevelTwo(game, screen)
	drawLevelThree(game, screen)
	drawPlayerAnimations(game, screen)
	drawHud(game, screen)
	//drawEnemy1Animations(game, screen)
	//drawEnemy2Animations(game, screen)
}
func drawLevelTwo(game Game, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	for tileY := 0; tileY < game.maps.level2.level.Height; tileY += 1 {
		for tileX := 0; tileX < game.maps.level2.level.Width; tileX += 1 {
			drawOptions.GeoM.Reset()
			TileXpos := float64(game.maps.level2.level.TileWidth * tileX)
			TileYpos := float64(game.maps.level2.level.TileHeight * tileY)
			drawOptions.GeoM.Translate(TileXpos, TileYpos)
			tileToDraw :=
				game.maps.level2.level.Layers[0].Tiles[tileY*game.maps.level2.level.Width+tileX]
			ebitenTileToDraw := game.maps.level2.tileHash[tileToDraw.ID]
			screen.DrawImage(ebitenTileToDraw,
				drawOptions)
		}
	}
	for tileY := 0; tileY < game.maps.level2.level.Height; tileY += 1 {
		for tileX := 0; tileX < game.maps.level2.level.Width; tileX += 1 {
			drawOptions.GeoM.Reset()
			TileXpos := float64(game.maps.level2.level.TileWidth * tileX)
			TileYpos := float64(game.maps.level2.level.TileHeight * tileY)
			drawOptions.GeoM.Translate(TileXpos, TileYpos)
			tileToDraw :=
				game.maps.level2.level.Layers[1].Tiles[tileY*game.maps.level2.level.Width+tileX]
			ebitenTileToDraw := game.maps.level2.tileHash[tileToDraw.ID]
			screen.DrawImage(ebitenTileToDraw,
				drawOptions)
		}
	}
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
func drawEnemy2Animations(game Game, screen *ebiten.Image) {
	drawLevel2Enemies(game, screen)
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
	ebitenutil.DebugPrintAt(screen, "Health: "+strconv.Itoa(game.player.health), 70, 770)
	ebitenutil.DebugPrintAt(screen, "Armor: "+strconv.Itoa(game.player.armor), 70, 790)
	ebitenutil.DebugPrintAt(screen, "Magic Power: "+strconv.Itoa(game.player.magicPower), 70, 810)
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
	ebitenutil.DebugPrintAt(screen, "Fireball                  LClick", 70, 830)
	ebitenutil.DrawRect(screen, 145, 835, float64(barWidth),
		float64(barHeight), color.Gray{192})
	ebitenutil.DrawRect(screen, 145, 835, float64(barWidth)*fireBallCoolDown,
		float64(barHeight), color.RGBA{255, 0, 0, 255})

	ebitenutil.DebugPrintAt(screen, "Electricity               RClick", 70, 850)
	ebitenutil.DrawRect(screen, 145, 855, float64(barWidth),
		float64(barHeight), color.Gray{192})
	ebitenutil.DrawRect(screen, 145, 855, float64(barWidth)*electricityCoolDown,
		float64(barHeight), color.RGBA{255, 0, 0, 255})

	ebitenutil.DebugPrintAt(screen, "Nuke                      Q", 70, 870)
	ebitenutil.DrawRect(screen, 145, 875, float64(barWidth),
		float64(barHeight), color.Gray{192})
	ebitenutil.DrawRect(screen, 145, 875, float64(barWidth)*nukeCoolDown,
		float64(barHeight), color.RGBA{255, 0, 0, 255})

	ebitenutil.DebugPrintAt(screen, "Teleport                  E", 70, 890)
	ebitenutil.DrawRect(screen, 145, 895, float64(barWidth),
		float64(barHeight), color.Gray{192})
	ebitenutil.DrawRect(screen, 145, 895, float64(barWidth)*teleportCoolDown,
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
func drawLevel2Enemies(game Game, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}
	drawOptions.GeoM.Reset()
	for i, _ := range game.enemies.bee {
		drawOptions.GeoM.Reset()
		drawOptions.GeoM.Translate(game.enemies.bee[i].xLoc, game.enemies.bee[i].yLoc)
		screen.DrawImage(game.enemies.bee[i].enemySprite.SubImage(image.Rect(
			game.enemies.bee[i].frame*BEE_FRAME_WIDTH,
			game.enemies.bee[i].direction*BEE_FRAME_HEIGHT,
			game.enemies.bee[i].frame*BEE_FRAME_WIDTH+BEE_FRAME_WIDTH,
			game.enemies.bee[i].direction*BEE_FRAME_HEIGHT+BEE_FRAME_HEIGHT)).(*ebiten.Image),
			drawOptions)
	}
}
func drawLevelOneEnemiesAttack(game Game, screen *ebiten.Image) {
	drawOptions := &ebiten.DrawImageOptions{}

	for i, _ := range game.enemies.rainbowMan {
		for j, _ := range game.enemies.rainbowMan[i].attacks {
			drawOptions.GeoM.Reset()
			drawOptions.GeoM.Translate(game.enemies.rainbowMan[i].attacks[j].xLoc,
				game.enemies.rainbowMan[i].attacks[j].yLoc)
			screen.DrawImage(game.animations.enemy1Attack.animation.SubImage(image.Rect(
				game.enemies.rainbowMan[i].attacks[j].row*ENEMY_ATTACK_FRAME_WIDTH,
				game.enemies.rainbowMan[i].attacks[j].column*ENEMY_ATTACK_FRAME_HEIGHT,
				game.enemies.rainbowMan[i].attacks[j].row*ENEMY_ATTACK_FRAME_WIDTH+ENEMY_ATTACK_FRAME_WIDTH,
				game.enemies.rainbowMan[i].attacks[j].column*ENEMY_ATTACK_FRAME_HEIGHT+ENEMY_ATTACK_FRAME_HEIGHT)).(*ebiten.Image),
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
