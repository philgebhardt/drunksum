package drunk

import (
	"fmt"
	"math"
	"strings"
)

const (
	PRINT = " .o+=*BOX@%&#/^SE!"
)

type Bishop interface {
	Print()
}

type bishop struct {
	width  int
	height int
	grid   []byte
}

func (b bishop) Print() {
	r := []rune(PRINT)
	fmt.Printf("_%s_\n", strings.Repeat(" ", b.width+2))
	for i, v := range b.grid {
		if i%b.width == 0 {
			fmt.Print(" ")
		}

		fmt.Printf("%c", r[int(math.Min(float64(v), float64(len(PRINT)-1)))])

		if i%b.width == b.width-1 {
			fmt.Println()
		}
	}
	fmt.Printf("_%s_\n", strings.Repeat(" ", b.width+2))
}

// / NewBishop creates a drunken bishop representation of the incoming data values. Largely modeled after http://www.dirk-loss.de/sshvis/drunken_bishop.pdf
func NewBishop(width, height int, data []byte) (Bishop, error) {
	grid := make([]byte, width*height)

	start := len(grid) / 2
	x := start
	lenData := len(data)
	for i := 0; i < lenData; i++ {
		v := data[i]
		v0 := 0x3 & v
		v1 := (0xC & v) >> 2

		offset := Offset(x, width, height, v0)
		x += offset
		grid[x]++

		offset = Offset(x, width, height, v1)
		grid[x]++
	}
	grid[start] = 15
	grid[x] = 16
	b := bishop{
		width:  width,
		height: height,
		grid:   grid,
	}
	return b, nil
}

func Offset(x, width, height int, v byte) int {
	switch x {
	case 0:
		return offseta(v)
	case width - 1:
		return offsetb(v)
	case width * (height - 1):
		return offsetc(v)
	case width*height - 1:
		return offsetd(v)
	default:
		if isTop(x, width) {
			return offsetT(v)
		} else if isBottom(x, width, height) {
			return offsetB(v)
		} else if isLeft(x, width) {
			return offsetL(v)
		} else if isRight(x, width) {
			return offsetR(v)
		} else {
			return offsetM(v)
		}
	}
}

func isTop(x, width int) bool {
	return x < width
}

func isBottom(x, width, height int) bool {
	return x > width*(height-1)
}

func isLeft(x, width int) bool {
	return x%width == 0
}

func isRight(x, width int) bool {
	return x%width == width-1
}

func offsetM(v byte) int {
	switch v {
	case 0:
		return -18
	case 1:
		return -16
	case 2:
		return 16
	case 3:
		return 18
	default:
		panic(fmt.Sprintf("invalid offsetM: %d", v))
	}
}

func offsetT(v byte) int {
	switch v {
	case 0:
		return -1
	case 1:
		return 1
	case 2:
		return 16
	case 3:
		return 18
	default:
		panic(fmt.Sprintf("invalid offsetT: %d", v))
	}
}

func offsetB(v byte) int {
	switch v {
	case 0:
		return -18
	case 1:
		return -16
	case 2:
		return -1
	case 3:
		return 1
	default:
		panic(fmt.Sprintf("invalid offsetB: %d", v))
	}
}

func offsetL(v byte) int {
	switch v {
	case 0:
		return -17
	case 1:
		return -16
	case 2:
		return 17
	case 3:
		return 18
	default:
		panic(fmt.Sprintf("invalid offsetL: %d", v))
	}
}

func offsetR(v byte) int {
	switch v {
	case 0:
		return -18
	case 1:
		return -17
	case 2:
		return 16
	case 3:
		return 17
	default:
		panic(fmt.Sprintf("invalid offsetR: %d", v))
	}
}

func offseta(v byte) int {
	switch v {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 17
	case 3:
		return 18
	default:
		panic(fmt.Sprintf("invalid offseta: %d", v))
	}
}

func offsetb(v byte) int {
	switch v {
	case 0:
		return -1
	case 1:
		return 0
	case 2:
		return 16
	case 3:
		return 17
	default:
		panic(fmt.Sprintf("invalid offsetb: %d", v))
	}
}

func offsetc(v byte) int {
	switch v {
	case 0:
		return -17
	case 1:
		return -16
	case 2:
		return 0
	case 3:
		return 1
	default:
		panic(fmt.Sprintf("invalid offsetc: %d", v))
	}
}

func offsetd(v byte) int {
	switch v {
	case 0:
		return -18
	case 1:
		return -17
	case 2:
		return -1
	case 3:
		return 0
	default:
		panic(fmt.Sprintf("invalid offsetd: %d", v))
	}
}
