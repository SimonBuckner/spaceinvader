package main

import "github.com/SimonBuckner/screen2d"

/*
	TTL measure time in frames
*/
const (
	// Screen settings
	originalWidth  = 224
	originalHeight = 256

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
	alienColWidth  = 16

	alienStartX = 10
	alienStartY = originalHeight - 0x78

	alienFrameTimer = 1000
)

const (
	// Player settings
	playerHeight = 8
	playerwidth  = 16

	// Start positions for player props
	playerX = 1
	playerY = originalHeight - (playerHeight * 4)

	// shipSpeed = 60
)

const (
	// Player Shot seettings
	playerShotSpeed     = 240
	playerShotMissedY   = 28
	playerShotMissedTTL = 500
)

const scoreTitle = "SCORE<1> HI-SCORE SCORE<2>"

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
