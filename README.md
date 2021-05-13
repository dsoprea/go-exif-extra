[![Build Status](https://travis-ci.com/dsoprea/go-exif-extra.svg?branch=master)](https://travis-ci.com/dsoprea/go-exif-extra)
[![codecov](https://codecov.io/gh/dsoprea/go-exif-extra/branch/master/graph/badge.svg?token=Twxyx7kpAa)](https://codecov.io/gh/dsoprea/go-exif-extra)
[![Go Report Card](https://goreportcard.com/badge/github.com/dsoprea/go-exif-extra)](https://goreportcard.com/report/github.com/dsoprea/go-exif-extra)
[![Go Reference](https://pkg.go.dev/badge/github.com/dsoprea/go-exif-extra.svg)](https://pkg.go.dev/github.com/dsoprea/go-exif-extra)

# Overview

This repository provides EXIF and image support that transcends any one format
but lives close to the implementation.


# Tree Index

Filesystem paths or trees can be loaded in an index. Files that have a given tag
are grouped together. Values of all tags or specific tags can also be searched
case-insensitively.

This operation will identify and parse most any format that supports EXIF.

See the Go Reference link for usage examples.


# Tools


## ee_filesystem_indexer

This will allow you to browse through the tags of all loaded images.

```
$ ee_filesystem_indexer -p /general_images/discrete/2019/Trips/20190131\ Lisbon_Brussels_Netherlands_Lux_Germany/Images/Sony_RX100-VI/DCIM/10480201
(0) [IFD/Exif/Iop] InteroperabilityIndex (0x0001)    (16) [IFD/Exif] FNumber (0x829d)                     (32) [IFD/Exif] PixelYDimension (0xa003)             (48) [IFD1] Orientation (0x0112)
(1) [IFD/Exif/Iop] InteroperabilityVersion (0x0002)  (17) [IFD/Exif] FileSource (0xa300)                  (33) [IFD/Exif] RecommendedExposureIndex (0x8832)    (49) [IFD1] ResolutionUnit (0x0128)
(2) [IFD/Exif] BrightnessValue (0x9203)              (18) [IFD/Exif] Flash (0x9209)                       (34) [IFD/Exif] Saturation (0xa409)                  (50) [IFD1] Software (0x0131)
(3) [IFD/Exif] ColorSpace (0xa001)                   (19) [IFD/Exif] FlashpixVersion (0xa000)             (35) [IFD/Exif] SceneCaptureType (0xa406)            (51) [IFD1] XResolution (0x011a)
(4) [IFD/Exif] ComponentsConfiguration (0x9101)      (20) [IFD/Exif] FocalLength (0x920a)                 (36) [IFD/Exif] SceneType (0xa301)                   (52) [IFD1] YCbCrPositioning (0x0213)
(5) [IFD/Exif] CompressedBitsPerPixel (0x9102)       (21) [IFD/Exif] FocalLengthIn35mmFilm (0xa405)       (37) [IFD/Exif] SensitivityType (0x8830)             (53) [IFD1] YResolution (0x011b)
(6) [IFD/Exif] Contrast (0xa408)                     (22) [IFD/Exif] ISOSpeedRatings (0x8827)             (38) [IFD/Exif] Sharpness (0xa40a)                   (54) [IFD] DateTime (0x0132)
(7) [IFD/Exif] CustomRendered (0xa401)               (23) [IFD/Exif] LensSpecification (0xa432)           (39) [IFD/Exif] UserComment (0x9286)                 (55) [IFD] ImageDescription (0x010e)
(8) [IFD/Exif] DateTimeDigitized (0x9004)            (24) [IFD/Exif] LightSource (0x9208)                 (40) [IFD/Exif] WhiteBalance (0xa403)                (56) [IFD] Make (0x010f)
(9) [IFD/Exif] DateTimeOriginal (0x9003)             (25) [IFD/Exif] MakerNote (0x927c)                   (41) [IFD1] Compression (0x0103)                     (57) [IFD] Model (0x0110)
(10) [IFD/Exif] DigitalZoomRatio (0xa404)            (26) [IFD/Exif] MaxApertureValue (0x9205)            (42) [IFD1] DateTime (0x0132)                        (58) [IFD] Orientation (0x0112)
(11) [IFD/Exif] ExifVersion (0x9000)                 (27) [IFD/Exif] MeteringMode (0x9207)                (43) [IFD1] ImageDescription (0x010e)                (59) [IFD] ResolutionUnit (0x0128)
(12) [IFD/Exif] ExposureBiasValue (0x9204)           (28) [IFD/Exif] OffsetTime (0x9010)                  (44) [IFD1] JPEGInterchangeFormat (0x0201)           (60) [IFD] Software (0x0131)
(13) [IFD/Exif] ExposureMode (0xa402)                (29) [IFD/Exif] OffsetTimeDigitized (0x9012)         (45) [IFD1] JPEGInterchangeFormatLength (0x0202)     (61) [IFD] XResolution (0x011a)
(14) [IFD/Exif] ExposureProgram (0x8822)             (30) [IFD/Exif] OffsetTimeOriginal (0x9011)          (46) [IFD1] Make (0x010f)                            (62) [IFD] YCbCrPositioning (0x0213)
(15) [IFD/Exif] ExposureTime (0x829a)                (31) [IFD/Exif] PixelXDimension (0xa002)             (47) [IFD1] Model (0x0110)                           (63) [IFD] YResolution (0x011b)

Enter the number of a found tag (or 'q' to quit): 20

Occurrences of [FocalLength] in IFD [IFD/Exif]:

(5) files with value [1017/100]:

/general_images/discrete/2019/Trips/20190131 Lisbon_Brussels_Netherlands_Lux_Germany/Images/Sony_RX100-VI/DCIM/10480201/DSC00871.JPG
/general_images/discrete/2019/Trips/20190131 Lisbon_Brussels_Netherlands_Lux_Germany/Images/Sony_RX100-VI/DCIM/10480201/DSC01005.JPG
/general_images/discrete/2019/Trips/20190131 Lisbon_Brussels_Netherlands_Lux_Germany/Images/Sony_RX100-VI/DCIM/10480201/DSC01006.JPG
/general_images/discrete/2019/Trips/20190131 Lisbon_Brussels_Netherlands_Lux_Germany/Images/Sony_RX100-VI/DCIM/10480201/DSC01172.JPG
/general_images/discrete/2019/Trips/20190131 Lisbon_Brussels_Netherlands_Lux_Germany/Images/Sony_RX100-VI/DCIM/10480201/DSC01401.JPG


(4) files with value [1038/100]:

/general_images/discrete/2019/Trips/20190131 Lisbon_Brussels_Netherlands_Lux_Germany/Images/Sony_RX100-VI/DCIM/10480201/DSC01189.JPG
/general_images/discrete/2019/Trips/20190131 Lisbon_Brussels_Netherlands_Lux_Germany/Images/Sony_RX100-VI/DCIM/10480201/DSC01190.JPG
/general_images/discrete/2019/Trips/20190131 Lisbon_Brussels_Netherlands_Lux_Germany/Images/Sony_RX100-VI/DCIM/10480201/DSC01191.JPG
/general_images/discrete/2019/Trips/20190131 Lisbon_Brussels_Netherlands_Lux_Germany/Images/Sony_RX100-VI/DCIM/10480201/DSC01192.JPG


(10) files with value [1060/100]:

/general_images/discrete/2019/Trips/20190131 Lisbon_Brussels_Netherlands_Lux_Germany/Images/Sony_RX100-VI/DCIM/10480201/DSC00608.JPG
/general_images/discrete/2019/Trips/20190131 Lisbon_Brussels_Netherlands_Lux_Germany/Images/Sony_RX100-VI/DCIM/10480201/DSC00609.JPG
/general_images/discrete/2019/Trips/20190131 Lisbon_Brussels_Netherlands_Lux_Germany/Images/Sony_RX100-VI/DCIM/10480201/DSC00610.JPG
/general_images/discrete/2019/Trips/20190131 Lisbon_Brussels_Netherlands_Lux_Germany/Images/Sony_RX100-VI/DCIM/10480201/DSC00611.JPG
/general_images/discrete/2019/Trips/20190131 Lisbon_Brussels_Netherlands_Lux_Germany/Images/Sony_RX100-VI/DCIM/10480201/DSC00612.JPG
/general_images/discrete/2019/Trips/20190131 Lisbon_Brussels_Netherlands_Lux_Germany/Images/Sony_RX100-VI/DCIM/10480201/DSC00613.JPG
/general_images/discrete/2019/Trips/20190131 Lisbon_Brussels_Netherlands_Lux_Germany/Images/Sony_RX100-VI/DCIM/10480201/DSC00614.JPG
/general_images/discrete/2019/Trips/20190131 Lisbon_Brussels_Netherlands_Lux_Germany/Images/Sony_RX100-VI/DCIM/10480201/DSC00615.JPG
/general_images/discrete/2019/Trips/20190131 Lisbon_Brussels_Netherlands_Lux_Germany/Images/Sony_RX100-VI/DCIM/10480201/DSC00616.JPG
/general_images/discrete/2019/Trips/20190131 Lisbon_Brussels_Netherlands_Lux_Germany/Images/Sony_RX100-VI/DCIM/10480201/DSC01065.JPG

...
```


## ee_filesystem_searcher

This will search through the values of the tags of all loaded images.

```
$ ee_filesystem_searcher -p /general_images/discrete/2019/Trips/20190131\ Lisbon_Brussels_Netherlands_Lux_Germany/Images/Sony_RX100-VI/DCIM/10480201 --tag FocalLength --value 954/100
(2) results were found.

/general_images/discrete/2019/Trips/20190131 Lisbon_Brussels_Netherlands_Lux_Germany/Images/Sony_RX100-VI/DCIM/10480201/DSC00701.JPG  IFD/Exif  FocalLength  954/100
/general_images/discrete/2019/Trips/20190131 Lisbon_Brussels_Netherlands_Lux_Germany/Images/Sony_RX100-VI/DCIM/10480201/DSC00928.JPG  IFD/Exif  FocalLength  954/100
```
