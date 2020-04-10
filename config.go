package main

import "github.com/SimonBuckner/screen2d"

const (
	// Screen settings
	originalWidth  = 224
	originalHeight = 256

	winWidth  = 1024
	winHeight = 768
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

	shipSpeed float32 = 60
	// 	shipExplodeTTL = 15

	// 	shotSpeed      = 540
	// 	shotExplodeTTL = 60
	// 	shotMissedY    = 25
)
const (
	// Player Shot seettings
	playerShotSpeed   float32 = 500
	playerShotMissedY float32 = 25
	// 	shotExplodeTTL = 15
	// 	shotExplodeTTL = 60
)

const title = "SCORE-1   HI-SCORE   SCORE-2"

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
