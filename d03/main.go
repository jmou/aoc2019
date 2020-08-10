// XXX bitmap or scanning algo?
package main

import (
	"bufio"
	"fmt"
	"github.com/golang-collections/go-datastructures/bitarray"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x, y int
}

// XXX grok ownership
func parse(path string) ([]Coord, error) {
	components := strings.Split(path, ",")
	coords := make([]Coord, len(components))
	coord := Coord{}
	coords[0] = coord
	for i, component := range components {
		length, err := strconv.Atoi(component[1:])
		if err != nil {
			return nil, err
		}

		switch component[0] {
		case 'U':
			coord.y += length
		case 'R':
			coord.x += length
		case 'D':
			coord.y -= length
		case 'L':
			coord.x -= length
		}
		coords[i] = coord
	}
	return coords, nil
}

func coords_to_bitmap(coords []Coord) bitarray.BitArray {
	boundmin, boundmax := Coord{}, Coord{}
	for _, c := range coords {
		if c.x < boundmin.x {
			boundmin.x = c.x
		}
		if c.x > boundmax.x {
			boundmax.x = c.x
		}
		if c.y < boundmin.y {
			boundmin.y = c.y
		}
		if c.y > boundmax.y {
			boundmax.y = c.y
		}
	}
	width := boundmax.x + 1 - boundmin.x
	height := boundmax.y + 1 - boundmin.y

	bitpath := bitarray.NewBitArray(uint64(width * height))
	c0 := coords[0]
	for _, c1 := range coords[1:] {
		if c0.x == c1.x {
			x := c0.x - boundmin.x
			y0 := c0.y - boundmin.y
			y1 := c1.y - boundmin.y
			if y1 < y0 {
				y0, y1 = y1, y0
			}
			for y := y0; y <= y1; y++ {
				bitpath.SetBit(uint64(y*width + x))
			}
		} else if c0.y == c1.y {
			y := c0.y - boundmin.y
			x0 := c0.x - boundmin.x
			x1 := c1.x - boundmin.x
			if x1 < x0 {
				x0, x1 = x1, x0
			}
			for x := x0; x <= x1; x++ {
				bitpath.SetBit(uint64(y*width + x))
			}
		} else {
			panic("non-orthogonal path")
		}
		c0 = c1
	}

	return bitpath
}

func main() {
	var bitpaths [2]bitarray.BitArray
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 2; i++ {
		scanner.Scan()
		coords, err := parse(scanner.Text())
		if err != nil {
			panic(err)
		}
		// XXX redo bounding
		bitpaths[i] = coords_to_bitmap(coords)
	}
	fmt.Println(bitpaths[0].And(bitpaths[1]))
}
