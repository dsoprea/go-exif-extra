package imageformats

import (
	"github.com/dsoprea/go-tiff-image-structure/v2"
)

const (
	TiffMediaType = "tiff"
)

func init() {
	register(
		TiffMediaType,
		tiffstructure.NewTiffMediaParser(),
		[]string{".tiff", ".tif"})
}
