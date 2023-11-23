package main

import (
	"math/rand"
)

func randomPowerUp(game *Game, xLoc int, yLoc int) {
	powerUpType := rand.Intn(3)
	if powerUpType == 0 {
		dropHeartPowerUp(game, xLoc, yLoc, powerUpType)
	} else if powerUpType == 1 {
		dropSpeedBoost(game, xLoc, yLoc, powerUpType)
	} else if powerUpType == 2 {
		dropManaPot(game, xLoc, yLoc, powerUpType)
	}
}
func animateHearts(game *Game) {
	for i, _ := range game.powerUps {
		if game.powerUps[i].powerUpType == HEALTH_POT {
			game.powerUps[i].frameDelay += 1
			if game.powerUps[i].frameDelay%4 == 0 {
				game.powerUps[i].row += 1
				if game.powerUps[i].row >= HEARTS_PER_SHEET {
					game.powerUps[i].row = 0
				}
			}
		}
	}
}
func animateSpeedBoost(game *Game) {
	for i, _ := range game.powerUps {
		if game.powerUps[i].powerUpType == SPEED_POT {
			game.powerUps[i].frameDelay += 1
			if game.powerUps[i].frameDelay%8 == 0 {
				game.powerUps[i].row += 1
				if game.powerUps[i].row >= SPEED_BOOST_PER_SHEET {
					game.powerUps[i].row = 0

				}
			}
		}
	}
}
func animateManaPot(game *Game) {
	for i, _ := range game.powerUps {
		if game.powerUps[i].powerUpType == MANA_POT {
			game.powerUps[i].frameDelay += 1
			if game.powerUps[i].frameDelay%4 == 0 {
				game.powerUps[i].row += 1
				if game.powerUps[i].row >= MANA_POT_FRAMES_PER_SHEET {
					game.powerUps[i].row = 0

				}
			}
		}
	}
}
func animateKey(game *Game) {
	for i, _ := range game.powerUps {
		if game.powerUps[i].powerUpType == KEY {
			game.powerUps[i].frameDelay += 1
			if game.powerUps[i].frameDelay%4 == 0 {
				game.powerUps[i].row += 1
				if game.powerUps[i].row >= KEY_FRAMES_PER_SHEET {
					game.powerUps[i].row = 0
				}
			}
		}
	}
}
func dropHeartPowerUp(game *Game, xLoc int, yLoc int, powerUpType int) {
	game.powerUps = append(game.powerUps,
		PowerUps{
			powerUpSprite: game.animations.heart,
			xLoc:          xLoc,
			yLoc:          yLoc,
			powerUpType:   powerUpType,
		})
}
func dropSpeedBoost(game *Game, xLoc int, yLoc int, powerUpType int) {
	game.powerUps = append(game.powerUps,
		PowerUps{
			powerUpSprite: game.animations.speedBoost,
			xLoc:          xLoc,
			yLoc:          yLoc,
			powerUpType:   powerUpType,
		})
}
func dropManaPot(game *Game, xLoc int, yLoc int, powerUpType int) {
	game.powerUps = append(game.powerUps,
		PowerUps{
			powerUpSprite: game.animations.manaPot,
			xLoc:          xLoc,
			yLoc:          yLoc,
			powerUpType:   powerUpType,
		})
}
func dropKey(game *Game, xLoc int, yLoc int, powerUpType int) {
	game.powerUps = append(game.powerUps,
		PowerUps{
			powerUpSprite: game.animations.key,
			xLoc:          xLoc,
			yLoc:          yLoc,
			powerUpType:   powerUpType,
		})
}
func removePowerUp(game *Game, i int) {
	game.powerUps[i] = game.powerUps[len(game.powerUps)-1]
	game.powerUps = append(game.powerUps[:len(game.powerUps)-1])
}
