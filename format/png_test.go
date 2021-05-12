package imageformats

import (
	"path"
	"testing"

	"io/ioutil"

	"github.com/dsoprea/go-exif-extra/common"
	"github.com/dsoprea/go-logging/v2"
)

func getTestPngBytes() (data []byte, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = log.Wrap(state.(error))
		}
	}()

	assetsPath := exifextracommon.GetTestAssetsPath()
	filepath := path.Join(assetsPath, "image.png")

	data, err = ioutil.ReadFile(filepath)
	log.PanicIf(err)

	return data, nil
}

func TestGetFormatForExtension_Png(t *testing.T) {
	name, mp := GetFormatForExtension(".png")
	if name != PngMediaType {
		t.Fatalf("HEIF not correctly found by extension.")
	}

	data, err := getTestPngBytes()
	log.PanicIf(err)

	_, err = mp.ParseBytes(data)
	log.PanicIf(err)
}

func TestGetFormatForBytes_Png(t *testing.T) {
	data, err := getTestPngBytes()
	log.PanicIf(err)

	name, mp := GetFormatForBytes(data)
	if name != PngMediaType {
		t.Fatalf("HEIF not correctly found by extension.")
	}

	_, err = mp.ParseBytes(data)
	log.PanicIf(err)
}
