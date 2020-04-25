# spaceinvader
Classic Space Invaders in Go and SDL2

## Notes on the Origianl Hardware

Any information here is taken from the very excellent breakdown of the original
game ROM by Mikael Agren on the Computer Archeology. You can see it here.

    https://computerarcheology.com/Arcade/SpaceInvaders/

The original screen had a resolution of 256x224 pixels, rotated 90 degrees 
anti-clockwise. That makes the screen 256 pixels high but 224 pixels wide.
The original game screen offsets start in the bottom left hand corner of
the screen and finish int he top right.

## Screen Coordinates

The screen rotation makes the original hardware x, y origin the bottom left
hand corner of the screen with the last pixel in the top right. To convert the
original game coordinates (offsets from video memory start address) to X & Y
coordinates we can use in the game the following should be done.

    os = original offset to be converted to x & y
    sv = video memory start address
    oo = original offset
    rw = row width

    x = (os - vs) mod rw
    y = rw - ((os - vs) / rw)

## Notes on Object Positions

The alien rack is based on the bottom left hand (after rotation) alien. Starting
Y coord for each round is. Once round 8 is reached, the aliens go no further.

    Round 1 :  112
    Round 2 :  128
    Round 3 :  144
    Round 4 :  152
    Round 5 :  160
    Round 6 :  160
    Round 7 :  160
    Round 8 :  160
  


## Notes on Game Timing Loops


0 move player
- mid screen
- 60 pixels/sec
- explode for 500ms

1 player shot
- 240 pixels/sec

2 alien shot 1
- Rolling shot
- 80 pixels/sec
- aliens < 8 = 100 pixels/sec
- column above player

3 alien shot 2
- Plunger shot
- 80 pixels/sec
- aliens < 8 = 100 pixels/sec
- Column order = 01 07 01 01 01 04 0B 01 06 03 01 01 0B 09 02 08

4 alien shot 3 (flying saucer)
- Squiggly shot
- 80 pixels/sec
- aliens < 8 = 100 pixels/sec
- Column order = 0B 01 06 03 01 01 0B 09 02 08 02 0B 04 07 0A
Saucer
- every 600 game looks
- player shots even from the right
- player shots odd from the left
- score linked to shots
- if shots > 8 then loop through this table 
- 100 050 050 100 150 100 100 050 300 100 100 100 050 150 100
At the start of every level keep a count of the shots you make as they explode -- whether they hit anything or not. Count to 8 and start over. From then on count to 15 and start over. You want the 15th shot to be the one that hits the saucer (or the first 8th if you can manage it). 


Shot reload rate 
<0x0200 = 0x30
<0x1000 = 0x10
<0x2000 = 0x0B
<0x3000 = 0x08
>0x3000 = 0x07


Alien draw
- 1 per end screen interrupt

## Game Loop

