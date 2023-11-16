package main

import (
	"math"
	"time"
)

func moveLevelOneEnemy(game *Game) {
	for i, _ := range game.enemies.enemy1 {
		if game.enemies.enemy1[i].isChasing {
			getEnemyMovementDirection(game)
			enemy1DeltaXY(game, i)
		}
	}
}
func getEnemyMovementDirection(game *Game) {
	for i, _ := range game.enemies.enemy1 {
		game.enemies.enemy1[i].degreesDirection = math.Atan2(float64(game.player.yLoc-game.enemies.enemy1[i].yLoc),
			float64(game.player.xLoc-game.enemies.enemy1[i].xLoc))
		game.enemies.enemy1[i].radiansDirection = game.enemies.enemy1[i].degreesDirection * 180 / math.Pi
		translateRadiansToCardinalDirection(game, i)
	}
}
func enemy1DeltaXY(game *Game, i int) {
	game.enemies.enemy1[i].xLoc += math.Cos(game.enemies.enemy1[i].degreesDirection) / 2
	game.enemies.enemy1[i].yLoc += math.Sin(game.enemies.enemy1[i].degreesDirection) / 2
}
func translateRadiansToCardinalDirection(game *Game, i int) {
	if -157.5 <= game.enemies.enemy1[i].radiansDirection && game.enemies.enemy1[i].radiansDirection < -112.5 {
		game.enemies.enemy1[i].direction = NORTH_WEST
	} else if -112.5 <= game.enemies.enemy1[i].radiansDirection && game.enemies.enemy1[i].radiansDirection < -67.5 {
		game.enemies.enemy1[i].direction = NORTH
	} else if -67.5 <= game.enemies.enemy1[i].radiansDirection && game.enemies.enemy1[i].radiansDirection < -22.5 {
		game.enemies.enemy1[i].direction = NORTH_EAST
	} else if -22.5 <= game.enemies.enemy1[i].radiansDirection && game.enemies.enemy1[i].radiansDirection < 22.5 {
		game.enemies.enemy1[i].direction = EAST
	} else if 22.5 <= game.enemies.enemy1[i].radiansDirection && game.enemies.enemy1[i].radiansDirection < 67.5 {
		game.enemies.enemy1[i].direction = SOUTH_EAST
	} else if 67.5 <= game.enemies.enemy1[i].radiansDirection && game.enemies.enemy1[i].radiansDirection < 112.5 {
		game.enemies.enemy1[i].direction = SOUTH
	} else if 112.5 <= game.enemies.enemy1[i].radiansDirection && game.enemies.enemy1[i].radiansDirection < 157.5 {
		game.enemies.enemy1[i].direction = SOUTH_WEST
	} else if game.enemies.enemy1[i].radiansDirection >= 157.5 || game.enemies.enemy1[i].radiansDirection <= -157.5 {
		game.enemies.enemy1[i].direction = WEST
	}
}
func animateEnemy1(game *Game) {
	for i, _ := range game.enemies.enemy1 {
		if game.enemies.enemy1[i].isChasing {
			incrementEnemy1FrameDelay(game, i)
			if game.enemies.enemy1[i].frameDelay%FRAMES_PER_SHEET == 0 {
				incrementEnemy1Frame(game, i)
				if game.enemies.enemy1[i].frame >= FRAMES_PER_SHEET {
					resetEnemyFrame(game, i)
				}
			}
		}
	}
}
func incrementEnemy1FrameDelay(game *Game, i int) {
	game.enemies.enemy1[i].frameDelay += 1
}
func incrementEnemy1Frame(game *Game, i int) {
	game.enemies.enemy1[i].frame += 1
}
func resetEnemyFrame(game *Game, i int) {
	game.enemies.enemy1[i].frame = 0
}
func enemy1Attackers(game *Game) {
	for i, _ := range game.enemies.enemy1 {
		checkEnemyFiringRange(game, i)
	}
}
func checkEnemyFiringRange(game *Game, i int) {
	if game.enemies.enemy1[i].isInRange {
		if time.Since(game.enemies.enemy1[i].shotTimer) >= LEFT_CLICK_SHOT_CLOCK {
			enemy1Attack(game, i)
			resetEnemy1AttackCooldown(game, i)
			playPewPewSound(game)
		}
	}
}
func enemy1Attack(game *Game, i int) {
	fireballTrajectory := math.Atan2(float64(game.player.yLoc)-game.enemies.enemy1[i].yLoc,
		float64(game.player.xLoc)-game.enemies.enemy1[i].xLoc)
	game.enemies.enemy1[i].attacks = append(game.enemies.enemy1[i].attacks,
		setupEnemyAttack(fireballTrajectory, *game, i))
}
func resetEnemy1AttackCooldown(game *Game, i int) {
	game.enemies.enemy1[i].shotTimer = time.Now()
}
func moveEnemy1Fireballs(game *Game) {
	for i, _ := range game.enemies.enemy1 {
		for j, _ := range game.enemies.enemy1[i].attacks {
			game.enemies.enemy1[i].attacks[j].xLoc +=
				math.Cos(game.enemies.enemy1[i].attacks[j].trajectory) * 3
			game.enemies.enemy1[i].attacks[j].yLoc +=
				math.Sin(game.enemies.enemy1[i].attacks[j].trajectory) * 3
		}
	}
}
func animateEnemy1Fireball(game *Game) {
	for i, _ := range game.enemies.enemy1 {
		for j, _ := range game.enemies.enemy1[i].attacks {
			incrementEnemy1FireballFrameDelay(game, i, j)
			if game.enemies.enemy1[i].attacks[j].frameDelay%2 == 0 {
				incrementEnemyFireballRow(game, i, j)
				if game.enemies.enemy1[i].attacks[j].row >= FRAMES_PER_SHEET {
					rewindEnemy1FireballAnimation(game, i, j)
				}
			}
		}
	}
}
func incrementEnemy1FireballFrameDelay(game *Game, i int, j int) {
	game.enemies.enemy1[i].attacks[j].frameDelay += 1
}
func incrementEnemyFireballRow(game *Game, i int, j int) {
	game.enemies.enemy1[i].attacks[j].row += 1
}
func rewindEnemy1FireballAnimation(game *Game, i int, j int) {
	game.enemies.enemy1[i].attacks[j].column += 1
	game.enemies.enemy1[i].attacks[j].row = 0
}
func removeDeadLevelOneEnemies(game *Game) {
	for i, _ := range game.enemies.enemy1 {
		if i < len(game.enemies.enemy1) && game.enemies.enemy1[i].health <= 0 {
			game.enemies.enemy1[i] = game.enemies.enemy1[len(game.enemies.enemy1)-1]
			game.enemies.enemy1 = append(game.enemies.enemy1[:len(game.enemies.enemy1)-1])
		}
	}
}
func removeSpentEnemy1Fireball(game *Game, i int, j int) {
	game.enemies.enemy1[i].attacks[j] =
		game.enemies.enemy1[i].attacks[len(game.enemies.enemy1[i].attacks)-1]
	game.enemies.enemy1[i].attacks =
		game.enemies.enemy1[i].attacks[:len(game.enemies.enemy1[i].attacks)-1]
}
