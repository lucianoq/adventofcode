package main

import "fmt"

const (
	tileSize      = 10
	tilesPerImage = 12
	imageSize     = (tileSize - 2) * tilesPerImage
)

var pattern = []string{
	"                  # ",
	"#    ##    ##    ###",
	" #  #  #  #  #  #   ",
}

func main() {
	tiles := parse()

	// this is one of the corner
	// we can build from here but it has to be
	// in the right orientation and flip
	corner := 1187
	tiles[corner].Flip()
	tiles[corner].Rotate()

	// This is the 12*12 matrix to handle the right tile, with correct
	// orientation and flip, in the right place.
	places := [tilesPerImage][tilesPerImage]*Tile{}

	places[0][0] = tiles[corner]
	delete(tiles, corner)

	// for every place in the image
	for i := 0; i < tilesPerImage; i++ {
	Place:
		for j := 0; j < tilesPerImage; j++ {

			if j == 0 {
				// First element has been decided already
				if i == 0 {
					continue Place
				}

				// first element of the row must match with top
				for id, tile := range tiles {
					for f := 0; f < 2; f++ {
						for r := 0; r < 4; r++ {
							if Overlap(places[i-1][j].Edge(Down), tile.Edge(Up)) {
								places[i][j] = tile
								delete(tiles, id)
								continue Place
							}
							tile.Rotate()
						}
						tile.Flip()
					}
				}
			}

			// only for j>0

			// for every remaining tile
			for id, tile := range tiles {
				for f := 0; f < 2; f++ {
					for r := 0; r < 4; r++ {
						if Overlap(places[i][j-1].Edge(Right), tile.Edge(Left)) {
							places[i][j] = tile
							delete(tiles, id)
							continue Place
						}
						tile.Rotate()
					}
					tile.Flip()
				}
			}
		}
	}

	image := MergeTiles(places)

	// Number of hashes that are not in sea monsters
	// are simply the number of all hashes without the number of hashes
	// in a single sea monster, times the number of sea monsters,
	// assuming sea monsters are never overlapped

	// count Hashes in the image
	hashes := 0
	for i := 0; i < imageSize; i++ {
		for j := 0; j < imageSize; j++ {
			if image[i][j] {
				hashes++
			}
		}
	}

	// count sea monsters in the image
	seaMonsters := 0
	for i := 0; i < imageSize-len(pattern); i++ {
		for j := 0; j < imageSize-len(pattern[0]); j++ {
			if SeaMonster(i, j, image) {
				seaMonsters++
			}
		}
	}

	fmt.Println(hashes - 15*seaMonsters)
}

// MergeTiles builds the final image,
// merging the 12*12 tiles, skipping the borders
func MergeTiles(places [tilesPerImage][tilesPerImage]*Tile) [imageSize][imageSize]bool {
	finalImage := [imageSize][imageSize]bool{}

	// I and J loop 12 times, as tiles in the final image
	for I := 0; I < tilesPerImage; I++ {
		for J := 0; J < tilesPerImage; J++ {

			// i and j loop 8 times, from 0 to 8-1
			// we need to take the i+1 and the j+1
			// to always skip the first (border)
			// arriving at 10-2 let us skip the last too (border)
			for i := 0; i < tileSize-2; i++ {
				for j := 0; j < tileSize-2; j++ {
					finalImage[(tileSize-2)*I+i][(tileSize-2)*J+j] = places[I][J][i+1][j+1]
				}
			}
		}
	}

	return finalImage
}

func SeaMonster(x, y int, image [imageSize][imageSize]bool) bool {
	for i := 0; i < len(pattern); i++ {
		for j := 0; j < len(pattern[0]); j++ {
			if pattern[i][j] == ' ' {
				continue
			}

			if !image[x+i][y+j] {
				return false
			}
		}
	}
	return true
}

func (t *Tile) Rotate() {
	newGrid := [tileSize][tileSize]bool{}
	for i := 0; i < tileSize; i++ {
		for j := 0; j < tileSize; j++ {
			newGrid[j][tileSize-i-1] = t[i][j]
		}
	}
	*t = newGrid
}

func (t *Tile) Flip() {
	for i := 0; i < tileSize/2; i++ {
		for j := 0; j < tileSize; j++ {
			t[i][j], t[tileSize-i-1][j] = t[tileSize-i-1][j], t[i][j]
		}
	}
}
