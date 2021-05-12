package exifextra

import (
	"fmt"
	"path"
	"sort"
	"strings"

	"io/fs"
	"io/ioutil"
	"path/filepath"

	"github.com/dsoprea/go-exif/v3"
	"github.com/dsoprea/go-logging/v2"
	"github.com/dsoprea/go-utility/v2/filesystem"

	"github.com/dsoprea/go-exif-extra/format"
)

var (
	treeIndexLogger = log.NewLogger("exifextra")
)

// IndexedTag uniquely describes a particular tag.
type IndexedTag struct {
	// IfdPath is the name of the IFD path at which the tag was found.
	IfdPath string

	// TagId is the ID of the tag.
	TagId uint16

	// TagName is the name of the tag.
	TagName string
}

// String returns a string representation of the IndexedTag.
func (it IndexedTag) String() string {
	return fmt.Sprintf("Tag<IFD=[%s] ID=(0x%04x) NAME=[%s]>", it.IfdPath, it.TagId, it.TagName)
}

// IndexedValue is the value for a particular tag in a particular file.
type IndexedValue struct {
	// Filepath is the fully-qualified path of this file.
	Filepath string

	// ValuePhrase is the stringified value of the tag for this file.
	ValuePhrase string
}

// String returns a string representation of the IndexedValue.
func (iv IndexedValue) String() string {
	return fmt.Sprintf("Value<FILE-PATH=[%s] VALUE=[%s]>", iv.Filepath, iv.ValuePhrase)
}

// TreeIndex is an index of the tags of many files.
type TreeIndex struct {
	indexed    map[IndexedTag][]IndexedValue
	addedFiles map[string]struct{}
	addedCount int
}

// NewTreeIndex returns a new TreeIndex.
func NewTreeIndex() *TreeIndex {
	indexed := make(map[IndexedTag][]IndexedValue)
	addedFiles := make(map[string]struct{})

	return &TreeIndex{
		indexed:    indexed,
		addedFiles: addedFiles,
	}
}

// Index is the currently-loaded index.
func (ti *TreeIndex) Index() map[IndexedTag][]IndexedValue {
	return ti.indexed
}

// GetFirstIfdMatches returns the files and values for the given tag in the
// first IFD it was found for. Often, if a tag is presentfor both IFD0/IFD1 and
// the Exif IFD, it will just be duplicated. So, we're just providing a
// convenience function to bypass the obligatory consideration of which, when a)
// people won't generally know which IFD a tag *should* be in, and they wouldn't
// have a preference or even care.
//
// Note that if a particular IFD or looking-up using a tag-ID rather than a name
// is desired, the indexed should be referenced directly.
//
// This method merely returns an empty list if the given tag wasn't found in any
// indexed file.
func (ti *TreeIndex) GetFirstIfdMatches(tagName string) []IndexedValue {
	for it, values := range ti.indexed {
		if it.TagName != tagName {
			continue
		}

		return values
	}

	return []IndexedValue{}
}

// AddedFiles returns a list of the files that have been parsed.
func (ti *TreeIndex) AddedFiles() (files []string) {
	files = make([]string, len(ti.addedFiles))

	i := 0
	for filepath := range ti.addedFiles {
		files[i] = filepath
		i++
	}

	return files
}

// AddTree add the given path, recursively.
func (ti *TreeIndex) AddTree(rootPath string) (err error) {
	defer func() {
		if errRaw := recover(); errRaw != nil {
			err = log.Wrap(errRaw.(error))
		}
	}()

	visitorCb := func(path string, de fs.DirEntry, err error) (errOut error) {
		defer func() {
			if errRaw := recover(); errRaw != nil {
				errOut = log.Wrap(errRaw.(error))
			}
		}()

		if err != nil {
			log.Panicf("Could not walk [%s]: %s", path, err.Error())
		} else if de.IsDir() == false {
			return nil
		}

		errOut = ti.AddPath(path)
		log.PanicIf(errOut)

		return nil
	}

	err = filepath.WalkDir(rootPath, visitorCb)
	log.PanicIf(err)

	return nil
}

