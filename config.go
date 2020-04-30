package main

import "github.com/SimonBuckner/screen2d"

/*
	TTL measure time in frames
*/
const (
	// Screen settings
	originalWidth  = 28 * 8 // 224
	originalHeight = 256

	// winWidth  = 250
	// winHeight = 300
	winWidth  = 1024
	winHeight = 768
)

const (
	// Play Mode settings
	pmReadyTTL      = 6 //23 * 2 // number of on/offs
	pmReadyDelayTTL = 3 // Delay between each on/off
)

const (
	// Aliens settings
	alienRows = 5
	alienCols = 11

	alienRowHeight = 16
	alienColWidth  = 15

	alienStartX = 14
	alienStartY = originalHeight - 0x78

	alienFrameTimer = 1000
	alienMinX       = 2 * 8
	alienMaxX       = 25 * 8
)

const (
	// Player settings
	playerHeight = 8
	playerwidth  = 16

	// Start positions for player props
	playerX = 1
	playerY = originalHeight - (playerHeight * 4)

	// shipSpeed = 60
	playerMinX = 2 * 8
	playerMaxX = 26*8 - playerwidth
)

const (
	// Player Shot seettings
	playerShotSpeed     = 240
	playerShotMissedY   = 28
	playerShotMissedTTL = 500
)

const (
	// Score Settings
	scoreTitleX = 1 * 8
	scoreTitleY = 0
	scoreTitle  = "SCORE<1> HI-SCORE SCORE<2>"
	scoreY      = 16
	scoreP1X    = 3 * 8
	scoreHiX    = 22 * 8
	scoreP2X    = 21 * 8
)

const (
	// Keys to image assets
	keyAlienSprCYA screen2d.SpriteMapKey = iota
	keyAlienSprCYB
	keyAlienSprA0
	keyAlienSprA1
	keyAlienSprB0
	keyAlienSprB1
	keyAlienSprC0
	keyAlienSprC1
	keyPlayerSprite
	keyPlrBlowupSprite0
	keyPlrBlowupSprite1
	keyPlayerShotSpr
	keyShotExploding
	keyAlienShotExploding
	keyAlienExplode
	keySquiglyShot0
	keySquiglyShot1
	keySquiglyShot2
	keySquiglyShot3
	keyPlungerShot0
	keyPlungerShot1
	keyPlungerShot2
	keyPlungerShot3
	keyRollShot0
	keyRollShot1
	keyRollShot2
	keyRollShot3
	keyShieldImage
	keySpriteSaucer
	keySpriteSaucerExp
	keyAlienSprCA
	keyAlienSprCB
)
