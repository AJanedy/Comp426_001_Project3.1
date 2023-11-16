package main

import (
	"github.com/co0p/tankism/lib/collision"
	"math"
)

func checkProximityAndCollisions(game *Game) {

	//checkArmorPotCollision(game)
	//checkHealthPotCollision(game)
	//checkManaPotCollision(game)
	checkShotsTakenAtEnemy1(game)
	checkEnemy1ShotsFired(game)
	checkPlayerEnemy1Proximity(game)
}
func checkShotsTakenAtEnemy1(game *Game) {
	for i, _ := range game.enemies.enemy1 {
		for j, _ := range game.player.fireballs {
			if isEnemy1Hit(game, i, j) {
				playOofSound(game)
				removeFireball(game, j)
			}
		}
	}
}
func isEnemy1Hit(game *Game, i int, j int) bool {
	if i < len(game.enemies.enemy1) && j < len(game.player.fireballs) {
		attackBounds := collision.BoundingBox{
			X:      float64(game.player.fireballs[j].xLoc),
			Y:      float64(game.player.fireballs[j].yLoc),
			Width:  float64(PLAYER_FIREBALL_FRAME_WIDTH - 15),
			Height: float64(PLAYER_FIREBALL_FRAME_HEIGHT - 15),
		}
		enemyBounds := collision.BoundingBox{
			X:      float64(game.enemies.enemy1[i].xLoc),
			Y:      float64(game.enemies.enemy1[i].yLoc),
			Width:  float64(PLAYER_FRAME_WIDTH - 25),
			Height: float64(PLAYER_FRAME_HEIGHT - 25),
		}
		if collision.AABBCollision(attackBounds, enemyBounds) {
			game.enemies.enemy1[i].health -= game.player.fireballs[j].damage
			return true
		}
	}
	return false
}
func checkEnemy1ShotsFired(game *Game) {
	for i, _ := range game.enemies.enemy1 {
		for j, _ := range game.enemies.enemy1[i].attacks {
			if isPlayer1HitByEnemy1(game, i, j) {
				playOwSound(game)
				lowerPlayerHealth(game, game.enemies.enemy1[i].attackPower)
				removeSpentEnemy1Fireball(game, i, j)
			}
		}
	}
}
func isPlayer1HitByEnemy1(game *Game, i int, j int) bool {

	if i < len(game.enemies.enemy1) && j < len(game.enemies.enemy1[i].attacks) {
		attackBounds := collision.BoundingBox{
			X:      float64(game.enemies.enemy1[i].attacks[j].xLoc),
			Y:      float64(game.enemies.enemy1[i].attacks[j].yLoc),
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
func checkPlayerEnemy1Proximity(game *Game) {
	for i, _ := range game.enemies.enemy1 {
		playerX, playerY := game.player.xLoc, game.player.yLoc
		enemyX, enemyY := game.enemies.enemy1[i].xLoc, game.enemies.enemy1[i].yLoc
		distance := math.Sqrt(math.Pow(float64(playerX-enemyX), 2) + math.Pow(float64(playerY-enemyY), 2))
		if 200 <= int(distance) && int(distance) <= game.enemies.enemy1[i].distanceThreshold {
			game.enemies.enemy1[i].isChasing = true
		} else if int(distance) < 200 {
			game.enemies.enemy1[i].isInRange = true
			game.enemies.enemy1[i].isChasing = false
		} else {
			game.enemies.enemy1[i].isInRange = false
			game.enemies.enemy1[i].isChasing = false
		}
	}
}

//func checkArmorPotCollision(game *Game) bool {
//	playerBounds := collision.BoundingBox{
//		X:      float64(game.player.xLoc),
//		Y:      float64(game.player.yLoc),
//		Width:  float64(PLAYER_FRAME_WIDTH),
//		Height: float64(PLAYER_FRAME_HEIGHT),
//	}
//	barrierBounds := collision.BoundingBox{
//		X:      float64(game.powerUps.armorPot.xLoc),
//		Y:      float64(game.powerUps.armorPot.yLoc),
//		Width:  float64(10),
//		Height: float64(10),
//	}
//	if collision.AABBCollision(playerBounds, barrierBounds) {
//		//game.player.armor += 1
//		return true
//	}
//	return false
//}
//func checkHealthPotCollision(game *Game) bool {
//	playerBounds := collision.BoundingBox{
//		X:      float64(game.player.xLoc),
//		Y:      float64(game.player.yLoc),
//		Width:  float64(PLAYER_FRAME_WIDTH),
//		Height: float64(PLAYER_FRAME_HEIGHT),
//	}
//	barrierBounds := collision.BoundingBox{
//		X:      float64(game.powerUps.healthPot.xLoc),
//		Y:      float64(game.powerUps.healthPot.yLoc),
//		Width:  float64(10),
//		Height: float64(10),
//	}
//	if collision.AABBCollision(playerBounds, barrierBounds) {
//		//game.player.health += 1
//		return true
//	}
//	return false
//}
//func checkManaPotCollision(game *Game) bool {
//	playerBounds := collision.BoundingBox{
//		X:      float64(game.player.xLoc),
//		Y:      float64(game.player.yLoc),
//		Width:  float64(PLAYER_FRAME_WIDTH),
//		Height: float64(PLAYER_FRAME_HEIGHT),
//	}
//	barrierBounds := collision.BoundingBox{
//		X:      float64(game.powerUps.manaPot.xLoc),
//		Y:      float64(game.powerUps.manaPot.yLoc),
//		Width:  float64(10),
//		Height: float64(10),
//	}
//	if collision.AABBCollision(playerBounds, barrierBounds) {
//		//game.player.magicPower += 1
//		return true
//	}
//	return false
//}
