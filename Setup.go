package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/lafriks/go-tiled"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"log"
	"math"
	"math/rand"
	"os"
	"path"
	"time"
)

func setupGameWindow() {
	windowWidth := WINDOW_WIDTH
	windowHeight := WINDOW_HEIGHT
	ebiten.SetWindowSize(windowWidth, windowHeight)
}
func setupGameStruct() Game {
	game := Game{
		maps:         setupAllMaps(),
		level:        1,
		barrierTiles: make([]BarrierTile, 0, 448),
		player:       setupPlayerAsset(),
		powerUps:     make([]PowerUps, 0, 20),
		cursor:       setupCursor(),
		animations:   setUpAnimations(),
		gameSounds:   setupGameSounds(),
		gameTimers:   setupGameTimers(),
		typeface:     setupTypeFace(),
		enemies:      setupEnemyAssets(),
		gameHUD:      setupHUD(),
	}
	buildLevel2(&game)
	buildLevel3(&game)
	return game
}
func constructLevel3BarrierTiles(game *Game) {
	for i, _ := range game.maps.level3 {
		for tileY := 0; tileY < game.maps.level3[i].level.Height; tileY += 1 {
			for tileX := 0; tileX < game.maps.level3[i].level.Width; tileX += 1 {
				TileXpos := float64(game.maps.level3[i].level.TileWidth * tileX)
				TileYpos := float64(game.maps.level3[i].level.TileHeight * tileY)
				TileHeight := float64(game.maps.level3[i].level.TileHeight)
				TileWidth := float64(game.maps.level3[i].level.TileWidth)

				if game.maps.level3[i].level.Layers[1].
					Tiles[tileY*game.maps.level3[i].level.Width+tileX].ID != 0 {
					game.barrierTiles = append(game.barrierTiles, BarrierTile{
						barrierTile: *game.maps.level3[i].level.Layers[1].
							Tiles[tileY*game.maps.level3[i].level.Width+tileX],
						xLoc:   int(TileXpos),
						yLoc:   int(TileYpos),
						height: int(TileHeight),
						width:  int(TileWidth),
					})
				}
			}
		}
	}
}
func setupAllMaps() AllMaps {
	level3Maps := make([]Map, 0, 8)
	gameMaps := AllMaps{
		level2: Map{},
		level3: level3Maps,
	}
	return gameMaps
}
func buildLevel2(game *Game) {
	game.maps.level2 = setupMap(path.Join("maps", "bee_hive.tmx"))
}
func buildLevel3(game *Game) {
	game.maps.level3 = append(game.maps.level3,
		setupMap(path.Join("maps", "boss_level_map_1.tmx")),
		setupMap(path.Join("maps", "boss_level_map_2.tmx")),
		setupMap(path.Join("maps", "boss_level_map_3.tmx")),
		setupMap(path.Join("maps", "boss_level_map_4.tmx")),
		setupMap(path.Join("maps", "boss_level_map_5.tmx")),
		setupMap(path.Join("maps", "boss_level_map_6.tmx")),
		setupMap(path.Join("maps", "boss_level_map_7.tmx")),
		setupMap(path.Join("maps", "boss_level_map_8.tmx")))
	constructLevel3BarrierTiles(game)
}
func setupMap(fileName string) Map {
	mapPointer, err := tiled.LoadFile(fileName)
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}
	gameMap := makeEbitenImagesFromMap(*mapPointer)
	levelMap := Map{
		level:    mapPointer,
		tileHash: gameMap,
	}
	return levelMap
}
func setupPlayerAsset() Player {
	newPlayer := Player{
		playerSprite: LoadEmbeddedImage("", "walking_man3.png"),
		fireballs:    make([]Attack, 0, 20),
		nukes:        make([]Attack, 0, 3),
		xLoc:         WINDOW_WIDTH / 6,
		yLoc:         WINDOW_HEIGHT / 2,
		direction:    4,
		maxHealth:    100,
		moveSpeed:    1,
		health:       100,
		magicPower:   1,
		canMove:      true,
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
		rainbowManAttack: Animation{
			animation: LoadEmbeddedImage("", "energy3.png"),
		},
		beeAttack: Animation{
			animation: LoadEmbeddedImage("", "energy4.png"),
		},
		heart: Animation{
			animation: LoadEmbeddedImage("", "heart.png"),
		},
		speedBoost: Animation{
			animation: LoadEmbeddedImage("", "speed_boost.png"),
		},
		manaPot: Animation{
			animation: LoadEmbeddedImage("", "mana_pot.png"),
		},
		key: Animation{
			animation: LoadEmbeddedImage("", "key.png"),
		},
		gameOver: Animation{
			animation: LoadEmbeddedImage("", "game_over.png"),
		},
		victory: Animation{
			animation: LoadEmbeddedImage("", "victory.png"),
		},
		questGiver: Animation{
			animation: LoadEmbeddedImage("", "weird_yellow_dude.png"),
		},
		questInstructions: Animation{
			animation: LoadEmbeddedImage("", "quest_instructions.png"),
		},
	}
	return AllAnimations
}
func setupGameSounds() SoundEffects {
	const SOUND_SAMPLE_RATE = 48000
	soundContext := audio.NewContext(SOUND_SAMPLE_RATE)

	gameSounds := SoundEffects{
		playerFire:    LoadWav("avidikadivi.wav", soundContext),
		playerNuke:    LoadWav("avada_kedavra.wav", soundContext),
		teleport:      LoadWav("teleport.wav", soundContext),
		playerStruck:  LoadWav("oof.wav", soundContext),
		powerUp:       LoadWav("powerup.wav", soundContext),
		enemyFired:    LoadWav("pewpew.wav", soundContext),
		enemyStruck:   LoadWav("ow.wav", soundContext),
		enemyKilled:   LoadWav("blegh.wav", soundContext),
		gameMusic:     LoadWav("music.wav", soundContext),
		bossMusic:     LoadWav("final_boss.wav", soundContext),
		gameOverSound: LoadWav("game_over.wav", soundContext),
		gameWonSound:  LoadWav("fanfare.wav", soundContext),
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
		rainbowMan: make([]Enemy, 0, 15),
		bees:       make([]Enemy, 0, 15),
		smileyMan: Enemy{
			enemySprite: LoadEmbeddedImage("", "creepySmily4.png"),
			xLoc:        500,
			yLoc:        200,
			speed:       5,
			health:      1000,
		},
	}
	makeRainbowMen(&allEnemies)
	makeBees(&allEnemies)

	return allEnemies
}

