package imageformats

import (
	"github.com/dsoprea/go-png-image-structure/v2"
)

const (
	PngMediaType = "png"
)

func init() {
	register(
		PngMediaType,
		pngstructure.NewPngMediaParser(),
		[]string{".png"})
}
