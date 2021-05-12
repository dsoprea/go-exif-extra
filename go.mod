module github.com/dsoprea/go-exif-extra

go 1.16

// replace github.com/dsoprea/go-utility/v2 => ../go-utility/v2
// replace github.com/dsoprea/go-png-image-structure/v2 => ../go-png-image-structure/v2
// replace github.com/dsoprea/go-jpeg-image-structure/v2 => ../go-jpeg-image-structure/v2
// replace github.com/dsoprea/go-heic-exif-extractor/v2 => ../go-heic-exif-extractor/v2
// replace github.com/dsoprea/go-tiff-image-structure/v2 => ../go-tiff-image-structure/v2
// replace github.com/dsoprea/go-webp-image-structure => ../go-webp-image-structure

require (
	github.com/dsoprea/go-exif/v3 v3.0.0-20210512055020-8213cfabc61b
	github.com/dsoprea/go-heic-exif-extractor/v2 v2.0.0-20210512044107-62067e44c235
	github.com/dsoprea/go-jpeg-image-structure/v2 v2.0.0-20210512043942-b434301c6836
	github.com/dsoprea/go-logging v0.0.0-20200710184922-b02d349568dd
	github.com/dsoprea/go-png-image-structure v0.0.0-20210512210324-29b889a6093d // indirect
	github.com/dsoprea/go-png-image-structure/v2 v2.0.0-20210512044023-23bdd883ee8e
	github.com/dsoprea/go-tiff-image-structure/v2 v2.0.0-20210512044046-dc78da6a809b
	github.com/dsoprea/go-utility/v2 v2.0.0-20200717064901-2fccff4aa15e
	github.com/dsoprea/go-webp-image-structure v0.0.0-20210512044215-f98af2b0401e
)
