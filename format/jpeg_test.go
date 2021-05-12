package imageformats

import (
	"path"
	"testing"

	"io/ioutil"

	"github.com/dsoprea/go-exif-extra"
	"github.com/dsoprea/go-logging/v2"
)

func getTestJpegBytes() (data []byte, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = log.Wrap(state.(error))
		}
	}()

	assetsPath := exifextra.GetTestAssetsPath()
	filepath := path.Join(assetsPath, "image.jpg")

	data, err = ioutil.ReadFile(filepath)
	log.PanicIf(err)

	return data, nil
}

func TestGetFormatForExtension_Jpeg(t *testing.T) {
	name, mp := GetFormatForExtension(".jpg")
	if name != JpegMediaType {
		t.Fatalf("HEIF not correctly found by extension.")
	}

	data, err := getTestJpegBytes()
	log.PanicIf(err)

	_, err = mp.ParseBytes(data)
	log.PanicIf(err)
}

func TestGetFormatForBytes_Jpeg(t *testing.T) {
	data, err := getTestJpegBytes()
	log.PanicIf(err)

	name, mp := GetFormatForBytes(data)
	if name != JpegMediaType {
		t.Fatalf("HEIF not correctly found by extension.")
	}

	_, err = mp.ParseBytes(data)
	log.PanicIf(err)
}
