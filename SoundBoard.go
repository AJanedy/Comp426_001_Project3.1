package main

func playAvidiKidiviSound(game *Game) {
	game.gameSounds.playerFire.Rewind()
	game.gameSounds.playerFire.Play()
}
func playOofSound(game *Game) {
	game.gameSounds.enemyStruck.Rewind()
	game.gameSounds.enemyStruck.Play()
}
func playOwSound(game *Game) {
	game.gameSounds.playerStruck.Rewind()
	game.gameSounds.playerStruck.Play()
}
func playPewPewSound(game *Game) {
	game.gameSounds.enemyFired.Rewind()
	game.gameSounds.enemyFired.Play()
}
func playAvadaKedavraSound(game *Game) {
	game.gameSounds.playerNuke.Rewind()
	game.gameSounds.playerNuke.Play()
}
func playGameMusic(game *Game) {
	game.gameSounds.gameMusic.Rewind()
	game.gameSounds.gameMusic.SetVolume(.4)
	game.gameSounds.gameMusic.Play()
}
func playBossMusic(game *Game) {
	game.gameSounds.bossMusic.Rewind()
	game.gameSounds.bossMusic.SetVolume(.4)
	game.gameSounds.bossMusic.Play()
}
func playTeleportSound(game *Game) {
	game.gameSounds.teleport.Rewind()
	game.gameSounds.teleport.Play()
}
func playGameOverSound(game *Game) {
	game.gameSounds.gameOverSound.Rewind()
	game.gameSounds.gameOverSound.Play()
}
func playVictorySound(game *Game) {
	game.gameSounds.gameWonSound.Rewind()
	game.gameSounds.gameWonSound.Play()
}
func playEnemyKilledSound(game *Game) {
	game.gameSounds.enemyKilled.Rewind()
	game.gameSounds.enemyKilled.Play()
}
func playPowerUpSound(game *Game) {
	game.gameSounds.powerUp.Rewind()
	game.gameSounds.powerUp.Play()
}
