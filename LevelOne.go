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
func animateRainbowMen(game *Game) {
	for i, _ := range game.enemies.rainbowMan {
		if game.enemies.rainbowMan[i].isChasing {
			game.enemies.rainbowMan[i].frameDelay += 1
			if game.enemies.rainbowMan[i].frameDelay%FRAMES_PER_SHEET == 0 {
				game.enemies.rainbowMan[i].frame += 1
				if game.enemies.rainbowMan[i].frame >= FRAMES_PER_SHEET {
					game.enemies.rainbowMan[i].frame = 0
				}
			}
		}
	}
}
func attackingRainbowMen(game *Game) {
	for i, _ := range game.enemies.rainbowMan {
		checkRainbowManFiringRange(game, i)
	}
}
func checkRainbowManFiringRange(game *Game, i int) {
	if game.enemies.rainbowMan[i].isInRange {
		if time.Since(game.enemies.rainbowMan[i].shotTimer) >= FIREBALL_SHOT_CLOCK {
			rainbowManAttack(game, i)
			resetRainbowManAttackCooldown(game, i)
			playPewPewSound(game)
		}
	}
}
func rainbowManAttack(game *Game, i int) {
	fireballTrajectory := math.Atan2(float64(game.player.yLoc)-game.enemies.rainbowMan[i].yLoc,
		float64(game.player.xLoc)-game.enemies.rainbowMan[i].xLoc)
	game.enemies.rainbowMan[i].attacks = append(game.enemies.rainbowMan[i].attacks,
		setupRainbowManAttack(fireballTrajectory, *game, i))
}
func resetRainbowManAttackCooldown(game *Game, i int) {
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
			game.enemies.rainbowMan[i].attacks[j].frameDelay += 1
			if game.enemies.rainbowMan[i].attacks[j].frameDelay%2 == 0 {
				game.enemies.rainbowMan[i].attacks[j].row += 1
				if game.enemies.rainbowMan[i].attacks[j].row >= FRAMES_PER_SHEET {
					game.enemies.rainbowMan[i].attacks[j].column += 1
					game.enemies.rainbowMan[i].attacks[j].row = 0
				}
			}
		}
	}
}
func removeDeadRainbowMen(game *Game) {
	for i, _ := range game.enemies.rainbowMan {
		if i < len(game.enemies.rainbowMan) && game.enemies.rainbowMan[i].health <= 0 {
			randomPowerUp(game, int(game.enemies.rainbowMan[i].xLoc), int(game.enemies.rainbowMan[i].yLoc))
			game.enemies.rainbowMan[i] = game.enemies.rainbowMan[len(game.enemies.rainbowMan)-1]
			game.enemies.rainbowMan = append(game.enemies.rainbowMan[:len(game.enemies.rainbowMan)-1])
			playEnemyKilledSound(game)
			if len(game.enemies.rainbowMan) == 0 {
				dropKey(game, 480, 480, KEY)
			}
		}
	}
}
func removeSpentRainbowManFireball(game *Game, i int, j int) {
	game.enemies.rainbowMan[i].attacks[j] =
		game.enemies.rainbowMan[i].attacks[len(game.enemies.rainbowMan[i].attacks)-1]
	game.enemies.rainbowMan[i].attacks =
		game.enemies.rainbowMan[i].attacks[:len(game.enemies.rainbowMan[i].attacks)-1]

}
