package exifextra

import (
	"fmt"
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

func TestTreeIndex_GetFirstIfdMatches(t *testing.T) {
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

func ExampleTreeIndex_GetFirstIfdMatches() {
	ti := NewTreeIndex()

	assetsPath := exifextracommon.GetTestAssetsPath()

	err := ti.AddPath(assetsPath)
	log.PanicIf(err)

	values := ti.GetFirstIfdMatches("Software")

	byFilename := make(map[string]string)
	for _, value := range values {
		byFilename[filepath.Base(value.Filepath)] = value.ValuePhrase
	}

	for _, filename := range []string{"image.tiff", "image.webp"} {
		fmt.Printf("%s: %s\n", filename, byFilename[filename])
	}

	// Output:
	// image.tiff: Mac OS X 10.5.8 (9L31a)
	// image.webp: GIMP 2.10.24
}

func TestTreeIndex_Search(t *testing.T) {
	ti := NewTreeIndex()

	assetsPath := exifextracommon.GetTestAssetsPath()

	err := ti.AddPath(assetsPath)
	log.PanicIf(err)

	hits, err := ti.Search("GIMP", nil)
	log.PanicIf(err)

	if len(hits) != 1 {
		t.Fatalf("Not exactly one result")
	}

	sr := hits[0]

	if filepath.Base(sr.Filepath) != "image.webp" {
		t.Fatalf("Search result file not correct: [%s]", filepath.Base(sr.Filepath))
	} else if sr.IfdPath != "IFD" {
		t.Fatalf("Search result IFD not correct: [%s]", sr.IfdPath)
	} else if sr.TagName != "Software" {
		t.Fatalf("Search result tag not correct: [%s]", sr.TagName)
	} else if sr.ValuePhrase != "GIMP 2.10.24" {
		t.Fatalf("Search result value not correct: [%s]", sr.ValuePhrase)
	}
}

func ExampleTreeIndex_Search() {
	ti := NewTreeIndex()

	assetsPath := exifextracommon.GetTestAssetsPath()

	err := ti.AddPath(assetsPath)
	log.PanicIf(err)

	hits, err := ti.Search("GIMP", nil)
	log.PanicIf(err)

	sr := hits[0]

	fmt.Printf("File: %s\n", filepath.Base(sr.Filepath))
	fmt.Printf("IfdPath: %s\n", sr.IfdPath)
	fmt.Printf("TagName: %s\n", sr.TagName)
	fmt.Printf("Value: %s\n", sr.ValuePhrase)

	// Output:
	//
	// File: image.webp
	// IfdPath: IFD
	// TagName: Software
	// Value: GIMP 2.10.24
}
