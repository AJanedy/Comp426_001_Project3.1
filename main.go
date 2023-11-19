package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	setupGameWindow()
	game := setupGameStruct()
	ebiten.SetCursorMode(ebiten.CursorModeHidden)
	ebiten.RunGame(&game)
}
