package imageformats

import (
	"github.com/dsoprea/go-jpeg-image-structure/v2"
)

const (
	JpegMediaType = "jpeg"
)

func init() {
	register(
		JpegMediaType,
		jpegstructure.NewJpegMediaParser(),
		[]string{".jpg", ".jpeg"})
}
