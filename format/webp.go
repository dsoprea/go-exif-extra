package imageformats

import (
	"github.com/dsoprea/go-webp-image-structure"
)

const (
	WebpMediaType = "webp"
)

func init() {
	register(
		WebpMediaType,
		webp.NewWebpMediaParser(),
		[]string{".webp"})
}
