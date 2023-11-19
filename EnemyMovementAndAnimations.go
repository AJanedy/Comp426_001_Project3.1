package main

import (
	"math"
	"time"
)

func moveRainbowMen(game *Game) {
	for i, _ := range game.enemies.rainbowMan {
		if game.enemies.rainbowMan[i].isChasing {
			getRainbowMenMovementDirection(game)
			rainbowMenDeltaXY(game, i)
		}
	}
}

//	func moveBee(game *Game) {
//		for i, _ := range game.enemies.bee {
//			//if game.enemies.bee[i].yLoc <= game.enemies.bee[i].startingYLoc &&
//			//	game {
//
//			}
//		}
//	}
func getRainbowMenMovementDirection(game *Game) {
	for i, _ := range game.enemies.rainbowMan {
		game.enemies.rainbowMan[i].degreesDirection = math.Atan2(float64(game.player.yLoc-game.enemies.rainbowMan[i].yLoc),
			float64(game.player.xLoc-game.enemies.rainbowMan[i].xLoc))
		game.enemies.rainbowMan[i].radiansDirection = game.enemies.rainbowMan[i].degreesDirection * 180 / math.Pi
		translateRadiansToCardinalDirection(game, i)
	}
}
func rainbowMenDeltaXY(game *Game, i int) {
	game.enemies.rainbowMan[i].xLoc += math.Cos(game.enemies.rainbowMan[i].degreesDirection) / 2
	game.enemies.rainbowMan[i].yLoc += math.Sin(game.enemies.rainbowMan[i].degreesDirection) / 2
}
func translateRadiansToCardinalDirection(game *Game, i int) {
	if -157.5 <= game.enemies.rainbowMan[i].radiansDirection && game.enemies.rainbowMan[i].radiansDirection < -112.5 {
		game.enemies.rainbowMan[i].direction = NORTH_WEST
	} else if -112.5 <= game.enemies.rainbowMan[i].radiansDirection && game.enemies.rainbowMan[i].radiansDirection < -67.5 {
		game.enemies.rainbowMan[i].direction = NORTH
	} else if -67.5 <= game.enemies.rainbowMan[i].radiansDirection && game.enemies.rainbowMan[i].radiansDirection < -22.5 {
		game.enemies.rainbowMan[i].direction = NORTH_EAST
	} else if -22.5 <= game.enemies.rainbowMan[i].radiansDirection && game.enemies.rainbowMan[i].radiansDirection < 22.5 {
		game.enemies.rainbowMan[i].direction = EAST
	} else if 22.5 <= game.enemies.rainbowMan[i].radiansDirection && game.enemies.rainbowMan[i].radiansDirection < 67.5 {
		game.enemies.rainbowMan[i].direction = SOUTH_EAST
	} else if 67.5 <= game.enemies.rainbowMan[i].radiansDirection && game.enemies.rainbowMan[i].radiansDirection < 112.5 {
		game.enemies.rainbowMan[i].direction = SOUTH
	} else if 112.5 <= game.enemies.rainbowMan[i].radiansDirection && game.enemies.rainbowMan[i].radiansDirection < 157.5 {
		game.enemies.rainbowMan[i].direction = SOUTH_WEST
	} else if game.enemies.rainbowMan[i].radiansDirection >= 157.5 || game.enemies.rainbowMan[i].radiansDirection <= -157.5 {
		game.enemies.rainbowMan[i].direction = WEST
	}
}
func animateBee(game *Game) {
	for i, _ := range game.enemies.bee {
		game.enemies.bee[i].frameDelay += 1
		if game.enemies.bee[i].frameDelay%3 == 0 {
			game.enemies.bee[i].frame += 1
			if game.enemies.bee[i].frame >= BEE_FRAMES_PER_SHEET {
				game.enemies.bee[i].direction += 1
				game.enemies.bee[i].frame = 0
				if game.enemies.bee[i].direction >= BEE_FRAMES_PER_SHEET {
					game.enemies.bee[i].direction = 0
				}
			}
		}
	}
}
func animateRainbowMen(game *Game) {
	for i, _ := range game.enemies.rainbowMan {
		if game.enemies.rainbowMan[i].isChasing {
			incrementEnemy1FrameDelay(game, i)
			if game.enemies.rainbowMan[i].frameDelay%FRAMES_PER_SHEET == 0 {
				incrementEnemy1Frame(game, i)
				if game.enemies.rainbowMan[i].frame >= FRAMES_PER_SHEET {
					resetEnemyFrame(game, i)
				}
			}
		}
	}
}
func animateLevel3(game *Game) {
	game.maps.frameDelay += 1
	if game.maps.frameDelay%FRAMES_PER_SHEET == 0 {
		game.maps.frame += 1
		if game.maps.frame >= 8 {
			game.maps.frame = 1
		}
	}
}
func incrementEnemy1FrameDelay(game *Game, i int) {
	game.enemies.rainbowMan[i].frameDelay += 1
}
func incrementEnemy1Frame(game *Game, i int) {
	game.enemies.rainbowMan[i].frame += 1
}
func resetEnemyFrame(game *Game, i int) {
	game.enemies.rainbowMan[i].frame = 0
}
func attackingRainbowMen(game *Game) {
	for i, _ := range game.enemies.rainbowMan {
		checkEnemyFiringRange(game, i)
	}
}
func checkEnemyFiringRange(game *Game, i int) {
	if game.enemies.rainbowMan[i].isInRange {
		if time.Since(game.enemies.rainbowMan[i].shotTimer) >= FIREBALL_SHOT_CLOCK {
			enemy1Attack(game, i)
			resetEnemy1AttackCooldown(game, i)
			playPewPewSound(game)
		}
	}
}
func enemy1Attack(game *Game, i int) {
	fireballTrajectory := math.Atan2(float64(game.player.yLoc)-game.enemies.rainbowMan[i].yLoc,
		float64(game.player.xLoc)-game.enemies.rainbowMan[i].xLoc)
	game.enemies.rainbowMan[i].attacks = append(game.enemies.rainbowMan[i].attacks,
		setupEnemyAttack(fireballTrajectory, *game, i))
}
func resetEnemy1AttackCooldown(game *Game, i int) {
	game.enemies.rainbowMan[i].shotTimer = time.Now()
}
func moveRainbowMenFireballs(game *Game) {
	for i, _ := range game.enemies.rainbowMan {
		for j, _ := range game.enemies.rainbowMan[i].attacks {
			game.enemies.rainbowMan[i].attacks[j].xLoc +=
				math.Cos(game.enemies.rainbowMan[i].attacks[j].trajectory) * 3
			game.enemies.rainbowMan[i].attacks[j].yLoc +=
				math.Sin(game.enemies.rainbowMan[i].attacks[j].trajectory) * 3
		}
	}
}
func animateRainbowMenFireballs(game *Game) {
	for i, _ := range game.enemies.rainbowMan {
		for j, _ := range game.enemies.rainbowMan[i].attacks {
			incrementEnemy1FireballFrameDelay(game, i, j)
			if game.enemies.rainbowMan[i].attacks[j].frameDelay%2 == 0 {
				incrementEnemyFireballRow(game, i, j)
				if game.enemies.rainbowMan[i].attacks[j].row >= FRAMES_PER_SHEET {
					rewindEnemy1FireballAnimation(game, i, j)
				}
			}
		}
	}
}
func incrementEnemy1FireballFrameDelay(game *Game, i int, j int) {
	game.enemies.rainbowMan[i].attacks[j].frameDelay += 1
}
func incrementEnemyFireballRow(game *Game, i int, j int) {
	game.enemies.rainbowMan[i].attacks[j].row += 1
}
func rewindEnemy1FireballAnimation(game *Game, i int, j int) {
	game.enemies.rainbowMan[i].attacks[j].column += 1
	game.enemies.rainbowMan[i].attacks[j].row = 0
}
func removeDeadRainbowMen(game *Game) {
	for i, _ := range game.enemies.rainbowMan {
		if i < len(game.enemies.rainbowMan) && game.enemies.rainbowMan[i].health <= 0 {
			game.enemies.rainbowMan[i] = game.enemies.rainbowMan[len(game.enemies.rainbowMan)-1]
			game.enemies.rainbowMan = append(game.enemies.rainbowMan[:len(game.enemies.rainbowMan)-1])
		}
	}
}
func removeSpentEnemy1Fireball(game *Game, i int, j int) {
	game.enemies.rainbowMan[i].attacks[j] =
		game.enemies.rainbowMan[i].attacks[len(game.enemies.rainbowMan[i].attacks)-1]
	game.enemies.rainbowMan[i].attacks =
		game.enemies.rainbowMan[i].attacks[:len(game.enemies.rainbowMan[i].attacks)-1]
}
