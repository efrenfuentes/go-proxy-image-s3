package image

import (
	"errors"
	"regexp"
	"strconv"

	"github.com/disintegration/imaging"
)

type geometryStruct struct {
	width  int
	height int
}

func Transform(filename string, resolution string) (*string, error) {
	if resolution == "original" {
		return &filename, nil
	}

	geometry, err := getGeometry(resolution)
	if err != nil {
		return nil, err
	}

	src, err := imaging.Open(filename, imaging.AutoOrientation(true))
	if err != nil {
		return nil, err
	}

	dstImage := imaging.Fill(src, geometry.width, geometry.height, imaging.Center, imaging.Box)

	outputFilename := resolution + "_" + filename
	err = imaging.Save(dstImage, outputFilename)
	if err != nil {
		return nil, err
	}

	return &outputFilename, nil
}

func getGeometry(geometry string) (*geometryStruct, error) {
	geometryRE, _ := regexp.Compile(`(\d+)x(\d+)|original`) // Prepare our regex
	result := geometryRE.FindStringSubmatch(geometry)

	if len(result) > 0 {
		width, _ := strconv.Atoi(result[1])
		height, _ := strconv.Atoi(result[2])

		return &geometryStruct{width: width, height: height}, nil
	}

	return nil, errors.New("invalid geometry")

}
