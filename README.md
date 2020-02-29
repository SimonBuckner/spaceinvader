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
  