package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fs := loadFS()
	fs.Defrag()
	fmt.Println(fs.Checksum())
}

func loadFS() FS {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	files := make(map[int]*Block)
	freeSpace := make([]*Block, 0, len(line)/2)

	cursor := 0
	for i, c := range line {
		num := int(c - '0')

		if i&1 == 0 {
			// on odd indexes, it's a file
			id := i / 2
			files[id] = &Block{
				ID:    id,
				Start: cursor,
				End:   cursor + num,
			}
			id++
		} else {
			// on even indexes, it's free space
			freeSpace = append(freeSpace, &Block{
				ID:    -1,
				Start: cursor,
				End:   cursor + num,
			})
		}

		cursor += num
	}

	return FS{
		Files:     files,
		FreeSpace: freeSpace,
		MaxFileID: len(line) / 2,
	}
}

type Block struct {
	ID, Start, End int
}

func (b Block) Size() int {
	return b.End - b.Start
}

type FS struct {
	Files     map[int]*Block
	FreeSpace []*Block
	MaxFileID int
}

func (fs FS) Defrag() {
	for toMove := fs.MaxFileID; toMove >= 0; toMove-- {
		f := fs.Files[toMove]

		for i := 0; i < len(fs.FreeSpace); i++ {
			currEmpty := fs.FreeSpace[i]

			// Has to be big enough
			if currEmpty.Size() < f.Size() {
				continue
			}

			// Do not move to the right
			if currEmpty.Start > f.Start {
				break
			}

			// Found a valid empty space. Move file
			f.Start, f.End = currEmpty.Start, currEmpty.Start+f.Size()
			currEmpty.Start += f.Size()

			// If the emptySpace doesn't exist anymore, clean it up
			if currEmpty.Size() == 0 {
				fs.FreeSpace = append(fs.FreeSpace[:i], fs.FreeSpace[i+1:]...)
			}

			// After the file is moved, go to the next file
			break
		}
	}
}

func (fs FS) Checksum() int {
	sum := 0
	for _, f := range fs.Files {
		for i := f.Start; i < f.End; i++ {
			sum += f.ID * i
		}
	}
	return sum
}
