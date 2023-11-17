package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/lafriks/go-tiled"
	"golang.org/x/image/font"
	"time"
)

const (
	WINDOW_WIDTH           = 960
	WINDOW_HEIGHT          = 960
	PLAYER_FRAME_WIDTH     = 31
	PLAYER_FRAME_HEIGHT    = 46
	FRAMES_PER_SHEET       = 8
	FIREBALL_SHOT_CLOCK    = time.Second * 3
	ELECTRICITY_SHOT_CLOCK = time.Second * 5
	NUKE_SHOT_CLOCK        = time.Second * 15
	TELEPORT_SHOT_CLOCK    = time.Second * 10
)
const (
	NORTH = iota
	NORTH_EAST
	EAST
	SOUTH_EAST
	SOUTH
	SOUTH_WEST
	WEST
	NORTH_WEST
	DISAPPEARING
	REAPPEARING
	STAND_STILL
)
const (
	PLAYER_FIREBALL_FRAME_WIDTH  = 38
	PLAYER_FIREBALL_FRAME_HEIGHT = 37
	ENEMY_ATTACK_FRAME_WIDTH     = 63
	ENEMY_ATTACK_FRAME_HEIGHT    = 63
	EXPLOSION_FRAME_WIDTH        = 100
	EXPLOSION_FRAME_HEIGHT       = 73
	EXPLOSION_FRAMES_PER_SHEET   = 5
	ICE_FRAME_WIDTH              = 100
	ICE_FRAME_HEIGHT             = 186
	ICE_FRAMES_PER_SHEET         = 6
)
const (
	ATTACK_1_BASE_DAMAGE = 25
)

type Map struct {
	level    *tiled.Map
	tileHash map[uint32]*ebiten.Image
}
type AllMaps struct {
	level1 Map
	level2 Map
	level3 Map
}
type Player struct {
	playerSprite *ebiten.Image
	fireballs    []Attack
	nukes        []Attack
	iceWalls     []Attack
	xLoc         float64
	yLoc         float64
	maxHealth    int
	health       int
	magicPower   int
	armor        int
	castSpeed    int
	direction    int
	canMove      bool
	isMoving     bool
	frame        int
	frameDelay   int
}
type Movement struct {
	moveNorth     bool
	moveNorthEast bool
	moveEast      bool
	moveSouthEast bool
	moveSouth     bool
	moveSouthWest bool
	moveWest      bool
	moveNorthWest bool
}
type Enemy struct {
	enemySprite       *ebiten.Image
	attacks           []Attack
	shotTimer         time.Time
	xLoc              float64
	yLoc              float64
	health            int
	attackPower       int
	degreesDirection  float64
	radiansDirection  float64
	direction         int
	speed             int
	distanceThreshold int
	isMoving          bool
	isChasing         bool
	isInRange         bool
	frame             int
	frameDelay        int
}
type AllEnemies struct {
	enemy1 []Enemy
	enemy2 []Enemy
	enemy3 []Enemy
}
type Attack struct {
	animation  *ebiten.Image
	xLoc       float64
	yLoc       float64
	trajectory float64
	targetX    float64
	targetY    float64
	damage     int
	row        int
	column     int
	frame      int
	frameDelay int
}
type Animation struct {
	animation *ebiten.Image
}
type AllAnimations struct {
	fireball     Animation
	explosion    Animation
	enemy1Attack Animation
	iceWall      Animation
}
type SoundEffects struct {
	playerFire    *audio.Player
	playerStruck  *audio.Player
	playerKilled  *audio.Player
	powerUp1      *audio.Player
	powerUp2      *audio.Player
	powerUp3      *audio.Player
	enemy1Strike  *audio.Player
	enemy2Strike  *audio.Player
	enemy3Strike  *audio.Player
	enemy1Struck  *audio.Player
	enemy2Struck  *audio.Player
	enemy3Struck  *audio.Player
	enemy1Killed  *audio.Player
	enemy2Killed  *audio.Player
	enemy3Killed  *audio.Player
	gameOverSound *audio.Player
	gameWonSound  *audio.Player
}
type HUD struct {
	WASD *ebiten.Image
}
type Cursor struct {
	cursor *ebiten.Image
	xLoc   float64
	yLoc   float64
}
type GameTimers struct {
	fireballTimer    time.Time
	electricityTimer time.Time
	nukeTimer        time.Time
	teleportTimer    time.Time
}
type Game struct {
	maps       AllMaps
	cursor     Cursor
	player     Player
	direction  Movement
	enemies    AllEnemies
	animations AllAnimations
	gameSounds SoundEffects
	gameTimers GameTimers
	typeface   font.Face
	gameHUD    HUD
}
