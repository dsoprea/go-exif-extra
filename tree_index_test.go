package exifextra

import (
	"reflect"
	"sort"
	"testing"

	"path/filepath"

	"github.com/dsoprea/go-exif-extra/common"
	"github.com/dsoprea/go-logging/v2"
)

func TestTreeIndex_AddTree(t *testing.T) {
	ti := NewTreeIndex()

	assetsPath := exifextracommon.GetTestAssetsPath()

	err := ti.AddTree(assetsPath)
	log.PanicIf(err)

	expected := []string{
		"image.heic",
		"image.jpg",
		"image.png",
		"image.tiff",
		"image.webp",
	}

	actual := ti.AddedFiles()

	assetsPathLen := len(assetsPath)
	for i, filepath := range actual {
		if filepath[:assetsPathLen] == assetsPath {
			actual[i] = filepath[assetsPathLen+1:]
		}
	}

	sort.Strings(actual)
	sort.Strings(expected)

	if reflect.DeepEqual(actual, expected) != true {
		t.Fatalf("Files list not correct: %v", actual)
	}
}

func TestTreeIndex_AddPath(t *testing.T) {
	ti := NewTreeIndex()

	assetsPath := exifextracommon.GetTestAssetsPath()

	err := ti.AddPath(assetsPath)
	log.PanicIf(err)

	expected := []string{
		"image.heic",
		"image.jpg",
		"image.png",
		"image.tiff",
		"image.webp",
	}

	actual := ti.AddedFiles()

	assetsPathLen := len(assetsPath)
	for i, filepath := range actual {
		if filepath[:assetsPathLen] == assetsPath {
			actual[i] = filepath[assetsPathLen+1:]
		}
	}

	sort.Strings(actual)
	sort.Strings(expected)

	if reflect.DeepEqual(actual, expected) != true {
		t.Fatalf("Files list not correct: %v", actual)
	}
}

func TestGetFirstIfdMatches(t *testing.T) {
	ti := NewTreeIndex()

	assetsPath := exifextracommon.GetTestAssetsPath()

	err := ti.AddPath(assetsPath)
	log.PanicIf(err)

	values := ti.GetFirstIfdMatches("Software")

	actual := make(map[string]string)
	for _, value := range values {
		actual[filepath.Base(value.Filepath)] = value.ValuePhrase
	}

	expected := map[string]string{
		"image.tiff": "Mac OS X 10.5.8 (9L31a)",
		"image.webp": "GIMP 2.10.24",
	}

	if reflect.DeepEqual(actual, expected) != true {
		t.Fatalf("Files list not correct: %v", actual)
	}
}
