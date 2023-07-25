package dcc

import (
	"image/color"

	"github.com/gravestench/dcc/pkg"
)

type (
	DCC              = pkg.DCC
	Direction        = pkg.Direction
	Frame            = pkg.Frame
	Cell             = pkg.Cell
	PixelBuffer      = pkg.PixelBuffer
	PixelBufferEntry = pkg.PixelBufferEntry
	CompressionFlag  = pkg.CompressionFlag
)

func FromBytes(data []byte) (*DCC, error) {
	return pkg.FromBytes(data)
}

func New() *DCC {
	return pkg.New()
}

func DefaultPalette() *color.Palette {
	return pkg.DefaultPalette()
}

func Dir64ToDcc(direction, numDirections int) int {
	return pkg.Dir64ToDcc(direction, numDirections)
}
