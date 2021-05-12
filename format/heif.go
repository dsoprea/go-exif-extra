package imageformats

import (
	"github.com/dsoprea/go-heic-exif-extractor/v2"
)

const (
	HeifMediaType = "heif"
)

func init() {
	register(
		HeifMediaType,
		heicexif.NewHeicExifMediaParser(),
		[]string{".heic", ".heif"})
}
