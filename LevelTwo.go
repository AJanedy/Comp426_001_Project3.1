package main

import (
	"math"
	"time"
)

func animateBee(game *Game) {
	for i, _ := range game.enemies.bees {
		game.enemies.bees[i].frameDelay += 1
		if game.enemies.bees[i].frameDelay%3 == 0 {
			game.enemies.bees[i].frame += 1
			spinBees(game, i)
			if game.enemies.bees[i].frame >= BEE_FRAMES_PER_SHEET {
				game.enemies.bees[i].direction += 1
				game.enemies.bees[i].frame = 0
				if game.enemies.bees[i].direction >= BEE_FRAMES_PER_SHEET {
					game.enemies.bees[i].direction = 0
				}
			}
		}
	}
}
func spinBees(game *Game, i int) {
	game.enemies.bees[i].angle += 10
	game.enemies.bees[i].xLoc = game.enemies.bees[i].xLoc + 2*
		math.Cos((math.Pi/180)*game.enemies.bees[i].angle)
	game.enemies.bees[i].yLoc = game.enemies.bees[i].yLoc + 2*
		math.Sin((math.Pi/180)*game.enemies.bees[i].angle)
}
func removeDeadBees(game *Game) {
	for i, _ := range game.enemies.bees {
		if i < len(game.enemies.bees) && game.enemies.bees[i].health <= 0 {
			randomPowerUp(game, int(game.enemies.bees[i].xLoc+25), int(game.enemies.bees[i].yLoc+25))
			game.enemies.bees[i] = game.enemies.bees[len(game.enemies.bees)-1]
			game.enemies.bees = append(game.enemies.bees[:len(game.enemies.bees)-1])
			playEnemyKilledSound(game)
			if len(game.enemies.bees) == 0 {
				dropKey(game, 480, 480, KEY)
			}
		}
	}
}
func attackingBees(game *Game) {
	for i, _ := range game.enemies.bees {
		checkBeeFiringRange(game, i)
	}
}
func checkBeeFiringRange(game *Game, i int) {
	if game.enemies.bees[i].isInRange {
		if time.Since(game.enemies.bees[i].shotTimer) >= BEE_SHOT_CLOCK {
			beeAttack(game, i)
			resetBeeAttackCooldown(game, i)
			playPewPewSound(game)
		}
	}
}
func beeAttack(game *Game, i int) {
	fireballTrajectory := math.Atan2(float64(game.player.yLoc)-game.enemies.bees[i].yLoc,
		float64(game.player.xLoc)-game.enemies.bees[i].xLoc)
	game.enemies.bees[i].attacks = append(game.enemies.bees[i].attacks,
		setupBeeAttack(fireballTrajectory, *game, i))
}
func resetBeeAttackCooldown(game *Game, i int) {
	game.enemies.bees[i].shotTimer = time.Now()
}
func moveBeeAttack(game *Game) {
	for i, _ := range game.enemies.bees {
		for j, _ := range game.enemies.bees[i].attacks {
			game.enemies.bees[i].attacks[j].xLoc +=
				math.Cos(game.enemies.bees[i].attacks[j].trajectory) * 3
			game.enemies.bees[i].attacks[j].yLoc +=
				math.Sin(game.enemies.bees[i].attacks[j].trajectory) * 3
		}
	}
}
func animateBeeAttack(game *Game) {
	for i, _ := range game.enemies.bees {
		for j, _ := range game.enemies.bees[i].attacks {
			game.enemies.bees[i].attacks[j].frameDelay += 1
			if game.enemies.bees[i].attacks[j].frameDelay%2 == 0 {
				game.enemies.bees[i].attacks[j].row += 1
				if game.enemies.bees[i].attacks[j].row >= FRAMES_PER_SHEET {
					game.enemies.bees[i].attacks[j].column += 1
					game.enemies.bees[i].attacks[j].row = 0
				}
			}
		}
	}
}
func removeSpentBeeAttack(game *Game, i int, j int) {
	game.enemies.bees[i].attacks[j] =
		game.enemies.bees[i].attacks[len(game.enemies.bees[i].attacks)-1]
	game.enemies.bees[i].attacks =
		game.enemies.bees[i].attacks[:len(game.enemies.bees[i].attacks)-1]
}
