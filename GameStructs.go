package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/lafriks/go-tiled"
	"golang.org/x/image/font"
	"golang.org/x/sys/windows"
	"time"
)

const (
	WINDOW_WIDTH        = 960
	WINDOW_HEIGHT       = 960
	PLAYER_FRAME_WIDTH  = 31
	PLAYER_FRAME_HEIGHT = 46
	FRAMES_PER_SHEET    = 8
	FIREBALL_SHOT_CLOCK = time.Second * 3
	BEE_SHOT_CLOCK      = time.Second * 2
	NUKE_SHOT_CLOCK     = time.Second
	TELEPORT_SHOT_CLOCK = time.Second * 10
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
	HEALTH_POT = iota
	SPEED_POT
	MANA_POT
	KEY
)
const (
	BEE_FRAME_WIDTH              = 100
	BEE_FRAME_HEIGHT             = 100
	BEE_FRAMES_PER_SHEET         = 5
	SMILEY_FRAME_WIDTH           = 151
	SMILEY_FRAME_HEIGHT          = 237
	SMILEY_FRAMES_PER_SHEET      = 6
	PLAYER_FIREBALL_FRAME_WIDTH  = 38
	PLAYER_FIREBALL_FRAME_HEIGHT = 37
	ENEMY_ATTACK_FRAME_WIDTH     = 63
	ENEMY_ATTACK_FRAME_HEIGHT    = 63
	EXPLOSION_FRAME_WIDTH        = 100
	EXPLOSION_FRAME_HEIGHT       = 73
	EXPLOSION_FRAMES_PER_SHEET   = 5
	HEART_FRAME_WIDTH            = 30
	HEART_FRAME_HEIGHT           = 30
	HEARTS_PER_SHEET             = 10
	SPEED_BOOST_FRAME_WIDTH      = 30
	SPEED_BOOST_FRAME_HEIGHT     = 33
	SPEED_BOOST_PER_SHEET        = 6
	MANA_POT_FRAME_WIDTH         = 21
	MANA_POT_FRAME_HEIGHT        = 44
	MANA_POT_FRAMES_PER_SHEET    = 7
	KEY_FRAME_WIDTH              = 30
	KEY_FRAME_HEIGHT             = 60
	KEY_FRAMES_PER_SHEET         = 9
)
const (
	FIREBALL_BASE_DAMAGE    = 25
	NUKE_ATTACK_BASE_DAMAGE = 2
)

type Map struct {
	level    *tiled.Map
	tileHash map[uint32]*ebiten.Image
}
type AllMaps struct {
	level1     Map
	level2     Map
	level3     []Map
	frame      int
	frameDelay int
}
type BarrierTile struct {
	barrierTile tiled.LayerTile
	xLoc        int
	yLoc        int
	height      int
	width       int
}
type Player struct {
	playerSprite  *ebiten.Image
	fireballs     []Attack
	nukes         []Attack
	xLoc          float64
	yLoc          float64
	keysCollected int
	maxHealth     int
	moveSpeed     float64
	health        int
	magicPower    float64
	direction     int
	canMove       bool
	isMoving      bool
	frame         int
	frameDelay    int
}
type PowerUps struct {
	powerUpSprite Animation
	xLoc          int
	yLoc          int
	powerUpType   int
	row           int
	column        int
	frame         int
	frameDelay    int
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
	startingXLoc      float64
	startingYLoc      float64
	angle             float64
	health            float64
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
	rainbowMan []Enemy
	bees       []Enemy
	smileyMan  Enemy
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
	fireball         Animation
	explosion        Animation
	rainbowManAttack Animation
	beeAttack        Animation
	gameOver         Animation
	victory          Animation
	heart            Animation
	speedBoost       Animation
	manaPot          Animation
	key              Animation
}
type SoundEffects struct {
	playerFire    *audio.Player
	playerNuke    *audio.Player
	teleport      *audio.Player
	playerStruck  *audio.Player
	powerUp       *audio.Player
	enemyFired    *audio.Player
	enemyKilled   *audio.Player
	enemyStruck   *audio.Player
	gameOverSound *audio.Player
	gameWonSound  *audio.Player
	gameMusic     *audio.Player
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
	maps         AllMaps
	level        int
	barrierTiles []BarrierTile
	wallLocation windows.Coord
	wallDetected bool
	cursor       Cursor
	player       Player
	powerUps     []PowerUps
	direction    Movement
	enemies      AllEnemies
	animations   AllAnimations
	gameSounds   SoundEffects
	gameTimers   GameTimers
	typeface     font.Face
	gameHUD      HUD
}
