package main

import (
	"embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
	"log"
	"path"
)

//go:embed assets/*
var EmbeddedAssets embed.FS

func LoadEmbeddedImage(folderName string, imageName string) *ebiten.Image {
	embeddedFile, err := EmbeddedAssets.Open(path.Join("assets", folderName, imageName))
	if err != nil {
		log.Fatal("failed to load embedded image ", imageName, err)
	}
	ebitenImage, _, err := ebitenutil.NewImageFromReader(embeddedFile)
	if err != nil {
		fmt.Println("Error loading tile image:", imageName, err)
	}
	return ebitenImage
}

//go:embed sounds/*
var EmbeddedSounds embed.FS

func LoadWav(soundFile string, context *audio.Context) *audio.Player {
	embeddedSoundFile, err := EmbeddedSounds.Open(path.Join("sounds", soundFile))
	if err != nil {
		fmt.Println("Error Loading sound: ", err)
	}
	gameSound, err := wav.DecodeWithoutResampling(embeddedSoundFile)
	if err != nil {
		fmt.Println("Error interpreting sound file: ", err)
	}
	soundPlayer, err := context.NewPlayer(gameSound)
	if err != nil {
		fmt.Println("Couldn't create sound player: ", err)
	}
	return soundPlayer
}

func makeEbitenImagesFromMap(tiledMap tiled.Map) map[uint32]*ebiten.Image {
	idToImage := make(map[uint32]*ebiten.Image)
	for _, tile := range tiledMap.Tilesets[0].Tiles {
		ebitenImageTile, _, err :=
			ebitenutil.NewImageFromFile(path.Join("tiles", tile.Image.Source))
		if err != nil {
			fmt.Println("Error loading tile image:",
				tile.Image.Source, err)
		}
		idToImage[tile.ID] = ebitenImageTile
	}
	return idToImage
}
