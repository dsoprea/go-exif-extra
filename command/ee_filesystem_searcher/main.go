package main

import (
	"fmt"
	"os"

	"github.com/dsoprea/go-exif-extra"
	"github.com/dsoprea/go-logging/v2"
	"github.com/jessevdk/go-flags"
	"github.com/ryanuber/columnize"
)

type parameters struct {
	Path      string   `short:"p" long:"path" required:"true" description:"Path to load recursively"`
	Value     string   `short:"V" long:"value" required:"true" description:"Substring to search"`
	Tags      []string `short:"t" long:"tag" description:"A tag to search within. Optional. May be provided multiple times."`
	IsVerbose bool     `short:"v" long:"verbose" description:"Print logging"`
}

var (
	arguments = new(parameters)
)

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

	hits, err := ti.Search(arguments.Value, arguments.Tags)
	log.PanicIf(err)

	fmt.Printf("(%d) results were found.\n", len(hits))
	fmt.Printf("\n")

	if len(hits) == 0 {
		return
	}

	lines := make([]string, len(hits))
	for i, hit := range hits {
		lines[i] = fmt.Sprintf("%s | %s | %s | %s", hit.Filepath, hit.IfdPath, hit.TagName, hit.ValuePhrase)
	}

	fmt.Println(columnize.SimpleFormat(lines))
}
