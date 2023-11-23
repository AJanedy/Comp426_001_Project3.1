package main

import (
	"github.com/co0p/tankism/lib/collision"
	"golang.org/x/sys/windows"
	"math"
)

func checkBeesHitByFireballs(game *Game) {
	for i, _ := range game.enemies.bees {
		for j, _ := range game.player.fireballs {
			if isBeeHitByFireball(game, i, j) {
				removeFireball(game, j)
			}
		}
	}
}
func checkRainbowMenHitByFireballs(game *Game) {
	for i, _ := range game.enemies.rainbowMan {
		for j, _ := range game.player.fireballs {
			if isRainbowManHitByFireball(game, i, j) {
				playOofSound(game)
				removeFireball(game, j)
			}
		}
	}
}
func checkRainbowMenHitByNukes(game *Game) {
	for i, _ := range game.enemies.rainbowMan {
		for j, _ := range game.player.nukes {
			isRainbowManHitByNuke(game, i, j)
		}
	}
}
func checkBeesHitByNukes(game *Game) {
	for i, _ := range game.enemies.bees {
		for j, _ := range game.player.nukes {
			isBeeHitByNuke(game, i, j)
		}
	}
}
func isBeeHitByNuke(game *Game, i int, j int) {
	if i < len(game.enemies.bees) && j < len(game.player.nukes) {
		attackBounds := collision.BoundingBox{
			X:      float64(game.player.nukes[j].xLoc),
			Y:      float64(game.player.nukes[j].yLoc),
			Width:  float64(EXPLOSION_FRAME_WIDTH),
			Height: float64(EXPLOSION_FRAME_HEIGHT),
		}
		enemyBounds := collision.BoundingBox{
			X:      float64(game.enemies.bees[i].xLoc),
			Y:      float64(game.enemies.bees[i].yLoc),
			Width:  float64(BEE_FRAME_WIDTH - 20),
			Height: float64(BEE_FRAME_HEIGHT - 20),
		}
		if collision.AABBCollision(attackBounds, enemyBounds) {
			game.enemies.bees[i].health -= float64(NUKE_ATTACK_BASE_DAMAGE) * game.player.magicPower
		}
	}
}
func isRainbowManHitByNuke(game *Game, i int, j int) {
	if i < len(game.enemies.rainbowMan) && j < len(game.player.nukes) {
		attackBounds := collision.BoundingBox{
			X:      float64(game.player.nukes[j].xLoc),
			Y:      float64(game.player.nukes[j].yLoc),
			Width:  float64(EXPLOSION_FRAME_WIDTH),
			Height: float64(EXPLOSION_FRAME_HEIGHT),
		}
		enemyBounds := collision.BoundingBox{
			X:      float64(game.enemies.rainbowMan[i].xLoc),
			Y:      float64(game.enemies.rainbowMan[i].yLoc),
			Width:  float64(PLAYER_FRAME_WIDTH),
			Height: float64(PLAYER_FRAME_HEIGHT),
		}
		if collision.AABBCollision(attackBounds, enemyBounds) {
			game.enemies.rainbowMan[i].health -= float64(NUKE_ATTACK_BASE_DAMAGE) * game.player.magicPower
		}
	}
}
func isBeeHitByFireball(game *Game, i int, j int) bool {
	if i < len(game.enemies.bees) && j < len(game.player.fireballs) {
		attackBounds := collision.BoundingBox{
			X:      float64(game.player.fireballs[j].xLoc),
			Y:      float64(game.player.fireballs[j].yLoc),
			Width:  float64(PLAYER_FIREBALL_FRAME_WIDTH - 15),
			Height: float64(PLAYER_FIREBALL_FRAME_HEIGHT - 15),
		}
		enemyBounds := collision.BoundingBox{
			X:      float64(game.enemies.bees[i].xLoc),
			Y:      float64(game.enemies.bees[i].yLoc),
			Width:  float64(BEE_FRAME_WIDTH),
			Height: float64(BEE_FRAME_HEIGHT),
		}
		if collision.AABBCollision(attackBounds, enemyBounds) {
			game.enemies.bees[i].health -= float64(FIREBALL_BASE_DAMAGE) * game.player.magicPower
			playOofSound(game)
			return true
		}
	}
	return false
}
func isSmileyManHitByFireball(game *Game, i int) bool {
	if i < len(game.player.fireballs) {
		attackBounds := collision.BoundingBox{
			X:      float64(game.player.fireballs[i].xLoc),
			Y:      float64(game.player.fireballs[i].yLoc),
			Width:  float64(PLAYER_FIREBALL_FRAME_WIDTH - 15),
			Height: float64(PLAYER_FIREBALL_FRAME_HEIGHT - 15),
		}
		enemyBounds := collision.BoundingBox{
			X:      float64(game.enemies.smileyMan.xLoc),
			Y:      float64(game.enemies.smileyMan.yLoc),
			Width:  float64(SMILEY_FRAME_WIDTH - 25),
			Height: float64(SMILEY_FRAME_HEIGHT - 25),
		}
		if collision.AABBCollision(attackBounds, enemyBounds) {
			game.enemies.smileyMan.health -= float64(FIREBALL_BASE_DAMAGE) * game.player.magicPower
			removeFireball(game, i)
			playOofSound(game)
			if game.enemies.smileyMan.health <= 0 {
				playVictorySound(game)
			}
			return true
		}
	}
	return false
}
func isSmileyManHitByNuke(game *Game, i int) bool {
	if i < len(game.player.nukes) {
		attackBounds := collision.BoundingBox{
			X:      float64(game.player.nukes[i].xLoc),
			Y:      float64(game.player.nukes[i].yLoc),
			Width:  float64(EXPLOSION_FRAME_WIDTH - 15),
			Height: float64(EXPLOSION_FRAME_HEIGHT - 15),
		}
		enemyBounds := collision.BoundingBox{
			X:      float64(game.enemies.smileyMan.xLoc),
			Y:      float64(game.enemies.smileyMan.yLoc),
			Width:  float64(SMILEY_FRAME_WIDTH),
			Height: float64(SMILEY_FRAME_HEIGHT),
		}
		if collision.AABBCollision(attackBounds, enemyBounds) {
			game.enemies.smileyMan.health -= float64(NUKE_ATTACK_BASE_DAMAGE)
			if game.enemies.smileyMan.health <= 0 {
				playVictorySound(game)
			}
			return true
		}
	}
	return false
}
func didSmileyManCatchPlayer(game *Game) {
	attackBounds := collision.BoundingBox{
		X:      float64(game.player.xLoc),
		Y:      float64(game.player.yLoc),
		Width:  float64(PLAYER_FRAME_WIDTH),
		Height: float64(PLAYER_FRAME_HEIGHT),
	}
	enemyBounds := collision.BoundingBox{
		X:      float64(game.enemies.smileyMan.xLoc),
		Y:      float64(game.enemies.smileyMan.yLoc),
		Width:  float64(SMILEY_FRAME_WIDTH),
		Height: float64(SMILEY_FRAME_HEIGHT),
	}
	if collision.AABBCollision(attackBounds, enemyBounds) {
		lowerPlayerHealth(game, game.player.maxHealth)
	}
}
func isRainbowManHitByFireball(game *Game, i int, j int) bool {
	if i < len(game.enemies.rainbowMan) && j < len(game.player.fireballs) {
		attackBounds := collision.BoundingBox{
			X:      float64(game.player.fireballs[j].xLoc),
			Y:      float64(game.player.fireballs[j].yLoc),
			Width:  float64(PLAYER_FIREBALL_FRAME_WIDTH - 15),
			Height: float64(PLAYER_FIREBALL_FRAME_HEIGHT - 15),
		}
		enemyBounds := collision.BoundingBox{
			X:      float64(game.enemies.rainbowMan[i].xLoc),
			Y:      float64(game.enemies.rainbowMan[i].yLoc),
			Width:  float64(PLAYER_FRAME_WIDTH - 25),
			Height: float64(PLAYER_FRAME_HEIGHT - 25),
		}
		if collision.AABBCollision(attackBounds, enemyBounds) {
			game.enemies.rainbowMan[i].health -= float64(FIREBALL_BASE_DAMAGE) * game.player.magicPower
			return true
		}
	}
	return false
}
func checkRainbowMenShotsFired(game *Game) {
	for i, _ := range game.enemies.rainbowMan {
		for j, _ := range game.enemies.rainbowMan[i].attacks {
			if isPlayerHitByRainbowMan(game, i, j) {
				playOwSound(game)
				lowerPlayerHealth(game, game.enemies.rainbowMan[i].attackPower)
				removeSpentRainbowManFireball(game, i, j)
			}
		}
	}
}
func isPlayerHitByRainbowMan(game *Game, i int, j int) bool {

	if i < len(game.enemies.rainbowMan) && j < len(game.enemies.rainbowMan[i].attacks) {
		attackBounds := collision.BoundingBox{
			X:      float64(game.enemies.rainbowMan[i].attacks[j].xLoc),
			Y:      float64(game.enemies.rainbowMan[i].attacks[j].yLoc),
			Width:  float64(ENEMY_ATTACK_FRAME_WIDTH - 15),
			Height: float64(ENEMY_ATTACK_FRAME_HEIGHT - 15),
		}
		playerBounds := collision.BoundingBox{
			X:      float64(game.player.xLoc),
			Y:      float64(game.player.yLoc),
			Width:  float64(PLAYER_FRAME_WIDTH - 25),
			Height: float64(PLAYER_FRAME_HEIGHT - 25),
		}
		if collision.AABBCollision(attackBounds, playerBounds) {
			return true
		}
	}
	return false
}
func checkBeeShotsFired(game *Game) {
	for i, _ := range game.enemies.bees {
		for j, _ := range game.enemies.bees[i].attacks {
			if isPlayerHitByBee(game, i, j) {
				playOwSound(game)
				lowerPlayerHealth(game, game.enemies.bees[i].attackPower)
				removeSpentBeeAttack(game, i, j)
			}
		}
	}
}
func isPlayerHitByBee(game *Game, i int, j int) bool {

	if i < len(game.enemies.bees) && j < len(game.enemies.bees[i].attacks) {
		attackBounds := collision.BoundingBox{
			X:      float64(game.enemies.bees[i].attacks[j].xLoc),
			Y:      float64(game.enemies.bees[i].attacks[j].yLoc),
			Width:  float64(ENEMY_ATTACK_FRAME_WIDTH - 15),
			Height: float64(ENEMY_ATTACK_FRAME_HEIGHT - 15),
		}
		playerBounds := collision.BoundingBox{
			X:      float64(game.player.xLoc),
			Y:      float64(game.player.yLoc),
			Width:  float64(PLAYER_FRAME_WIDTH - 25),
			Height: float64(PLAYER_FRAME_HEIGHT - 25),
		}
		if collision.AABBCollision(attackBounds, playerBounds) {
			return true
		}
	}
	return false
}
func checkPlayerRainbowMenProximity(game *Game) {
	for i, _ := range game.enemies.rainbowMan {
		playerX, playerY := game.player.xLoc, game.player.yLoc
		enemyX, enemyY := game.enemies.rainbowMan[i].xLoc, game.enemies.rainbowMan[i].yLoc
		distance := math.Sqrt(math.Pow(float64(playerX-enemyX), 2) + math.Pow(float64(playerY-enemyY), 2))
		if 200 <= int(distance) && int(distance) <= game.enemies.rainbowMan[i].distanceThreshold {
			game.enemies.rainbowMan[i].isChasing = true
		} else if int(distance) < 200 {
			game.enemies.rainbowMan[i].isInRange = true
			game.enemies.rainbowMan[i].isChasing = false
		} else {
			game.enemies.rainbowMan[i].isInRange = false
			game.enemies.rainbowMan[i].isChasing = false
		}
	}
}
func checkPlayerQuestGiverProximity(game *Game) {
	playerX, playerY := game.player.xLoc, game.player.yLoc
	questGiverX, questGiverY := float64(150), float64(150)
	distance := math.Sqrt(math.Pow(float64(playerX-questGiverX), 2) + math.Pow(float64(playerY-questGiverY), 2))
	if distance <= 100 {
		game.player.getMessage = true
	} else {
		game.player.getMessage = false
	}
}
func checkPlayerBeeProximity(game *Game) {
	for i, _ := range game.enemies.bees {
		playerX, playerY := game.player.xLoc, game.player.yLoc
		enemyX, enemyY := game.enemies.bees[i].xLoc, game.enemies.bees[i].yLoc
		distance := math.Sqrt(math.Pow(float64(playerX-enemyX), 2) + math.Pow(float64(playerY-enemyY), 2))
		if int(distance) <= game.enemies.bees[i].distanceThreshold-200 {
			game.enemies.bees[i].isInRange = true
		} else {
			game.enemies.bees[i].isInRange = false
		}
	}
}
func checkBarrierCollision(game *Game, i int) bool {
	playerBounds := collision.BoundingBox{
		X:      float64(game.player.xLoc),
		Y:      float64(game.player.yLoc),
		Width:  float64(PLAYER_FRAME_WIDTH),
		Height: float64(PLAYER_FRAME_HEIGHT),
	}
	barrierBounds := collision.BoundingBox{
		X:      float64(game.barrierTiles[i].xLoc),
		Y:      float64(game.barrierTiles[i].yLoc),
		Width:  float64(game.barrierTiles[i].width),
		Height: float64(game.barrierTiles[i].height),
	}
	if collision.AABBCollision(playerBounds, barrierBounds) {
		game.wallLocation = windows.Coord{int16(game.barrierTiles[i].xLoc),
			int16(game.barrierTiles[i].yLoc)}
		return true
	}
	return false
}
func checkPowerUpCollision(game *Game) {
	for i, _ := range game.powerUps {
		if i < len(game.powerUps) {
			playerBounds := collision.BoundingBox{
				X:      float64(game.player.xLoc),
				Y:      float64(game.player.yLoc),
				Width:  float64(PLAYER_FRAME_WIDTH),
				Height: float64(PLAYER_FRAME_HEIGHT),
			}
			powerUpBounds := collision.BoundingBox{
				X:      float64(game.powerUps[i].xLoc),
				Y:      float64(game.powerUps[i].yLoc),
				Width:  float64(25),
				Height: float64(25),
			}
			if collision.AABBCollision(playerBounds, powerUpBounds) {
				updatePlayerAttribute(game, game.powerUps[i].powerUpType)
				removePowerUp(game, i)
				playPowerUpSound(game)
			}
		}
	}
}
