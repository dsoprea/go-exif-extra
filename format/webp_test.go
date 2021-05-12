package imageformats

import (
	"path"
	"testing"

	"io/ioutil"

	"github.com/dsoprea/go-exif-extra"
	"github.com/dsoprea/go-logging/v2"
)

func getTestWebpBytes() (data []byte, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = log.Wrap(state.(error))
		}
	}()

	assetsPath := exifextra.GetTestAssetsPath()
	filepath := path.Join(assetsPath, "image.webp")

	data, err = ioutil.ReadFile(filepath)
	log.PanicIf(err)

	return data, nil
}

func TestGetFormatForExtension_Webp(t *testing.T) {
	name, mp := GetFormatForExtension(".webp")
	if name != WebpMediaType {
		t.Fatalf("HEIF not correctly found by extension.")
	}

	data, err := getTestWebpBytes()
	log.PanicIf(err)

	_, err = mp.ParseBytes(data)
	log.PanicIf(err)
}

func TestGetFormatForBytes_Webp(t *testing.T) {
	data, err := getTestWebpBytes()
	log.PanicIf(err)

	name, mp := GetFormatForBytes(data)
	if name != WebpMediaType {
		t.Fatalf("HEIF not correctly found by extension.")
	}

	_, err = mp.ParseBytes(data)
	log.PanicIf(err)
}
