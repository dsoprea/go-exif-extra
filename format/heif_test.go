package imageformats

import (
	"path"
	"testing"

	"io/ioutil"

	"github.com/dsoprea/go-exif-extra/common"
	"github.com/dsoprea/go-logging/v2"
)

func getTestHeifBytes() (data []byte, err error) {
	defer func() {
		if state := recover(); state != nil {
			err = log.Wrap(state.(error))
		}
	}()

	assetsPath := exifextracommon.GetTestAssetsPath()
	filepath := path.Join(assetsPath, "image.heic")

	data, err = ioutil.ReadFile(filepath)
	log.PanicIf(err)

	return data, nil
}

func TestGetFormatForExtension_Heif(t *testing.T) {
	name, mp := GetFormatForExtension(".heif")
	if name != HeifMediaType {
		t.Fatalf("HEIF not correctly found by extension.")
	}

	data, err := getTestHeifBytes()
	log.PanicIf(err)

	_, err = mp.ParseBytes(data)
	log.PanicIf(err)
}

func TestGetFormatForBytes_Heif(t *testing.T) {
	data, err := getTestHeifBytes()
	log.PanicIf(err)

	name, mp := GetFormatForBytes(data)
	if name != HeifMediaType {
		t.Fatalf("HEIF not correctly found by extension.")
	}

	_, err = mp.ParseBytes(data)
	log.PanicIf(err)
}
