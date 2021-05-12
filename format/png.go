package imageformats

import (
	"github.com/dsoprea/go-png-image-structure/v2"
)

const (
	// PngMediaType is the name of the registered PNG parser.
	PngMediaType = "png"
)

func init() {
	register(
		PngMediaType,
		pngstructure.NewPngMediaParser(),
		[]string{".png"})
}
