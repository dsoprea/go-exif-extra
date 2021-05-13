package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/dsoprea/go-exif-extra"
	"github.com/dsoprea/go-logging/v2"
	"github.com/eiannone/keyboard"
	"github.com/jessevdk/go-flags"
	"github.com/wayneashleyberry/terminal-dimensions"
)

type parameters struct {
	Path      string `short:"p" long:"path" required:"true" description:"Path to load recursively"`
	IsVerbose bool   `short:"v" long:"verbose" description:"Print logging"`
}

var (
	arguments = new(parameters)
)

func printGroupedFiles(valuePhrase string, files []string) (err error) {
	defer func() {
		if errRaw := recover(); errRaw != nil {
			err = errRaw.(error)
		}
	}()

	fmt.Printf("(%d) files with value [%s]:\n", len(files), valuePhrase)
	fmt.Printf("\n")

	for _, filepath := range files {
		fmt.Println(filepath)
	}

	fmt.Printf("\n")

	return nil
}

func printValues(it exifextra.IndexedTag, ti *exifextra.TreeIndex) (err error) {
	defer func() {
		if errRaw := recover(); errRaw != nil {
			err = errRaw.(error)
		}
	}()

	fmt.Printf("Occurrences of [%s] in IFD [%s]:\n", it.TagName, it.IfdPath)
	fmt.Printf("\n")

	index := ti.Index()
	values := index[it]

	groupedFiles := make(map[string][]string)

	for _, value := range values {
		if files, found := groupedFiles[value.ValuePhrase]; found == true {
			groupedFiles[value.ValuePhrase] = append(files, value.Filepath)
		} else {
			groupedFiles[value.ValuePhrase] = []string{value.Filepath}
		}
	}

	i := 0
	valuePhrases := make([]string, len(groupedFiles))
	for valuePhrase := range groupedFiles {
		valuePhrases[i] = valuePhrase
		i++
	}

	sort.Strings(valuePhrases)

	for _, value := range valuePhrases {
		printGroupedFiles(value, groupedFiles[value])

		fmt.Printf("\n")
	}

	fmt.Printf("\n")
	fmt.Printf("Press any key to continue.")

	_, _, err = keyboard.GetSingleKey()
	log.PanicIf(err)

	fmt.Printf("\n")

	return nil
}

func showTagsMenu(tagPhrases []string, mapping map[string]exifextra.IndexedTag, ti *exifextra.TreeIndex) (repeat bool, err error) {
	defer func() {
		if errRaw := recover(); errRaw != nil {
			err = errRaw.(error)
		}
	}()

	terminalWidthRaw, err := terminaldimensions.Width()
	log.PanicIf(err)

	terminalWidth := int(terminalWidthRaw)

	choices := make([]string, len(tagPhrases))
	maxChoiceWidth := 0
	for i, tagPhrase := range tagPhrases {
		choice := fmt.Sprintf("(%d) %s", i, tagPhrase)
		choices[i] = choice

		l := len(choice)
		if l > maxChoiceWidth {
			maxChoiceWidth = l
		}
	}

	spacingLength := 2
	maxChoiceWidth += spacingLength

	columnCount := int(math.Floor(float64(terminalWidth) / float64(maxChoiceWidth)))
	rowCount := int(math.Ceil(float64(len(tagPhrases)) / float64(columnCount)))

	for i := 0; i < rowCount; i++ {
		for j := 0; j < columnCount; j++ {
			choiceIndex := i + rowCount*j

			if choiceIndex >= len(choices) {
				continue
			}

			choice := choices[choiceIndex]

			fmt.Print(choice)

			// Print spacing prior to next column.
			if j < columnCount-1 {
				fmt.Print(strings.Repeat(" ", maxChoiceWidth-len(choice)))
			}
		}

		fmt.Print("\n")
	}

	fmt.Printf("\n")
	fmt.Printf("Enter the number of a found tag (or 'q' to quit): ")

	itemNumberPhrase := ""

	for {
		fmt.Scanln(&itemNumberPhrase)

		itemNumberPhrase = strings.ToLower(itemNumberPhrase)

		if itemNumberPhrase == "q" {
			return false, nil
		} else if itemNumberPhrase != "" {
			break
		}
	}

	fmt.Printf("\n")

	itemNumberRaw, err := strconv.ParseInt(itemNumberPhrase, 10, 32)
	log.PanicIf(err)

	itemNumber := int(itemNumberRaw)

	if itemNumber >= len(tagPhrases) {
		fmt.Printf("Invalid choice: (%d)", itemNumber)
		return true, nil
	}

	tagPhrase := tagPhrases[itemNumber]
	it := mapping[tagPhrase]

	err = printValues(it, ti)
	log.PanicIf(err)

	return true, nil
}

func main() {
	defer func() {
		if errRaw := recover(); errRaw != nil {
			err := errRaw.(error)
			log.PrintError(err)

			os.Exit(-2)
		}
	}()

	_, err := flags.Parse(arguments)
	if err != nil {
		os.Exit(-1)
	}

	if arguments.IsVerbose == true {
		cla := log.NewConsoleLogAdapter()
		log.AddAdapter("console", cla)

		scp := log.NewStaticConfigurationProvider()
		scp.SetLevelName("debug")

		log.LoadConfiguration(scp)
	}

	ti := exifextra.NewTreeIndex()

	err = ti.AddTree(arguments.Path)
	log.PanicIf(err)

	index := ti.Index()

	tagPhrases := make([]string, len(index))
	mapping := make(map[string]exifextra.IndexedTag)

	i := 0
	for it := range ti.Index() {
		tagPhrase := fmt.Sprintf("[%s] %s (0x%04x)", it.IfdPath, it.TagName, it.TagId)

		tagPhrases[i] = tagPhrase
		mapping[tagPhrase] = it

		i++
	}

	sort.Strings(tagPhrases)

	for {
		repeat, err := showTagsMenu(tagPhrases, mapping, ti)
		log.PanicIf(err)

		if repeat != true {
			break
		}

		fmt.Printf("\n")
	}
}