func makeBees(allEnemies *AllEnemies) {
	for i := 0; i < 12; i++ {
		angle := 2 * math.Pi * float64(i) / float64(12)
		x := WINDOW_WIDTH/2 + 200*math.Cos(angle)
		y := WINDOW_HEIGHT/2 + 200*math.Sin(angle)
		allEnemies.bees = append(allEnemies.bees, setupEnemy(int(x)+200, int(y)-100, "bee.png"))
	}
}
func makeRainbowMen(allEnemies *AllEnemies) {
	for i := 700; i < 1000; i += 100 {
		for j := 1; j < 10; j += 2 {
			allEnemies.rainbowMan = append(allEnemies.rainbowMan, setupEnemy(i, (j*60)+200, "rainbow_man.png"))
		}
	}
}
func setupEnemy(xLoc int, yLoc int, file string) Enemy {
	enemy := Enemy{
		enemySprite:       LoadEmbeddedImage("", file),
		xLoc:              float64(xLoc - 100),
		yLoc:              float64(yLoc - 50),
		startingXLoc:      float64(xLoc - 100),
		startingYLoc:      float64(xLoc - 50),
		attackPower:       15,
		health:            100,
		direction:         5,
		speed:             1,
		distanceThreshold: 400,
		shotTimer:         time.Now(),
		frame:             rand.Intn(6),
	}
	return enemy
}
func setupFireballAttack(xLoc int, yLoc int, game Game) Attack {
	attack := Attack{
		xLoc: float64(game.player.xLoc),
		yLoc: float64(game.player.yLoc),
		trajectory: math.Atan2(float64(yLoc)-game.player.yLoc,
			float64(xLoc)-game.player.xLoc),
		damage: FIREBALL_BASE_DAMAGE,
	}
	return attack
}

//	func setupHeart(xLox int, yLoc int, game Game) Animation {
//		heart := PowerUps{
//			powerUpSprite: ,
//
//		}
//	}
func setupNukeAttack(game Game) Attack {
	attack := Attack{
		xLoc:   game.cursor.xLoc,
		yLoc:   game.cursor.yLoc,
		damage: NUKE_ATTACK_BASE_DAMAGE,
	}
	return attack
}
func setupRainbowManAttack(trajectory float64, game Game, i int) Attack {
	attack := Attack{
		xLoc:       float64(game.enemies.rainbowMan[i].xLoc),
		yLoc:       float64(game.enemies.rainbowMan[i].yLoc),
		trajectory: trajectory,
		damage:     15,
	}
	return attack
}
func setupBeeAttack(trajectory float64, game Game, i int) Attack {
	attack := Attack{
		xLoc:       float64(game.enemies.bees[i].xLoc + 25),
		yLoc:       float64(game.enemies.bees[i].yLoc + 25),
		trajectory: trajectory,
		damage:     10,
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
		fireballTimer:    time.Now(),
		electricityTimer: time.Now(),
		nukeTimer:        time.Now(),
		teleportTimer:    time.Now(),
	}
	return gameTimers
}
