package imageformats

import (
	"github.com/dsoprea/go-jpeg-image-structure/v2"
)

const (
	// JpegMediaType is the name of the registered JPEG parser.
	JpegMediaType = "jpeg"
)

func init() {
	register(
		JpegMediaType,
		jpegstructure.NewJpegMediaParser(),
		[]string{".jpg", ".jpeg"})
}
