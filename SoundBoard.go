package main

func playAvidiKidiviSound(game *Game) {
	game.gameSounds.playerFire.Rewind()
	game.gameSounds.playerFire.Play()
}
func playOofSound(game *Game) {
	game.gameSounds.enemy1Struck.Rewind()
	game.gameSounds.enemy1Struck.Play()
}
func playOwSound(game *Game) {
	game.gameSounds.playerStruck.Rewind()
	game.gameSounds.playerStruck.Play()
}
func playPewPewSound(game *Game) {
	game.gameSounds.enemy1Strike.Rewind()
	game.gameSounds.enemy1Strike.Play()
}
