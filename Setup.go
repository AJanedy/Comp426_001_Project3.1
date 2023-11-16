package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"log"
	"math"
	"time"
)

func setupGameWindow() {
	windowWidth := WINDOW_WIDTH
	windowHeight := WINDOW_HEIGHT
	ebiten.SetWindowSize(windowWidth, windowHeight)
}
func setupGameStruct() Game {
	game := Game{
		player:     setupPlayerAsset(),
		cursor:     setupCursor(),
		animations: setUpAnimations(),
		gameSounds: setupGameSounds(),
		gameTimers: setupGameTimers(),
		typeface:   setupTypeFace(),
		enemies:    setupEnemyAssets(),
		gameHUD:    setupHUD(),
	}
	return game
}
func setupPlayerAsset() Player {
	newPlayer := Player{
		playerSprite: LoadEmbeddedImage("", "walking_man.png"),
		fireballs:    make([]Attack, 0, 20),
		nukes:        make([]Attack, 0, 3),
		xLoc:         WINDOW_WIDTH / 4,
		yLoc:         WINDOW_HEIGHT / 2,
		direction:    4,
		maxHealth:    100,
		health:       100,
		magicPower:   1,
		armor:        1,
		castSpeed:    1,
	}
	return newPlayer
}
func setupCursor() Cursor {
	gameCursor := Cursor{
		cursor: LoadEmbeddedImage("", "cursor.png"),
	}
	return gameCursor
}
func setUpAnimations() AllAnimations {
	AllAnimations := AllAnimations{
		fireball: Animation{
			animation: LoadEmbeddedImage("", "fireball.png"),
		},
		explosion: Animation{
			animation: LoadEmbeddedImage("", "explosion.png"),
		},
		enemy1Attack: Animation{
			animation: LoadEmbeddedImage("", "energy3.png"),
		},
		iceWall: Animation{
			animation: LoadEmbeddedImage("", "energy4.png"),
		},
	}
	return AllAnimations
}
func setupGameSounds() SoundEffects {
	const SOUND_SAMPLE_RATE = 48000
	soundContext := audio.NewContext(SOUND_SAMPLE_RATE)

	gameSounds := SoundEffects{
		playerFire:   LoadWav("avidikadivi.wav", soundContext),
		playerStruck: LoadWav("oof.wav", soundContext),
		enemy1Strike: LoadWav("pewpew.wav", soundContext),
		enemy1Struck: LoadWav("ow.wav", soundContext),
	}
	return gameSounds
}
func setupTypeFace() font.Face {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	gameFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	return gameFont
}
func setupEnemyAssets() AllEnemies {
	allEnemies := AllEnemies{
		enemy1: make([]Enemy, 0, 15),
		//enemy2: setupEnemy2(),
		//enemy3: setupEnemy3(),
	}
	for i := 700; i < 1000; i += 100 {
		for j := 1; j < (10); j += 2 {
			allEnemies.enemy1 = append(allEnemies.enemy1, setupEnemy(i, j*96))
		}
	}
	return allEnemies
}
func setupEnemy(xLoc int, yLoc int) Enemy {
	enemy := Enemy{
		enemySprite:       LoadEmbeddedImage("", "rainbow_man.png"),
		xLoc:              float64(xLoc),
		yLoc:              float64(yLoc),
		attackPower:       15,
		health:            100,
		direction:         6,
		speed:             1,
		distanceThreshold: 400,
		shotTimer:         time.Now(),
	}
	fmt.Println("FIXME: initializing enemy1")
	return enemy
}
func setupShootingAttack(xLoc int, yLoc int, game Game) Attack {
	attack := Attack{
		xLoc: float64(game.player.xLoc),
		yLoc: float64(game.player.yLoc),
		trajectory: math.Atan2(float64(yLoc)-game.player.yLoc,
			float64(xLoc)-game.player.xLoc),
		damage: ATTACK_1_BASE_DAMAGE,
	}
	return attack
}
func setupStationaryAttack(game Game) Attack {
	attack := Attack{
		xLoc: game.cursor.xLoc,
		yLoc: game.cursor.yLoc,
	}
	return attack
}
func setupEnemyAttack(trajectory float64, game Game, i int) Attack {
	attack := Attack{
		xLoc:       float64(game.enemies.enemy1[i].xLoc),
		yLoc:       float64(game.enemies.enemy1[i].yLoc),
		trajectory: trajectory,
		damage:     15,
	}
	return attack
}
func setupHUD() HUD {
	gameHud := HUD{
		WASD: LoadEmbeddedImage("", "WASD.png"),
	}
	return gameHud
}
func setupGameTimers() GameTimers {
	gameTimers := GameTimers{
		attack1Timer: time.Now(),
		attack2Timer: time.Now(),
		qAttackTimer: time.Now(),
		eAttackTimer: time.Now(),
	}
	return gameTimers
}