// AddPath will load the images from a single path into the index. It
func (ti *TreeIndex) AddPath(rootPath string) (err error) {
	defer func() {
		if errRaw := recover(); errRaw != nil {
			err = log.Wrap(errRaw.(error))
		}
	}()

	files, err := ioutil.ReadDir(rootPath)
	log.PanicIf(err)

	for _, fi := range files {
		filename := fi.Name()

		extension := filepath.Ext(filename)
		if extension == "" {
			continue
		}

		name, mp := imageformats.GetFormatForExtension(extension)
		if mp == nil {
			continue
		}

		imageFilepath := path.Join(rootPath, filename)

		mc, err := mp.ParseFile(imageFilepath)
		log.PanicIf(err)

		_, exifData, err := mc.Exif()
		if err != nil {
			treeIndexLogger.Warningf(nil, "Could not parse [%s]: %s", imageFilepath, err.Error())
			continue
		}

		treeIndexLogger.Debugf(nil, "File [%s] is handled format [%s].", imageFilepath, name)

		// Get tags.

		sb := rifs.NewSeekableBufferWithBytes(exifData)

		tags, _, err := exif.GetFlatExifDataUniversalSearchWithReadSeeker(sb, nil, false)
		log.PanicIf(err)

		thisAddedCount := 0
		for _, tag := range tags {
			if tag.ChildIfdPath != "" {
				continue
			}

			it := IndexedTag{
				IfdPath: tag.IfdPath,
				TagId:   tag.TagId,
				TagName: tag.TagName,
			}

			iv := IndexedValue{
				Filepath: imageFilepath,

				// Note that we're only sharing the first item. This will be
				// inconvenient/misleading if we want to see multiple items of
				// the very rare multiple-item tag.
				ValuePhrase: tag.FormattedFirst,
			}

			if knownValues, found := ti.indexed[it]; found == true {
				ti.indexed[it] = append(knownValues, iv)
			} else {
				ti.indexed[it] = []IndexedValue{iv}
			}

			thisAddedCount++
		}

		ti.addedCount += thisAddedCount
		ti.addedFiles[imageFilepath] = struct{}{}

		treeIndexLogger.Debugf(nil, "Indexed (%d) tags from file [%s]. (%d) total tags have been indexed. (%d) total unique tags have been found.", thisAddedCount, imageFilepath, ti.addedCount, len(ti.indexed))
	}

	return nil
}

// SearchResult represents a single hit.
type SearchResult struct {
	Filepath    string
	IfdPath     string
	TagName     string
	ValuePhrase string
}

// String returns a stringified representation of the result.
func (sr SearchResult) String() string {
	return fmt.Sprintf("SearchResult<FILE-PATH=[%s] IFD=[%s] TAG=[%s] VALUE=[%s]>", sr.Filepath, sr.IfdPath, sr.TagName, sr.ValuePhrase)
}

// Search does a case-insensitive search through the values of the given tags
// acros all indexed files. If no tags are given, all tags will be searched. The
// tag-names are also case-insensitive.
func (ti *TreeIndex) Search(query string, tagNames []string) (hits []SearchResult) {
	var filter sort.StringSlice

	if tagNames != nil && len(tagNames) > 0 {
		filter = make(sort.StringSlice, 0)

		for _, name := range tagNames {
			name = strings.ToLower(name)
			filter = append(filter, name)
		}

		filter.Sort()
	}

	query = strings.ToLower(query)

	hits = make([]SearchResult, 0)
	for it, values := range ti.indexed {
		// Filter tag names.

		if filter != nil {
			tagNameLower := strings.ToLower(it.TagName)
			i := filter.Search(tagNameLower)
			if i >= len(filter) || filter[i] != tagNameLower {
				continue
			}
		}

		// Filter tag values.

		for _, iv := range values {
			value := strings.ToLower(iv.ValuePhrase)

			if strings.Contains(value, query) == true {
				result := SearchResult{
					Filepath:    iv.Filepath,
					IfdPath:     it.IfdPath,
					TagName:     it.TagName,
					ValuePhrase: iv.ValuePhrase,
				}

				hits = append(hits, result)
			}
		}
	}

	return hits
}
