package main

import (
	"regexp"
	"strconv"
	"strings"
)

type ProductImage struct {
	ImagePath string
	ProductId string
	ImageId   string
	ImageName string
	Size      string
	Dpr       bool
	Ext       string
}

type ImageSize struct {
	Width  int
	Height int
	Mode   string
}

func ParseProductImageUrl(urlPath string) ProductImage {
	var re = regexp.MustCompile(productRe)
	matches := re.FindAllStringSubmatch(urlPath, -1)[0]
	image := ProductImage{
		ImagePath: matches[1],
		ProductId: matches[2],
		ImageId:   matches[3],
		ImageName: matches[4],
		Size:      matches[5],
		//Dpr:       matches[6],
		Ext:       matches[7],
	}
	if matches[6] != "" {
		image.Dpr = true
	}
	return image
}

func GetOriginalProductImagePath(image ProductImage) string {
	n := ""
	if image.ImageId == image.ImageName {
		n = image.ImageId
	} else {
		n = image.ImageId + "." + image.ImageName
	}
	file := "/wa-data/protected/shop/products/" + image.ImagePath + n + "." + image.Ext
	return file
}

func ParseSize(size string, dpr bool) ImageSize {
	mode := "unknown"
	arSize := strings.Split(size, "x")
	height := 0
	width, _ := strconv.Atoi(arSize[0])
	if len(arSize) == 2 {
		height, _ = strconv.Atoi(arSize[1])
	}

	if len(arSize) == 1 {
		mode = "max"
		height = width
	} else if width == height {
		mode = "crop"
	} else if width != 0 && height != 0 {
		mode = "rectangle"
	} else if width == 0 {
		mode = "height"
	} else if height == 0 {
		mode = "width"
	}

	return ImageSize{
		Width:  width,
		Height: height,
		Mode:   mode,
	}
}
