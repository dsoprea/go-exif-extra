package imageformats

import (
	"github.com/dsoprea/go-webp-image-structure"
)

const (
	// WebpMediaType is the name of the registered WEBP parser.
	WebpMediaType = "webp"
)

func init() {
	register(
		WebpMediaType,
		webp.NewWebpMediaParser(),
		[]string{".webp"})
}
