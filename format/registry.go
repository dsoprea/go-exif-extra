package imageformats

import (
	"strings"

	"github.com/dsoprea/go-logging/v2"
	"github.com/dsoprea/go-utility/v2/image"
)

var (
	registry           = make(map[string]riimage.MediaParser)
	extensionsRegistry = make(map[string]string)
)

func register(name string, mp riimage.MediaParser, extensions []string) {
	defer func() {
		if errRaw := recover(); errRaw != nil {
			err := errRaw.(error)
			log.Panic(err)
		}
	}()

	_, found := registry[name]
	if found == true {
		log.Panicf("format [%s] already registered", name)
	}

	registry[name] = mp

	for _, extension := range extensions {
		if extension[0] != '.' {
			log.Panicf("format [%s] extension [%s] must include the prefixing period", name, extension)
		}

		extension = strings.ToLower(extension)

		_, found := extensionsRegistry[extension]
		if found == true {
			log.Panicf("format [%s] extension [%s] already registered", name, extension)
		}

		extensionsRegistry[extension] = name
	}
}

// Formats returns all registered media-parsers.
func Formats() (formats []riimage.MediaParser) {
	defer func() {
		if errRaw := recover(); errRaw != nil {
			err := errRaw.(error)
			log.Panic(err)
		}
	}()

	formats = make([]riimage.MediaParser, len(registry))

	i := 0
	for _, mp := range registry {
		formats[i] = mp
		i += 1
	}

	return formats
}

// GetFormatForExtension returns a MediaParser for the given extension or nil if
// none associated.
func GetFormatForExtension(extension string) (name string, mp riimage.MediaParser) {
	defer func() {
		if errRaw := recover(); errRaw != nil {
			err := errRaw.(error)
			log.Panic(err)
		}
	}()

	if len(extension) == 0 {
		log.Panicf("the extension is an empty string")
	} else if extension[0] != '.' {
		log.Panicf("the given extension must have the prefixing period: [%s]", extension)
	}

	extension = strings.ToLower(extension)
	name = extensionsRegistry[extension]

	if name == "" {
		return "", nil
	}

	mp = registry[name]

	return name, mp
}

// GetFormatForBytes returns the first MP that can parse the given bytes.
// This has an obvious cost.
func GetFormatForBytes(data []byte) (name string, mp riimage.MediaParser) {
	defer func() {
		if errRaw := recover(); errRaw != nil {
			err := errRaw.(error)
			log.Panic(err)
		}
	}()

	for name, mp = range registry {
		if mp.LooksLikeFormat(data) == true {
			return name, mp
		}
	}

	return "", nil
}
